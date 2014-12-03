package ecs

import (
)

/* API on Instance */
// Create `Instance`
// TODO
type CreateInstanceResult struct {
    GlobalResult
    //...
}

func (self *Client) CreateInstance() (result CreateInstanceResult, errorResult *ErrorResult) {
    res := CreateInstanceResult{}
    //...
    return res, nil
}

// Start `Instance`
// TODO
type StartInstanceResult struct {
    GlobalResult
    //...
}

func (self *Client) StartInstance() (result StartInstanceResult, errorResult *ErrorResult) {
    res := StartInstanceResult{}
    //...
    return res, nil
}

// Stop `Instance`
// TODO
type StopInstanceResult struct {
    GlobalResult
    //...
}

func (self *Client) StopInstance() (result StopInstanceResult, errorResult *ErrorResult) {
    res := StopInstanceResult{}
    //...
    return res, nil
}

// Reboot `Instance`
// TODO
type RebootInstanceResult struct {
    GlobalResult
    //...
}

func (self *Client) RebootInstance() (result RebootInstanceResult, errorResult *ErrorResult) {
    res := RebootInstanceResult{}
    //...
    return res, nil
}

// Modify `Instance` attribute
type ModifyInstanceAttributeArgs struct {
    InstanceID string
    InstanceName string
    Description string
    Password string
    HostName string
}

type ModifyInstanceAttributeResult struct {
    GlobalResult
}

func (self *Client) ModifyInstanceAttribute(args *ModifyInstanceAttributeArgs) (result *ModifyInstanceAttributeResult, errorResult *ErrorResult) {
    req := NewRequest("DescribeRegions")
    res := new(ModifyInstanceAttributeResult)

    req.SetArgs(args)
    errRes := self.CallAPI(req, res)

    return res, errRes
}

// Describe `Instance` status
// TODO
type DescribeInstanceStatusResult struct {
    GlobalResult
    //...
}

func (self *Client) DescribeInstanceStatus() (result DescribeInstanceStatusResult, errorResult *ErrorResult) {
    res := DescribeInstanceStatusResult{}
    //...
    return res, nil
}

// Describe `Instance` attribute
// TODO
type DescribeInstanceAttributeResult struct {
    GlobalResult
    //...
}

func (self *Client) DescribeInstanceAttribute() (result DescribeInstanceAttributeResult, errorResult *ErrorResult) {
    res := DescribeInstanceAttributeResult{}
    //...
    return res, nil
}

// Delete `Instance`
// TODO
type DeleteInstanceResult struct {
    GlobalResult
    //...
}

func (self *Client) DeleteInstance() (result DeleteInstanceResult, errorResult *ErrorResult) {
    res := DeleteInstanceResult{}
    //...
    return res, nil
}

// Join `SecurityGroup`
// TODO
type JoinSecurityGroupResult struct {
    GlobalResult
    //...
}

func (self *Client) JoinSecurityGroup() (result JoinSecurityGroupResult, errorResult *ErrorResult) {
    res := JoinSecurityGroupResult{}
    //...
    return res, nil
}

// Leave `SecurityGroup`
// TODO
type LeaveSecurityGroupResult struct {
    GlobalResult
    //...
}

func (self *Client) LeaveSecurityGroup() (result LeaveSecurityGroupResult, errorResult *ErrorResult) {
    res := LeaveSecurityGroupResult{}
    //...
    return res, nil
}
