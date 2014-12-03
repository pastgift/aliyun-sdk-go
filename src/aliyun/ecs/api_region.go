package ecs

import (
)

/* API on Region */
// Get `Region` list
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
    req := NewRequest("DescribeRegions")
    res := new(DescribeRegionsResult)

    errRes := self.CallAPI(req, res)

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
    req := NewRequest("DescribeZones")
    res := new(DescribeZonesResult)

    //req.SetArg("RegionId", args.RegionId)
    req.SetArgs(args)
    errRes := self.CallAPI(req, res)

    // Additional error message for user
    if errRes == nil && len(res.Zones.Zone) == 0 {
        errRes = &ErrorResult{Message:"No Zone Found"}
    }

    return res, errRes
}
