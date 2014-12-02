package ecs

import (
    "log"
)

/* API on Region */
// Get `Region` list
type RegionList struct {
    requestId string `json:"RequestID"`
    Regions struct {
        Region []struct {
            RegionId string `json:"RegionID"`
            LocalName string `json:"LocalName"`
        } `json:"Region"`
    } `json:"Regions"`
}

func (self *Client) DescribeRegions() (result RegionList, err error) {
    res := RegionList{}

    req := NewRequest("DescribeRegions")

    resp, err := self.Do(req)
    if err != nil {
        log.Println("Do Request Error")
        return res, err
    }

    err = resp.FillJson(&res)
    if err != nil {
        log.Println("Load Json Error")
        return res, err
    }

    log.Println("Received:\n\t", resp.String())
    return res, err
}

// Get `Zone` list
type ZoneList struct {
    requestId string`json:"RequestID"`
    Zones struct {
        Zone []struct {
            ZoneId string `json:"ZoneId"`
            LocalName string `json:"LocalName"`
            AvailableDiskCategories struct {
                DiskCategories []string `json:"DiskCategories"`
            } `json:"AvailableDiskCategories"`
            AvailableResourceCreation struct {
                ResourceTypes []string `json:"ResourceTypes"`
            } `json:"AvailableResourceCreation"`
        } `json:"Zone"`
    } `json:"Zones"`
}

func (self *Client) DescribeZones(regionId string) (result ZoneList, err error) {
    res := ZoneList{}

    req := NewRequest("DescribeZones")
    req.SetArg("RegionId", regionId)

    resp, err := self.Do(req)
    if err != nil {
        log.Println("Do Request Error")
        return res, err
    }

    err = resp.FillJson(&res)
    if err != nil {
        log.Println("Load Json Error")
        return res, err
    }

    log.Println("Received:\n\t", resp.String())
    return res, err
}
