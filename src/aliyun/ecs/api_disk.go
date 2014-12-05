// API on Disk 

package ecs

import (
)

// Create `Disk`
type CreateDiskArgs struct {

}

type CreateDiskResult struct {
    GlobalResult
    // TODO
}

func (self *Client) CreateDisk(args *CreateDiskArgs) (result *CreateDiskResult, errorResult *ErrorResult) {
    res := new(CreateDiskResult)
    errRes := self.CallAPI("CreateDisk", args, res)

    return res, errRes
}

// Describe `Disk`s
const (
    DISK_TYPE_ALL       = "all"
    DISK_TYPE_SYSTEM    = "system"
    DISK_TYPE_DATA      = "data"

    DISK_CAT_ALL        = "all"
    DISK_CAT_CLOUD      = "cloud"
    DISK_CAT_EPHEMERAL  = "ephemeral"
    
    DISK_STATUS_IN_USE      = "In_use"
    DISK_STATUS_AVAILABLE   = "Available"
    DISK_STATUS_ATTACHING   = "Attaching"
    DISK_STATUS_DETACHING   = "Detaching"
    DISK_STATUS_CREATING    = "Creating"
    DISK_STATUS_REINITING   = "ReIniting"
    DISK_STATUS_ALL         = "All"
    
    DISK_PORTABLE_TRUE  = "True"
    DISK_PORTABLE_FALSE = "False"

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

type DescribeDisksResult struct {
    GlobalResult

    RegionId    string  `json:"RegionId"`
    TotalCount  int     `json:"TotalCount"`
    PageNumber  int     `json:"PageNumber"`
    PageSize    int     `json:"PageSize"`
    Disks       struct {
        Disk        []struct {
            AttachedTime        string  `json:"AttachedTime"`
            Category            string  `json:"Category"`
            CreationTime        string  `json:"CreationTime"`
            DeleteAutoSnapshot  string  `json:"DeleteAutoSnapshot"`
            DeleteWithInstance  string  `json:"DeleteWithInstance"`
            Description         string  `json:"Description"`
            DetachedTime        string  `json:"DetachedTime"`
            Device              string  `json:"Device"`
            DiskId              string  `json:"DiskId"`
            DiskName            string  `json:"DiskName"`
            ImageId             string  `json:"ImageId"`
            InstanceId          string  `json:"InstanceId"`
            OperationLocks      struct {
                OperationLock       []string    `json:"OperationLock"`    
            }                           `json:"OperationLocks"`
        }                   `json:"Disk"`
    }                   `json:"Disks"`
}

func (self *Client) DescribeDisks(args *DescribeDisksArgs) (result *DescribeDisksResult, errorResult *ErrorResult) {
    res := new(DescribeDisksResult)
    errRes := self.CallAPI("DescribeDisks", args, res)

    return res, errRes
}

// Attach `Disk`
type AttachDiskArgs struct {

}

type AttachDiskResult struct {
    GlobalResult
    // TODO
}

func (self *Client) AttachDisk(args *AttachDiskArgs) (result *AttachDiskResult, errorResult *ErrorResult) {
    res := new(AttachDiskResult)
    errRes := self.CallAPI("AttachDisk", args, res)

    return res, errRes
}

// Detach `Disk`
type DetachDiskArgs struct {

}

type DetachDiskResult struct {
    GlobalResult
    // TODO
}

func (self *Client) DetachDisk(args *DetachDiskArgs) (result *DetachDiskResult, errorResult *ErrorResult) {
    res := new(DetachDiskResult)
    errRes := self.CallAPI("DetachDisk", args, res)

    return res, errRes
}

// Modify `Disk` Attribute
type ModifyDiskAttributeArgs struct {

}

type ModifyDiskAttributeResult struct {
    GlobalResult
    // TODO
}

func (self *Client) ModifyDiskAttribute(args *ModifyDiskAttributeArgs) (result *ModifyDiskAttributeResult, errorResult *ErrorResult) {
    res := new(ModifyDiskAttributeResult)
    errRes := self.CallAPI("ModifyDiskAttribute", args, res)

    return res, errRes
}

// Delete `Disk`
type DeleteDiskArgs struct {

}

type DeleteDiskResult struct {
    GlobalResult
    // TODO
}

func (self *Client) DeleteDisk(args *DeleteDiskArgs) (result *DeleteDiskResult, errorResult *ErrorResult) {
    res := new(DeleteDiskResult)
    errRes := self.CallAPI("DeleteDisk", args, res)

    return res, errRes
}

// ReInit `Disk`
type ReInitDiskArgs struct {

}

type ReInitDiskResult struct {
    GlobalResult
    // TODO
}

func (self *Client) ReInitDisk(args *ReInitDiskArgs) (result *ReInitDiskResult, errorResult *ErrorResult) {
    res := new(ReInitDiskResult)
    errRes := self.CallAPI("ReInitDisk", args, res)

    return res, errRes
}

// Reset `Disk`
type ResetDiskArgs struct {

}

type ResetDiskResult struct {
    GlobalResult
    // TODO
}

func (self *Client) ResetDisk(args *ResetDiskArgs) (result *ResetDiskResult, errorResult *ErrorResult) {
    res := new(ResetDiskResult)
    errRes := self.CallAPI("ResetDisk", args, res)

    return res, errRes
}

// Replace `SystemDisk`
type ReplaceSystemDiskArgs struct {

}

type ReplaceSystemDiskResult struct {
    GlobalResult
    // TODO
}

func (self *Client) ReplaceSystemDisk(args *ReplaceSystemDiskArgs) (result *ReplaceSystemDiskResult, errorResult *ErrorResult) {
    res := new(ReplaceSystemDiskResult)
    errRes := self.CallAPI("ReplaceSystemDisk", args, res)

    return res, errRes
}