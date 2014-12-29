// API on Snapshot

package ecs

import ()

// Create `Snapshot`
type CreateSnapshotArgs struct {
	DiskId       string
	SnapshotName string
	Description  string
	ClientToken  string
}

type CreateSnapshotResult struct {
	GlobalResult

	SnapshotId string `json:"SnapshotId"`
}

func (self *Client) CreateSnapshot(args *CreateSnapshotArgs) (result *CreateSnapshotResult, errorResult *ErrorResult) {
	res := new(CreateSnapshotResult)
	errRes := self.CallAPI("CreateSnapshot", args, res)

	return res, errRes
}

/**************************************************/

// Delete `Snapshot`
type DeleteSnapshotArgs struct {
	SnapshotId string
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

type SnapshotType struct {
	SnapshotId     string `json:"SnapshotId"`
	SnapshotName   string `json:"SnapshotName"`
	Description    string `json:"Description"`
	Progress       string `json:"Progress"`
	SourceDiskId   string `json:"SourceDiskId"`
	SourceDiskSize int    `json:"SourceDiskSize"`
	SourceDiskType string `json:"SourceDiskType"`
	ProductCode    string `json:"ProductCode"`
	CreationTime   string `json:"CreationTime"`
}

type SnapshotType_Array struct {
	SnapshotResource []SnapshotType `json:"SnapshotResource"`
}

type DescribeSnapshotsResult struct {
	GlobalResult

	RegionId   string             `json:"RegionId"`
	TotalCount int                `json:"TotalCount"`
	PageNumber int                `json:"PageNumber"`
	PageSize   int                `json:"PageSize"`
	Snapshots  SnapshotType_Array `json:"Snapshots"`
}

func (self *Client) DescribeSnapshots(args *DescribeSnapshotsArgs) (result *DescribeSnapshotsResult, errorResult *ErrorResult) {
	res := new(DescribeSnapshotsResult)
	errRes := self.CallAPI("DescribeSnapshots", args, res)

	return res, errRes
}

/**************************************************/

// Modify `AutoSnapshot` Policy
type ModifyAutoSnapshotPolicyArgs struct {
	SystemDiskPolicyEnabled           string
	SystemDiskPolicyTimePeriod        int
	SystemDiskPolicyRetentionDays     int
	SystemDiskPolicyRetentionLastWeek string
	DataDiskPolicyEnabled             string
	DataDiskPolicyTimePeriod          int
	DataDiskPolicyRetentionDays       int
	DataDiskPolicyRetentionLastWeek   string
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

type AutoSnapshotPolicyType struct {
	SystemDiskPolicyEnabled           string `json:"SystemDiskPolicyEnabled"`
	SystemDiskPolicyTimePeriod        int    `json:"SystemDiskPolicyTimePeriod"`
	SystemDiskPolicyRetentionDays     int    `json:"SystemDiskPolicyRetentionDays"`
	SystemDiskPolicyRetentionLastWeek string `json:"SystemDiskPolicyRetentionLastWeek"`
	DataDiskPolicyEnabled             string `json:"DataDiskPolicyEnabled"`
	DataDiskPolicyTimePeriod          int    `json:"DataDiskPolicyTimePeriod"`
	DataDiskPolicyRetentionDays       int    `json:"DataDiskPolicyRetentionDays"`
	DataDiskPolicyRetentionLastWeek   string `json:"DataDiskPolicyRetentionLastWeek"`
}

type AutoSnapshotExecutionStatusType struct {
	SystemDiskExcutionStatus string `json:"SystemDiskExcutionStatus"`
	DataDiskExcutionStatus   string `json:"DataDiskExcutionStatus"`
}

type DescribeAutoSnapshotPolicyResult struct {
	GlobalResult

	AutoSnapshotPolicy         AutoSnapshotPolicyType          `json:"AutoSnapshotPolicy"`
	AutoSnapshotExcutionStatus AutoSnapshotExecutionStatusType `json:"AutoSnapshotExcutionStatus"`
}

func (self *Client) DescribeAutoSnapshotPolicy(args *DescribeAutoSnapshotPolicyArgs) (result *DescribeAutoSnapshotPolicyResult, errorResult *ErrorResult) {
	res := new(DescribeAutoSnapshotPolicyResult)
	errRes := self.CallAPI("DescribeAutoSnapshotPolicy", args, res)

	return res, errRes
}
