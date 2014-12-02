package ecs

import (
    "errors"
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

func (self *Client) DescribeRegions() (result *DescribeRegionsResult, err error) {
    req := NewRequest("DescribeRegions")
    res := new(DescribeRegionsResult)

    err = self.CallAPI(req, res)
    return res, err
}

// Get `Zone` list
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

func (self *Client) DescribeZones(regionId string) (result *DescribeZonesResult, err error) {
    req := NewRequest("DescribeZones")
    res := new(DescribeZonesResult)

    if regionId == "" {
        return res, errors.New("RegionId should not be blank")
    }

    req.SetArg("RegionId", regionId)

    err = self.CallAPI(req, res)
    return res, err
}
