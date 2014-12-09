// API on Security

package ecs

import (
)

// Create `SecurityGroup`
type CreateSecurityGroupArgs struct {

}

type CreateSecurityGroupResult struct {
    GlobalResult
    // TODO
}

func (self *Client) CreateSecurityGroup(args *CreateSecurityGroupArgs) (result *CreateSecurityGroupResult, errorResult *ErrorResult) {
    res := new(CreateSecurityGroupResult)
    errRes := self.CallAPI("CreateSecurityGroup", args, res)

    return res, errRes
}

/**************************************************/

// Authorize `SecurityGroup`
type AuthorizeSecurityGroupArgs struct {

}

type AuthorizeSecurityGroupResult struct {
    GlobalResult
    // TODO
}

func (self *Client) AuthorizeSecurityGroup(args *AuthorizeSecurityGroupArgs) (result *AuthorizeSecurityGroupResult, errorResult *ErrorResult) {
    res := new(AuthorizeSecurityGroupResult)
    errRes := self.CallAPI("AuthorizeSecurityGroup", args, res)

    return res, errRes
}

/**************************************************/

// Describe `SecurityGroup` Attribute
type DescribeSecurityGroupAttributeArgs struct {

}

type DescribeSecurityGroupAttributeResult struct {
    GlobalResult
    // TODO
}

func (self *Client) DescribeSecurityGroupAttribute(args *DescribeSecurityGroupAttributeArgs) (result *DescribeSecurityGroupAttributeResult, errorResult *ErrorResult) {
    res := new(DescribeSecurityGroupAttributeResult)
    errRes := self.CallAPI("DescribeSecurityGroupAttribute", args, res)

    return res, errRes
}

/**************************************************/

// Describe `SecurityGroups`
type DescribeSecurityGroupsArgs struct {

}

type DescribeSecurityGroupsResult struct {
    GlobalResult
    // TODO
}

func (self *Client) DescribeSecurityGroups(args *DescribeSecurityGroupsArgs) (result *DescribeSecurityGroupsResult, errorResult *ErrorResult) {
    res := new(DescribeSecurityGroupsResult)
    errRes := self.CallAPI("DescribeSecurityGroups", args, res)

    return res, errRes
}

/**************************************************/

// Revoke `SecurityGroup`
type RevokeSecurityGroupArgs struct {

}

type RevokeSecurityGroupResult struct {
    GlobalResult
    // TODO
}

func (self *Client) RevokeSecurityGroup(args *RevokeSecurityGroupArgs) (result *RevokeSecurityGroupResult, errorResult *ErrorResult) {
    res := new(RevokeSecurityGroupResult)
    errRes := self.CallAPI("RevokeSecurityGroup", args, res)

    return res, errRes
}

/**************************************************/

// Delete `SecurityGroup`
type DeleteSecurityGroupArgs struct {

}

type DeleteSecurityGroupResult struct {
    GlobalResult
    // TODO
}

func (self *Client) DeleteSecurityGroup(args *DeleteSecurityGroupArgs) (result *DeleteSecurityGroupResult, errorResult *ErrorResult) {
    res := new(DeleteSecurityGroupResult)
    errRes := self.CallAPI("DeleteSecurityGroup", args, res)

    return res, errRes
}