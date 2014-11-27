package main

import (
    "fmt"
    "strings"
    "io/ioutil"

    "aliyun/ecs"
)

func main() {
    fmt.Println("----- Test Start -----")

    // Get AccessKeyId and AccessKeySecret
    f, err := ioutil.ReadFile("src/_key.txt")
    if err != nil {
        fmt.Println("Can't open `_key.txt`")
        return
    }

    access_key        := strings.Split(string(f), "\n")
    access_key_id     := access_key[0]
    access_key_secret := access_key[1]

    fmt.Println("Useing AccessKeyId:\n\t",     access_key_id)
    fmt.Println("Useing AccessKeySecret:\n\t", access_key_secret)
    fmt.Println("")

    // Example usage
    ecs_client := ecs.NewEcsClient(access_key_id,access_key_secret)
    res, err := ecs_client.DescribeRegions()
    if err == nil {
        fmt.Println("Request URL:\n\t", ecs_client.ApiUrl.String())
        fmt.Println("API result:\n\t", res)
    }

    fmt.Println("----- Test End -----")
}
