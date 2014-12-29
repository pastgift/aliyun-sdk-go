// API on Region

package ecs

import ()

// Get `Region` list
type DescribeRegionsArgs struct {
	// NO ARGUMENT NEEDED
}

type RegionType struct {
	RegionId  string `json:"RegionID"`
	LocalName string `json:"LocalName"`
}

type RegionType_Array struct {
	Region []RegionType `json:"Region"`
}

type DescribeRegionsResult struct {
	GlobalResult

	Regions RegionType_Array `json:"Regions"`
}

func (self *Client) DescribeRegions() (result *DescribeRegionsResult, errorResult *ErrorResult) {
	args := new(DescribeRegionsArgs)
	res := new(DescribeRegionsResult)
	errRes := self.CallAPI("DescribeRegions", args, res)

	return res, errRes
}

/**************************************************/

// Get `Zone` list
type DescribeZonesArgs struct {
	RegionId string
}

type AvailableResourceCreationType struct {
	ResourceTypes []string `json:"ResourceTypes"`
}

type AvailableDiskCategoriesType struct {
	DiskCategories []string `json:"DiskCategories"`
}

type ZoneType struct {
	ZoneId                    string                        `json:"ZoneId"`
	LocalName                 string                        `json:"LocalName"`
	AvailableResourceCreation AvailableResourceCreationType `json:"AvailableResourceCreation"`
	AvailableDiskCategories   AvailableDiskCategoriesType   `json:"AvailableDiskCategories"`
}

type ZoneType_Array struct {
	Zone []ZoneType `json:"Zone"`
}

type DescribeZonesResult struct {
	GlobalResult

	Zones ZoneType_Array `json:"Zones"`
}

func (self *Client) DescribeZones(args *DescribeZonesArgs) (result *DescribeZonesResult, errorResult *ErrorResult) {
	res := new(DescribeZonesResult)
	errRes := self.CallAPI("DescribeZones", args, res)

	// If there is no Zone, API will not return Error
	if errRes == nil && len(res.Zones.Zone) == 0 {
		errRes = &ErrorResult{Message: "No Zone Found"}
	}

	return res, errRes
}
