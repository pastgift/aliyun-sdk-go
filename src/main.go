package main

import (
    "fmt"

    "aliyun/ecs"
)

func main() {
    fmt.Println("----- Test Start -----")

    ecs_client := ecs.NewEcsClient("access_key_id","access_key_secret")
    res, err := ecs_client.DescribeRegions()
    if err == nil {
        fmt.Println(res)
    }

    fmt.Println("----- Test End -----")
}
