// API on Region

package ecs

import (
)

// Get `Region` list
type DescribeRegionsArgs struct {
    // NO ARGUMENT NEEDED
}

type RegionST struct {
    RegionId    string      `json:"RegionID"`
    LocalName   string      `json:"LocalName"`
}

type RegionsST struct {
    Region      []RegionST  `json:"Region"`
}

type DescribeRegionsResult struct {
    GlobalResult

    Regions     RegionsST   `json:"Regions"`
}

func (self *Client) DescribeRegions() (result *DescribeRegionsResult, errorResult *ErrorResult) {
    args := new(DescribeRegionsArgs)
    res  := new(DescribeRegionsResult)
    errRes := self.CallAPI("DescribeRegions", args, res)

    return res, errRes
}

/**************************************************/

// Get `Zone` list
type DescribeZonesArgs struct {
    RegionId                    string
}

type AvailableDiskCategoriesST struct {
    DiskCategories              []string                    `json:"DiskCategories"`
}

type AvailableResourceCreationST struct {
    ResourceTypes               []string                    `json:"ResourceTypes"`
}

type ZoneST struct {
    ZoneId                      string                      `json:"ZoneId"`
    LocalName                   string                      `json:"LocalName"`
    AvailableDiskCategories     AvailableDiskCategoriesST   `json:"AvailableDiskCategories"`
    AvailableResourceCreation   AvailableResourceCreationST `json:"AvailableResourceCreation"`
}

type ZonesST struct {
    Zone                        []ZoneST                    `json:"Zone"`
}

type DescribeZonesResult struct {
    GlobalResult

    Zones                       ZonesST                     `json:"Zones"`
}

func (self *Client) DescribeZones(args *DescribeZonesArgs) (result *DescribeZonesResult, errorResult *ErrorResult) {
    res := new(DescribeZonesResult)
    errRes := self.CallAPI("DescribeZones", args, res)

    // If there is no Zone, API will not return Error
    if errRes == nil && len(res.Zones.Zone) == 0 {
        errRes = &ErrorResult{Message:"No Zone Found"}
    }

    return res, errRes
}
