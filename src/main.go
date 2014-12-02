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

    ecsClient.DescribeRegions()
    ecsClient.DescribeZones("cn-hangzhou")

    log.Println("----- Test End -----")
}
