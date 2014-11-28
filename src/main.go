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

    access_key        := strings.Split(string(f), "\n")
    access_key_id     := access_key[0]
    access_key_secret := access_key[1]

    log.Println("Useing AccessKeyId:    ",     access_key_id)
    log.Println("Useing AccessKeySecret:", access_key_secret)

    // Example usage
    ecs_client := ecs.NewECSClient(access_key_id,access_key_secret)
    result, err := ecs_client.DescribeRegions()
    if err == nil {
        log.Printf("API result:\n\t %+v\n", result)
    }

    log.Println("----- Test End -----")
}