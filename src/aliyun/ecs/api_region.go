package ecs

import (
)

/* API on Region */
// Get `Region` list
type DescribeRegionsArgs struct {
}

type DescribeRegionsResult struct {
    GlobalResult
    Regions struct {
        Region []struct {
            RegionId string `json:"RegionID"`
            LocalName string `json:"LocalName"`
        } `json:"Region"`
    } `json:"Regions"`
}

func (self *Client) DescribeRegions() (result *DescribeRegionsResult, errorResult *ErrorResult) {
    act  := "DescribeRegions"
    args := new(DescribeRegionsArgs)
    res  := new(DescribeRegionsResult)
    errRes := self.CallAPI(act, args, res)

    return res, errRes
}

// Get `Zone` list
type DescribeZonesArgs struct {
    RegionId string
}

type DescribeZonesResult struct {
    GlobalResult
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

func (self *Client) DescribeZones(args *DescribeZonesArgs) (result *DescribeZonesResult, errorResult *ErrorResult) {
    act := "DescribeZones"
    res := new(DescribeZonesResult)
    errRes := self.CallAPI(act, args, res)

    // If there is no Zone, API will not return Error
    if errRes == nil && len(res.Zones.Zone) == 0 {
        errRes = &ErrorResult{Message:"No Zone Found"}
    }

    return res, errRes
}
