// API on Instance

package ecs

import (
)

// Create `Instance`
type CreateInstanceArgs struct {
    RegionId                string
    ZoneId                  string
    ImageId                 string
    InstanceType            string
    SecurityGroupId         string
    InstanceName            string
    Description             string
    InternetChargeType      string
    InternetMaxBandwidthIn  int
    InternetMaxBandwidthOut int
    HostName                string
    Password                string
    SystemDisk.Category     string
    SystemDisk.DiskName     string
    SystemDisk.Description  string

    DataDisk_1_Size         int     `ArgName:"DataDisk.1.Size"`
    DataDisk_1_Category     string  `ArgName:"DataDisk.1.Category"`
    DataDisk_1_SnapshotId   string  `ArgName:"DataDisk.1.SnapshotId"`
    DataDisk_1_DiskName     string  `ArgName:"DataDisk.1.DiskName"`
    DataDisk_1_Description  string  `ArgName:"DataDisk.1.Description"`
    DataDisk_1_Device       string  `ArgName:"DataDisk.1.Device"`

    DataDisk_2_Size         int     `ArgName:"DataDisk.2.Size"`
    DataDisk_2_Category     string  `ArgName:"DataDisk.2.Category"`
    DataDisk_2_SnapshotId   string  `ArgName:"DataDisk.2.SnapshotId"`
    DataDisk_2_DiskName     string  `ArgName:"DataDisk.2.DiskName"`
    DataDisk_2_Description  string  `ArgName:"DataDisk.2.Description"`
    DataDisk_2_Device       string  `ArgName:"DataDisk.2.Device"`

    DataDisk_3_Size         int     `ArgName:"DataDisk.3.Size"`
    DataDisk_3_Category     string  `ArgName:"DataDisk.3.Category"`
    DataDisk_3_SnapshotId   string  `ArgName:"DataDisk.3.SnapshotId"`
    DataDisk_3_DiskName     string  `ArgName:"DataDisk.3.DiskName"`
    DataDisk_3_Description  string  `ArgName:"DataDisk.3.Description"`
    DataDisk_3_Device       string  `ArgName:"DataDisk.3.Device"`

    DataDisk_4_Size         int     `ArgName:"DataDisk.4.Size"`
    DataDisk_4_Category     string  `ArgName:"DataDisk.4.Category"`
    DataDisk_4_SnapshotId   string  `ArgName:"DataDisk.4.SnapshotId"`
    DataDisk_4_DiskName     string  `ArgName:"DataDisk.4.DiskName"`
    DataDisk_4_Description  string  `ArgName:"DataDisk.4.Description"`
    DataDisk_4_Device       string  `ArgName:"DataDisk.4.Device"`

    ClientToken             string
}

type CreateInstanceResult struct {
    GlobalResult

    InstanceId              string  `json:"InstanceId"`
}

func (self *Client) CreateInstance(args *CreateInstanceArgs) (result *CreateInstanceResult, errorResult *ErrorResult) {
    res := new(CreateInstanceResult)
    errRes := self.CallAPI("CreateInstance", args, res)

    return res, errRes
}

/**************************************************/

// Start `Instance`
type StartInstanceArgs struct {
    InstanceId  string  `json:"InstanceId"`
}

type StartInstanceResult struct {
    GlobalResult
}

func (self *Client) StartInstance(args *StartInstanceArgs) (result *StartInstanceResult, errorResult *ErrorResult) {
    res := new(StartInstanceResult)
    errRes := self.CallAPI("StartInstance", args, res)

    return res, errRes
}

/**************************************************/

// Stop `Instance`
type StopInstanceArgs struct {
    InstanceId  string  `json:"InstanceId"`
    ForceStop   string  `json:"ForceStop"`
}

type StopInstanceResult struct {
    GlobalResult
}

func (self *Client) StopInstance(args *StopInstanceArgs) (result *StopInstanceResult, errorResult *ErrorResult) {
    res := new(StopInstanceResult)
    errRes := self.CallAPI("StopInstance", args, res)

    return res, errRes
}

/**************************************************/

// Reboot `Instance`
type RebootInstanceArgs struct {
    InstanceId  string  `json:"InstanceId"`
    ForceStop   string  `json:"ForceStop"`
}

type RebootInstanceResult struct {
    GlobalResult
}

func (self *Client) RebootInstance(args *RebootInstanceArgs) (result *RebootInstanceResult, errorResult *ErrorResult) {
    res := new(RebootInstanceResult)
    errRes := self.CallAPI("RebootInstance", args, res)

    return res, errRes
}

/**************************************************/

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

/**************************************************/

// Describe `Instance` status
type DescribeInstanceStatusArgs struct {
    RegionId    string
    ZoneId      string
    PageNumber  string
    PageSize    string
}

type InstanceStatusItemType struct {
    InstanceId          string                      `json:"InstanceId"`
    Status              string                      `json:"Status"`
}

type InstanceStatusSetType struct {
    InstanceStatus      []InstanceStatusItemType    `json:InstanceStatus`
}

type DescribeInstanceStatusResult struct {
    GlobalResult

    TotalCount          int                         `json:"TotalCount"`
    PageNumber          int                         `json:"PageNumber"`
    PageSize            int                         `json:"PageSize"`
    InstanceStatuses    InstanceStatusSetType       `json:"InstanceStatuses"`
}

func (self *Client) DescribeInstanceStatus(args *DescribeInstanceStatusArgs) (result *DescribeInstanceStatusResult, errorResult *ErrorResult) {
    res := new(DescribeInstanceStatusResult)
    errRes := self.CallAPI("DescribeInstanceStatus", args, res)

    return res, errRes
}

/**************************************************/

// Describe `Instance` attribute
type DescribeInstanceAttributeArgs struct {
    InstanceId string
}

type OperationLocksType struct {
    LockReason              string                      `json:"LockReason"`
}

type OperationLocksType_Array struct {
    OperationLock           []OperationLocksType        `json:"OperationLock"`
}

type SecurityGroupIdSetType struct {
    SecurityGroupId         []string                    `json:"SecurityGroupId"`
}

type IpAddressSetType struct {
    IpAddress               []string                    `json:"IpAddress"`
}

type DescribeInstanceAttributeResult struct {
    GlobalResult

    InstanceId              string                      `json:"InstanceId"`
    InstanceName            string                      `json:"InstanceName"`
    Description             string                      `json:"Description"`
    ImageId                 string                      `json:"ImageId"`
    RegionId                string                      `json:"RegionId"`
    ZoneId                  string                      `json:"ZoneId"`
    ClusterId               string                      `json:"ClusterId"`
    InstanceType            string                      `json:"InstanceType"`
    HostName                string                      `json:"HostName"`
    Status                  string                      `json:"Status"`
    OperationLocks          OperationLocksType_Array    `json:"OperationLocks"`
    SecurityGroupIds        SecurityGroupIdSetType      `json:"SecurityGroupIds"`
    PublicIpAddress         IpAddressSetType            `json:"PublicIpAddress"`
    InnerIpAddress          IpAddressSetType            `json:"InnerIpAddress"`
    InternetMaxBandwidthIn  int                         `json:"InternetMaxBandwidthIn"`
    InternetMaxBandwidthOut int                         `json:"InternetMaxBandwidthOut"`
    InternetChargeType      string                      `json:"InternetChargeType"`
    CreationTime            string                      `json:"CreationTime"`

    //Note: Maybe API doc miss
    //VlanId                  string              `json:"VlanId"`
    //SerialNumber            string              `json:"SerialNumber"`   
}

func (self *Client) DescribeInstanceAttribute(args *DescribeInstanceAttributeArgs ) (result *DescribeInstanceAttributeResult, errorResult *ErrorResult) {
    res := new(DescribeInstanceAttributeResult)
    errRes := self.CallAPI("DescribeInstanceAttribute", args, res)

    return res, errRes
}

/**************************************************/

// Delete `Instance`
type DeleteInstanceArgs struct {
    InstanceId  string  `json:"InstanceId"`
}

type DeleteInstanceResult struct {
    GlobalResult
}

func (self *Client) DeleteInstance(args *DeleteInstanceArgs) (result *DeleteInstanceResult, errorResult *ErrorResult) {
    res := new(DeleteInstanceResult)
    errRes := self.CallAPI("DeleteInstance", args, res)

    return res, errRes
}

/**************************************************/

// Join `SecurityGroup`
type JoinSecurityGroupArgs struct {
    InstanceId      string  `json:"InstanceId"`
    SecurityGroupId string  `json:"SecurityGroupId"`
}

type JoinSecurityGroupResult struct {
    GlobalResult
}

func (self *Client) JoinSecurityGroup(args *JoinSecurityGroupArgs) (result *JoinSecurityGroupResult, errorResult *ErrorResult) {
    res := new(JoinSecurityGroupResult)
    errRes := self.CallAPI("JoinSecurityGroup", args, res)

    return res, errRes
}

/**************************************************/

// Leave `SecurityGroup`
type LeaveSecurityGroupArgs struct {
    InstanceId      string  `json:"InstanceId"`
    SecurityGroupId string  `json:"SecurityGroupIds"`
}

type LeaveSecurityGroupResult struct {
    GlobalResult
}

func (self *Client) LeaveSecurityGroup(args *LeaveSecurityGroupArgs) (result *LeaveSecurityGroupResult, errorResult *ErrorResult) {
    res := new(LeaveSecurityGroupResult)
    errRes := self.CallAPI("LeaveSecurityGroup", args, res)

    return res, errRes
}
