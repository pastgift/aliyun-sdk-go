// API on Instance

package ecs

import (
)

// Create `Instance`
// TODO
type CreateInstanceResult struct {
    GlobalResult
    //...
}

func (self *Client) CreateInstance() (result *CreateInstanceResult, errorResult *ErrorResult) {
    res := new(CreateInstanceResult)
    //...
    return res, nil
}

// Start `Instance`
// TODO
type StartInstanceResult struct {
    GlobalResult
    //...
}

func (self *Client) StartInstance() (result *StartInstanceResult, errorResult *ErrorResult) {
    res := new(StartInstanceResult)
    //...
    return res, nil
}

// Stop `Instance`
// TODO
type StopInstanceResult struct {
    GlobalResult
    //...
}

func (self *Client) StopInstance() (result *StopInstanceResult, errorResult *ErrorResult) {
    res := new(StopInstanceResult)
    //...
    return res, nil
}

// Reboot `Instance`
// TODO
type RebootInstanceResult struct {
    GlobalResult
    //...
}

func (self *Client) RebootInstance() (result *RebootInstanceResult, errorResult *ErrorResult) {
    res := new(RebootInstanceResult)
    //...
    return res, nil
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
// TODO
type DeleteInstanceResult struct {
    GlobalResult
    //...
}

func (self *Client) DeleteInstance() (result *DeleteInstanceResult, errorResult *ErrorResult) {
    res := new(DeleteInstanceResult)
    //...
    return res, nil
}

// Join `SecurityGroup`
// TODO
type JoinSecurityGroupResult struct {
    GlobalResult
    //...
}

func (self *Client) JoinSecurityGroup() (result *JoinSecurityGroupResult, errorResult *ErrorResult) {
    res := new(JoinSecurityGroupResult)
    //...
    return res, nil
}

// Leave `SecurityGroup`
// TODO
type LeaveSecurityGroupResult struct {
    GlobalResult
    //...
}

func (self *Client) LeaveSecurityGroup() (result *LeaveSecurityGroupResult, errorResult *ErrorResult) {
    res := new(LeaveSecurityGroupResult)
    //...
    return res, nil
}
