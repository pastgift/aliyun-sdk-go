package main

import (
    "log"
    "strings"
    "io/ioutil"

    "aliyun/ecs"
)

func main() {
    log.Println("----- Test Start -----")

    // Get AccessKeyId and AccessKeySecret
    f, err := ioutil.ReadFile("src/_key.txt")
    if err != nil {
        log.Println("Can't open `_key.txt`")
        return
    }

    accessKey       := strings.Split(string(f), "\n")
    accessKeyId     := accessKey[0]
    accessKeySecret := accessKey[1]

    log.Println("Useing AccessKeyId:    ", accessKeyId)
    log.Println("Useing AccessKeySecret:", accessKeySecret)

    // Example usage
    ecsClient := ecs.NewClient(accessKeyId,accessKeySecret)

    regionList, errRes := ecsClient.DescribeRegions()
    if errRes == nil {
        log.Println("First Region:", regionList.Regions.Region[0].RegionId)
    } else {
        log.Println("Error:", errRes.Message)
    }

    zoneList, errRes := ecsClient.DescribeZones(
        &ecs.DescribeZonesArgs{
            RegionId:"cn-hangzhou",
        })
    if errRes == nil {
        log.Println("First Zone:  ", zoneList.Zones.Zone[0].ZoneId)
    } else {
        log.Println("Error:", errRes.Message)
    }

    log.Println("----- Test End -----")
}