// API on Other

package ecs

import ()

// Describe `Instance` Types
type DescribeInstanceTypesArgs struct {
	// NO ARGUMENT NEEDED
}

type InstanceTypeItemType struct {
	InstanceTypeId string  `json:"InstanceTypeId"`
	CpuCoreCount   int     `json:"CpuCoreCount"`
	MemorySize     float64 `json:"MemorySize"`
}

type InstanceTypeItemType_Array struct {
	InstanceType []InstanceTypeItemType `json:"InstanceType"`
}

type DescribeInstanceTypesResult struct {
	GlobalResult

	InstanceTypes InstanceTypeItemType_Array `json:"InstanceTypes"`
}

func (self *Client) DescribeInstanceTypes(args *DescribeInstanceTypesArgs) (result *DescribeInstanceTypesResult, errorResult *ErrorResult) {
	res := new(DescribeInstanceTypesResult)
	errRes := self.CallAPI("DescribeInstanceTypes", args, res)

	return res, errRes
}
