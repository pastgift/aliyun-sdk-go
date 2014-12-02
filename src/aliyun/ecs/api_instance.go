package ecs

import (
    "errors"
)

/* API on Instance */
// Create `Instance`
// TODO
type CreateInstanceResult struct {
    GlobalResult
    //...
}

func (self *Client) CreateInstance() (result CreateInstanceResult, err error) {
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

func (self *Client) StartInstance() (result StartInstanceResult, err error) {
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

func (self *Client) StopInstance() (result StopInstanceResult, err error) {
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

func (self *Client) RebootInstance() (result RebootInstanceResult, err error) {
    res := RebootInstanceResult{}
    //...
    return res, nil
}

// Modify `Instance` attribute
type ModifyInstanceAttributeResult struct {
    GlobalResult
}

func (self *Client) ModifyInstanceAttribute(instanceID string, instanceName string, description string, password string, hostName string) (result *ModifyInstanceAttributeResult, err error) {
    req := NewRequest("DescribeRegions")
    res := new(ModifyInstanceAttributeResult)

    if instanceID == "" {
        return res, errors.New("InstanceID should not be blank")
    }

    req.SetArg("InstanceID", instanceID)
    req.SetArg("InstanceName", instanceName)
    req.SetArg("Description", description)
    req.SetArg("Password", password)
    req.SetArg("HostName", hostName)

    err = self.CallAPI(req, res)
    return res, err
}

// Describe `Instance` status
// TODO
type DescribeInstanceStatusResult struct {
    GlobalResult
    //...
}

func (self *Client) DescribeInstanceStatus() (result DescribeInstanceStatusResult, err error) {
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

func (self *Client) DescribeInstanceAttribute() (result DescribeInstanceAttributeResult, err error) {
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

func (self *Client) DeleteInstance() (result DeleteInstanceResult, err error) {
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

func (self *Client) JoinSecurityGroup() (result JoinSecurityGroupResult, err error) {
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

func (self *Client) LeaveSecurityGroup() (result LeaveSecurityGroupResult, err error) {
    res := LeaveSecurityGroupResult{}
    //...
    return res, nil
}
