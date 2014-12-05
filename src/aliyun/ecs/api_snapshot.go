// API on Snapshot 

package ecs

import (
)

// Create `Snapshot`
type CreateSnapshotArgs struct {

}

type CreateSnapshotResult struct {
    GlobalResult
    // TODO
}

func (self *Client) CreateSnapshot(args *CreateSnapshotArgs) (result *CreateSnapshotResult, errorResult *ErrorResult) {
    res := new(CreateSnapshotResult)
    errRes := self.CallAPI("CreateSnapshot", args, res)

    return res, errRes
}



// Delete `Snapshot`
type DeleteSnapshotArgs struct {

}

type DeleteSnapshotResult struct {
    GlobalResult

}

func (self *Client) DeleteSnapshot(args *DeleteSnapshotArgs) (result *DeleteSnapshotResult, errorResult *ErrorResult) {
    res := new(DeleteSnapshotResult)
    errRes := self.CallAPI("DeleteSnapshot", args, res)

    return res, errRes
}

// Describe `Snapshot`s
type DescribeSnapshotsArgs struct {

}

type DescribeSnapshotsResult struct {
    GlobalResult
    // TODO
}

func (self *Client) DescribeSnapshots(args *DescribeSnapshotsArgs) (result *DescribeSnapshotsResult, errorResult *ErrorResult) {
    res := new(DescribeSnapshotsResult)
    errRes := self.CallAPI("DescribeSnapshots", args, res)

    return res, errRes
}

// Modify `AutoSnapshot` Policy
type ModifyAutoSnapshotPolicyArgs struct {

}

type ModifyAutoSnapshotPolicyResult struct {
    GlobalResult
    // TODO
}

func (self *Client) ModifyAutoSnapshotPolicy(args *ModifyAutoSnapshotPolicyArgs) (result *ModifyAutoSnapshotPolicyResult, errorResult *ErrorResult) {
    res := new(ModifyAutoSnapshotPolicyResult)
    errRes := self.CallAPI("ModifyAutoSnapshotPolicy", args, res)

    return res, errRes
}


// Describe `AutoSnapshot` Policy
type DescribeAutoSnapshotPolicyArgs struct {

}

type DescribeAutoSnapshotPolicyResult struct {
    GlobalResult
    // TODO
}

func (self *Client) DescribeAutoSnapshotPolicy(args *DescribeAutoSnapshotPolicyArgs) (result *DescribeAutoSnapshotPolicyResult, errorResult *ErrorResult) {
    res := new(DescribeAutoSnapshotPolicyResult)
    errRes := self.CallAPI("DescribeAutoSnapshotPolicy", args, res)

    return res, errRes
}