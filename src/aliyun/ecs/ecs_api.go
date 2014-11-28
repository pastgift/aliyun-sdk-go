package ecs

import (
    "log"
)

type RegionJson struct {
    requestId string`json:"RequestID"`
    Regions struct {    
        Region []struct {   
            RegionId string `json:"RegionID"`
            LocalName string `json:"LocalName"`
        } `json:"Region"`
    } `json:"Regions"`
}

func (self *ECSClient) DescribeRegions() (result RegionJson, err error) {
    regionJson := RegionJson{}

    resp, err := self.Do(NewECSRequest("DescribeRegions"))
    if err != nil {
        log.Println("Do Request Error")
        return regionJson, err
    }

    // Output as string
    log.Println("Received:\n\t", resp.String())
    
    // Output as Struct
    err = resp.FillJson(&regionJson)
    if err != nil {
        log.Println("Load Json Error")
        return regionJson, err
    }

    return regionJson, err
}