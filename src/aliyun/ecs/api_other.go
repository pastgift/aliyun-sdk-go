// API on Other

package ecs

import (
)

// Describe `Instance` Types
type DescribeInstanceTypesArgs struct {

}

type DescribeInstanceTypesResult struct {
    GlobalResult
    
    InstanceTypes   struct {
        InstanceType    []struct {
            InstanceTypeId  string      `json:"InstanceTypeId"`
            CpuCoreCount    int         `json:"InstanceTypeId"`
            MemorySize      float       `json:"InstanceTypeId"`
        }                           `json:"InstanceType"`
    }                           `json:"InstanceTypes"`
}

func (self *Client) DescribeInstanceTypes(args *DescribeInstanceTypesArgs) (result *DescribeInstanceTypesResult, errorResult *ErrorResult) {
    res := new(DescribeInstanceTypesResult)
    errRes := self.CallAPI("DescribeInstanceTypes", args, res)

    return res, errRes
}