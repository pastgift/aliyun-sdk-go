// API on Snapshot

package ecs

import (
)

// Create `Snapshot`
type CreateSnapshotArgs struct {
    DiskId          string
    SnapshotName    string
    Description     string
    ClientToken     string
}

type CreateSnapshotResult struct {
    GlobalResult

    SnapshotId      string  `json:"SnapshotId"`
}

func (self *Client) CreateSnapshot(args *CreateSnapshotArgs) (result *CreateSnapshotResult, errorResult *ErrorResult) {
    res := new(CreateSnapshotResult)
    errRes := self.CallAPI("CreateSnapshot", args, res)

    return res, errRes
}

/**************************************************/

// Delete `Snapshot`
type DeleteSnapshotArgs struct {
    SnapshotId  string
}

type DeleteSnapshotResult struct {
    GlobalResult
}

func (self *Client) DeleteSnapshot(args *DeleteSnapshotArgs) (result *DeleteSnapshotResult, errorResult *ErrorResult) {
    res := new(DeleteSnapshotResult)
    errRes := self.CallAPI("DeleteSnapshot", args, res)

    return res, errRes
}

/**************************************************/

// Describe `Snapshot`s
type DescribeSnapshotsArgs struct {
    RegionId    string
    InstanceId  string
    DiskId      string
    SnapshotIds string
    PageNumber  int
    PageSize    int
}

type SnapshotResourceST struct {
    CreationTime        string                  `json:"CreationTime"`
    Description         string                  `json:"Description"`
    ProductCode         string                  `json:"ProductCode"`
    Progress            string                  `json:"Progress"`
    SnapshotId          string                  `json:"SnapshotId"`
    SnapshotName        string                  `json:"SnapshotName"`
    SourceDiskId        string                  `json:"SourceDiskId"`
    SourceDiskSize      string                  `json:"SourceDiskSize"`
    SourceDiskType      string                  `json:"SourceDiskType"`
}

type SnapshotsST struct {
    SnapshotResource    []SnapshotResourceST    `json:"SnapshotResource"`
}

type DescribeSnapshotsResult struct {
    GlobalResult

    RegionId            string                  `json:"RegionId"`
    TotalCount          int                     `json:"TotalCount"`
    PageNumber          int                     `json:"PageNumber"`
    PageSize            int                     `json:"PageSize"`
    Snapshots           SnapshotsST             `json:"Snapshots"`
}

func (self *Client) DescribeSnapshots(args *DescribeSnapshotsArgs) (result *DescribeSnapshotsResult, errorResult *ErrorResult) {
    res := new(DescribeSnapshotsResult)
    errRes := self.CallAPI("DescribeSnapshots", args, res)

    return res, errRes
}

/**************************************************/

// Modify `AutoSnapshot` Policy
type ModifyAutoSnapshotPolicyArgs struct {
    SystemDiskPolicyEnabled             string
    SystemDiskPolicyTimePeriod          int
    SystemDiskPolicyRetentionDays       int
    SystemDiskPolicyRetentionLastWeek   string
    DataDiskPolicyEnabled               string
    DataDiskPolicyTimePeriod            int
    DataDiskPolicyRetentionDays         int
    DataDiskPolicyRetentionLastWeek     string
}

type ModifyAutoSnapshotPolicyResult struct {
    GlobalResult
}

func (self *Client) ModifyAutoSnapshotPolicy(args *ModifyAutoSnapshotPolicyArgs) (result *ModifyAutoSnapshotPolicyResult, errorResult *ErrorResult) {
    res := new(ModifyAutoSnapshotPolicyResult)
    errRes := self.CallAPI("ModifyAutoSnapshotPolicy", args, res)

    return res, errRes
}

/**************************************************/

// Describe `AutoSnapshot` Policy
type DescribeAutoSnapshotPolicyArgs struct {
    // NO ARGUMENT NEEDED
}

type AutoSnapshotExcutionStatusST struct {
    DataDiskExcutionStatus              string                          `json:"DataDiskExcutionStatus"`
    SystemDiskExcutionStatus            string                          `json:"SystemDiskExcutionStatus"`
}

type AutoSnapshotPolicyST struct {
    DataDiskPolicyEnabled               string                          `json:"DataDiskPolicyEnabled"`
    DataDiskPolicyRetentionDays         string                          `json:"DataDiskPolicyRetentionDays"`
    DataDiskPolicyRetentionLastWeek     string                          `json:"DataDiskPolicyRetentionLastWeek"`
    DataDiskPolicyTimePeriod            string                          `json:"DataDiskPolicyTimePeriod"`
    SystemDiskPolicyEnabled             string                          `json:"SystemDiskPolicyEnabled"`
    SystemDiskPolicyRetentionDays       string                          `json:"SystemDiskPolicyRetentionDays"`
    SystemDiskPolicyRetentionLastWeek   string                          `json:"SystemDiskPolicyRetentionLastWeek"`
    SystemDiskPolicyTimePeriod          string                          `json:"SystemDiskPolicyTimePeriod"`
}

type DescribeAutoSnapshotPolicyResult struct {
    GlobalResult

    AutoSnapshotPolicy                  AutoSnapshotPolicyST            `json:"AutoSnapshotPolicy"`
    AutoSnapshotExcutionStatus          AutoSnapshotExcutionStatusST    `json:"AutoSnapshotExcutionStatus"`
}

func (self *Client) DescribeAutoSnapshotPolicy(args *DescribeAutoSnapshotPolicyArgs) (result *DescribeAutoSnapshotPolicyResult, errorResult *ErrorResult) {
    res := new(DescribeAutoSnapshotPolicyResult)
    errRes := self.CallAPI("DescribeAutoSnapshotPolicy", args, res)

    return res, errRes
}
