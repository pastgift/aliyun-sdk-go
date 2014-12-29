package main

import (
	"io/ioutil"
	"log"
	"strings"

	"aliyun/ecs"
	"aliyun/util"
)

func main() {
	log.Println("----- Test Start -----")
	log.Println(util.CreateSpecifiedTimestampString(2014, 01, 02, 03, 04, 05))

	// Get AccessKeyId and AccessKeySecret
	f, err := ioutil.ReadFile("src/_key.txt")
	if err != nil {
		log.Println("Can't open `_key.txt`")
		return
	}

	accessKey := strings.Split(string(f), "\n")
	accessKeyId := accessKey[0]
	accessKeySecret := accessKey[1]

	log.Println("Useing AccessKeyId:    ", accessKeyId)
	log.Println("Useing AccessKeySecret:", accessKeySecret)

	// Example usage
	ecsClient := ecs.NewClient(accessKeyId, accessKeySecret)

	// DescribeRegions
	regionList, errRes := ecsClient.DescribeRegions()
	if errRes == nil {
		log.Println("First Region:", regionList.Regions.Region[0].RegionId)
	} else {
		log.Println("Error:", errRes.Message)
	}

	// DescribeZones
	zoneList, errRes := ecsClient.DescribeZones(
		&ecs.DescribeZonesArgs{
			RegionId: "cn-hangzhou",
		})
	if errRes == nil {
		log.Println("First Zone:  ", zoneList.Zones.Zone[0].ZoneId)
	} else {
		log.Println("Error:", errRes.Message)
	}

	// DescribeImages
	_, errRes = ecsClient.DescribeImages(
		&ecs.DescribeImagesArgs{
			RegionId: "cn-hangzhou",
//			PageNumber:"1",
//			PageSize:"50",
//			ImageId : "m-23sh9otvi",
			ImageId : "m-237pyempt",
		})
	if errRes == nil {
		log.Println("Images Found:")
	} else {
		log.Println("Error:", errRes.Message)
	}
	log.Println("----- Test End -----")
}
