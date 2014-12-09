// API on Network

package ecs

import (
)

// Allocate Public Ip Address
type AllocatePublicIpAddressArgs struct {
    InstanceId  string
}

type AllocatePublicIpAddressResult struct {
    GlobalResult

    IpAddress   string  `json:"IpAddress"`
}

func (self *Client) AllocatePublicIpAddress(args *AllocatePublicIpAddressArgs) (result *AllocatePublicIpAddressResult, errorResult *ErrorResult) {
    res := new(AllocatePublicIpAddressResult)
    errRes := self.CallAPI("AllocatePublicIpAddress", args, res)

    return res, errRes
}
