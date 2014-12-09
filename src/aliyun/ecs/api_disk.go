// API on Disk

package ecs

import (
)

// Create `Disk`
type CreateDiskArgs struct {
    RegionId    string
    ZoneId      string
    DiskName    string
    Description string
    Size        int
    SnapshotId  string
    ClientToken string
}

type CreateDiskResult struct {
    GlobalResult

    DiskId      string `json:"DiskId"`
}

func (self *Client) CreateDisk(args *CreateDiskArgs) (result *CreateDiskResult, errorResult *ErrorResult) {
    res := new(CreateDiskResult)
    errRes := self.CallAPI("CreateDisk", args, res)

    return res, errRes
}

/**************************************************/

// Describe `Disk`s
const (
    DISK_TYPE_ALL                   = "all"
    DISK_TYPE_SYSTEM                = "system"
    DISK_TYPE_DATA                  = "data"

    DISK_CAT_ALL                    = "all"
    DISK_CAT_CLOUD                  = "cloud"
    DISK_CAT_EPHEMERAL              = "ephemeral"

    DISK_STATUS_IN_USE              = "In_use"
    DISK_STATUS_AVAILABLE           = "Available"
    DISK_STATUS_ATTACHING           = "Attaching"
    DISK_STATUS_DETACHING           = "Detaching"
    DISK_STATUS_CREATING            = "Creating"
    DISK_STATUS_REINITING           = "ReIniting"
    DISK_STATUS_ALL                 = "All"

    DISK_PORTABLE_TRUE              = "True"
    DISK_PORTABLE_FALSE             = "False"

    DISK_DELETE_WITH_INSTANCE_TRUE  = "True"
    DISK_DELETE_WITH_INSTANCE_FALSE = "False"

    DISK_DELETE_AUTO_SNAPSHOT_TRUE  = "True"
    DISK_DELETE_AUTO_SNAPSHOT_FALSE = "False"
)

type DescribeDisksArgs struct {
    RegionId            string
    ZoneId              string
    DiskIds             string
    InstanceId          string
    DiskType            string
    Category            string
    Status              string
    SnapshotId          string
    Portable            string
    DeleteWithInstance  string
    DeleteAutoSnapshot  string
    PageNumber          int
    PageSize            int
}

type OperationLocksST struct {
    OperationLock       []string            `json:"OperationLock"`
}

type DiskST struct {
    AttachedTime        string              `json:"AttachedTime"`
    Category            string              `json:"Category"`
    CreationTime        string              `json:"CreationTime"`
    DeleteAutoSnapshot  string              `json:"DeleteAutoSnapshot"`
    DeleteWithInstance  string              `json:"DeleteWithInstance"`
    Description         string              `json:"Description"`
    DetachedTime        string              `json:"DetachedTime"`
    Device              string              `json:"Device"`
    DiskId              string              `json:"DiskId"`
    DiskName            string              `json:"DiskName"`
    ImageId             string              `json:"ImageId"`
    InstanceId          string              `json:"InstanceId"`
    OperationLocks      OperationLocksST    `json:"OperationLocks"`
}

type DisksST struct {
    Disk                []DiskST            `json:"Disk"`
}

type DescribeDisksResult struct {
    GlobalResult

    RegionId            string              `json:"RegionId"`
    TotalCount          int                 `json:"TotalCount"`
    PageNumber          int                 `json:"PageNumber"`
    PageSize            int                 `json:"PageSize"`
    Disks               DisksST             `json:"Disks"`
}

func (self *Client) DescribeDisks(args *DescribeDisksArgs) (result *DescribeDisksResult, errorResult *ErrorResult) {
    res := new(DescribeDisksResult)
    errRes := self.CallAPI("DescribeDisks", args, res)

    return res, errRes
}

/**************************************************/

// Attach `Disk`
type AttachDiskArgs struct {
    InstanceId          string
    DiskId              string
    Device              string
    DeleteWithInstance  string
}

type AttachDiskResult struct {
    GlobalResult
}

func (self *Client) AttachDisk(args *AttachDiskArgs) (result *AttachDiskResult, errorResult *ErrorResult) {
    res := new(AttachDiskResult)
    errRes := self.CallAPI("AttachDisk", args, res)

    return res, errRes
}

/**************************************************/

// Detach `Disk`
type DetachDiskArgs struct {
    InstanceId  string
    DiskId      string
}

type DetachDiskResult struct {
    GlobalResult
}

func (self *Client) DetachDisk(args *DetachDiskArgs) (result *DetachDiskResult, errorResult *ErrorResult) {
    res := new(DetachDiskResult)
    errRes := self.CallAPI("DetachDisk", args, res)

    return res, errRes
}

/**************************************************/

// Modify `Disk` Attribute
type ModifyDiskAttributeArgs struct {
    DiskId              string
    DiskName            string
    Description         string
    DeleteWithInstance  string
}

type ModifyDiskAttributeResult struct {
    GlobalResult
}

func (self *Client) ModifyDiskAttribute(args *ModifyDiskAttributeArgs) (result *ModifyDiskAttributeResult, errorResult *ErrorResult) {
    res := new(ModifyDiskAttributeResult)
    errRes := self.CallAPI("ModifyDiskAttribute", args, res)

    return res, errRes
}

/**************************************************/

// Delete `Disk`
type DeleteDiskArgs struct {
    DiskId  string
}

type DeleteDiskResult struct {
    GlobalResult
}

func (self *Client) DeleteDisk(args *DeleteDiskArgs) (result *DeleteDiskResult, errorResult *ErrorResult) {
    res := new(DeleteDiskResult)
    errRes := self.CallAPI("DeleteDisk", args, res)

    return res, errRes
}

/**************************************************/

// ReInit `Disk`
type ReInitDiskArgs struct {
    DiskId  string
}

type ReInitDiskResult struct {
    GlobalResult
}

func (self *Client) ReInitDisk(args *ReInitDiskArgs) (result *ReInitDiskResult, errorResult *ErrorResult) {
    res := new(ReInitDiskResult)
    errRes := self.CallAPI("ReInitDisk", args, res)

    return res, errRes
}

/**************************************************/

// Reset `Disk`
type ResetDiskArgs struct {
    DiskId      string
    SnapshotId  string
}

type ResetDiskResult struct {
    GlobalResult
}

func (self *Client) ResetDisk(args *ResetDiskArgs) (result *ResetDiskResult, errorResult *ErrorResult) {
    res := new(ResetDiskResult)
    errRes := self.CallAPI("ResetDisk", args, res)

    return res, errRes
}

/**************************************************/

// Replace `SystemDisk`
type ReplaceSystemDiskArgs struct {
    InstanceId  string
    ImageId     string
    ClientToken string
}

type ReplaceSystemDiskResult struct {
    GlobalResult

    DiskId      string `json:"DiskId"`
}

func (self *Client) ReplaceSystemDisk(args *ReplaceSystemDiskArgs) (result *ReplaceSystemDiskResult, errorResult *ErrorResult) {
    res := new(ReplaceSystemDiskResult)
    errRes := self.CallAPI("ReplaceSystemDisk", args, res)

    return res, errRes
}
