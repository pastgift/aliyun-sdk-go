// API on Instance

package ecs

import (
)

// Create `Instance`
type CreateInstanceArgs struct {
}

type CreateInstanceResult struct {
    GlobalResult
    // TODO
}

func (self *Client) CreateInstance(args *CreateInstanceArgs) (result *CreateInstanceResult, errorResult *ErrorResult) {
    res := new(CreateInstanceResult)
    errRes := self.CallAPI("CreateInstance", args, res)

    return res, errRes
}

// Start `Instance`
type StartInstanceArgs struct {
}

type StartInstanceResult struct {
    GlobalResult
    // TODO
}

func (self *Client) StartInstance(args *StartInstanceArgs) (result *StartInstanceResult, errorResult *ErrorResult) {
    res := new(StartInstanceResult)
    errRes := self.CallAPI("StartInstance", args, res)

    return res, errRes
}

// Stop `Instance`
type StopInstanceArgs struct {
}

type StopInstanceResult struct {
    GlobalResult
    // TODO
}

func (self *Client) StopInstance(args *StopInstanceArgs) (result *StopInstanceResult, errorResult *ErrorResult) {
    res := new(StopInstanceResult)
    errRes := self.CallAPI("StopInstance", args, res)

    return res, errRes
}

// Reboot `Instance`
type RebootInstanceArgs struct {
}

type RebootInstanceResult struct {
    GlobalResult
    // TODO
}

func (self *Client) RebootInstance(args *RebootInstanceArgs) (result *RebootInstanceResult, errorResult *ErrorResult) {
    res := new(RebootInstanceResult)
    errRes := self.CallAPI("RebootInstance", args, res)

    return res, errRes
}

// Modify `Instance` attribute
type ModifyInstanceAttributeArgs struct {
    InstanceID      string
    InstanceName    string
    Description     string
    Password        string
    HostName        string
}

type ModifyInstanceAttributeResult struct {
    GlobalResult
}

func (self *Client) ModifyInstanceAttribute(args *ModifyInstanceAttributeArgs) (result *ModifyInstanceAttributeResult, errorResult *ErrorResult) {
    res := new(ModifyInstanceAttributeResult)
    errRes := self.CallAPI("ModifyInstanceAttribute", args, res)

    return res, errRes
}

// Describe `Instance` status
type DescribeInstanceStatusArgs struct {
    RegionId    string
    ZoneId      string
    PageNumber  string
    PageSize    string
}

type DescribeInstanceStatusResult struct {
    GlobalResult

    TotalCount          int `json:"TotalCount"`
    PageNumber          int `json:"PageNumber"`
    PageSize            int `json:"PageSize"`
    InstanceStatuses    struct {
        InstanceStatus      []struct {
            InstanceId          string  `json:"InstanceId"`
            Status              string  `json:"Status"`
        }                   `json:InstanceStatus`
    }                       `json:"InstanceStatuses"`
}

func (self *Client) DescribeInstanceStatus(args *DescribeInstanceStatusArgs) (result *DescribeInstanceStatusResult, errorResult *ErrorResult) {
    res := new(DescribeInstanceStatusResult)
    errRes := self.CallAPI("DescribeInstanceStatus", args, res)

    return res, errRes
}

// Describe `Instance` attribute
type DescribeInstanceAttributeArgs struct {
    InstanceId string
}

type DescribeInstanceAttributeResult struct {
    GlobalResult

    ClusterId               string  `json:"ClusterId"`
    CreationTime            string  `json:"CreationTime"`
    Description             string  `json:"Description"`
    HostName                string  `json:"HostName"`
    ImageId                 string  `json:"ImageId"`
    InnerIpAddress          struct {
        IpAddress               []string    `json:"IpAddress"`
    }                               `json:"InnerIpAddress"`
    InstanceId              string  `json:"InstanceId"`
    InstanceName            string  `json:"InstanceName"`
    InstanceType            string  `json:"InstanceType"`
    InternetChargeType      string  `json:"InternetChargeType"`
    InternetMaxBandwidthIn  int     `json:"InternetMaxBandwidthIn"`
    InternetMaxBandwidthOut int     `json:"InternetMaxBandwidthOut"`
    OperationLocks          struct {
        OperationLock           []string    `json:"OperationLock"`
    }                               `json:"OperationLocks"`
    PublicIpAddress         struct {
        IpAddress               []string    `json:"IpAddress"`
    }                               `json:"PublicIpAddress"`
    RegionId                string  `json:"RegionId"`
    SecurityGroupIds        struct {
        SecurityGroupId         []string    `json:"SecurityGroupId"`
    }                               `json:"SecurityGroupIds"`
    SerialNumber            string  `json:"SerialNumber"`
    Status                  string  `json:"Status"`
    VlanId                  string  `json:"VlanId"`
    ZoneId                  string  `json:"ZoneId"`
}

func (self *Client) DescribeInstanceAttribute(args *DescribeInstanceAttributeArgs ) (result *DescribeInstanceAttributeResult, errorResult *ErrorResult) {
    res := new(DescribeInstanceAttributeResult)
    errRes := self.CallAPI("DescribeInstanceAttribute", args, res)

    return res, errRes
}

// Delete `Instance`
type DeleteInstanceArgs struct {
}

type DeleteInstanceResult struct {
    GlobalResult
    // TODO
}

func (self *Client) DeleteInstance(args *DeleteInstanceArgs) (result *DeleteInstanceResult, errorResult *ErrorResult) {
    res := new(DeleteInstanceResult)
    errRes := self.CallAPI("DeleteInstance", args, res)

    return res, errRes
}

// Join `SecurityGroup`
type JoinSecurityGroupArgs struct {
}

type JoinSecurityGroupResult struct {
    GlobalResult
    // TODO
}

func (self *Client) JoinSecurityGroup(args *JoinSecurityGroupArgs) (result *JoinSecurityGroupResult, errorResult *ErrorResult) {
    res := new(JoinSecurityGroupResult)
    errRes := self.CallAPI("JoinSecurityGroup", args, res)

    return res, errRes
}

// Leave `SecurityGroup`
type LeaveSecurityGroupArgs struct {
}

type LeaveSecurityGroupResult struct {
    GlobalResult
    // TODO
}

func (self *Client) LeaveSecurityGroup(args *LeaveSecurityGroupArgs) (result *LeaveSecurityGroupResult, errorResult *ErrorResult) {
    res := new(LeaveSecurityGroupResult)
    errRes := self.CallAPI("LeaveSecurityGroup", args, res)

    return res, errRes
}