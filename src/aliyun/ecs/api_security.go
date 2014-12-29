// API on Security

package ecs

import ()

// Create `SecurityGroup`
type CreateSecurityGroupArgs struct {
	RegionId    string
	Description string
	ClientToken string
}

type CreateSecurityGroupResult struct {
	GlobalResult

	SecurityGroupId string `json:"SecurityGroupId"`
}

func (self *Client) CreateSecurityGroup(args *CreateSecurityGroupArgs) (result *CreateSecurityGroupResult, errorResult *ErrorResult) {
	res := new(CreateSecurityGroupResult)
	errRes := self.CallAPI("CreateSecurityGroup", args, res)

	return res, errRes
}

/**************************************************/

// Authorize `SecurityGroup`
type AuthorizeSecurityGroupArgs struct {
	SecurityGroupId string
	RegionId        string
	IpProtocol      string
	PortRange       string
	SourceGroupId   string
	SourceCidrIp    string
	Policy          string
	NicType         string
}

type AuthorizeSecurityGroupResult struct {
	GlobalResult
}

func (self *Client) AuthorizeSecurityGroup(args *AuthorizeSecurityGroupArgs) (result *AuthorizeSecurityGroupResult, errorResult *ErrorResult) {
	res := new(AuthorizeSecurityGroupResult)
	errRes := self.CallAPI("AuthorizeSecurityGroup", args, res)

	return res, errRes
}

/**************************************************/

// Describe `SecurityGroup` Attribute
type DescribeSecurityGroupAttributeArgs struct {
	SecurityGroupId string
	RegionId        string
	NicType         string
}

type PermissionType struct {
	IpProtocol    string `json:"IpProtocol"`
	PortRange     string `json:"PortRange"`
	SourceCidrIp  string `json:"SourceCidrIp"`
	SourceGroupId string `json:"SourceGroupId"`
	Policy        string `json:"Policy"`
	NicType       string `json:"NicType"`
}

type PermissionSetType struct {
	Permission []PermissionType `json:"Permission"`
}

type DescribeSecurityGroupAttributeResult struct {
	GlobalResult

	SecurityGroupId string            `json:"SecurityGroupId"`
	RegionId        string            `json:"RegionId"`
	Description     string            `json:"Description"`
	Permissions     PermissionSetType `json:"Permissions"`
}

func (self *Client) DescribeSecurityGroupAttribute(args *DescribeSecurityGroupAttributeArgs) (result *DescribeSecurityGroupAttributeResult, errorResult *ErrorResult) {
	res := new(DescribeSecurityGroupAttributeResult)
	errRes := self.CallAPI("DescribeSecurityGroupAttribute", args, res)

	return res, errRes
}

/**************************************************/

// Describe `SecurityGroups`
type DescribeSecurityGroupsArgs struct {
	RegionId   string
	PageNumber int
	PageSize   int
}

type SecurityGroupItemType struct {
	SecurityGroupId string `json:"SecurityGroupId"`
	Description     string `json:"Description"`
}

type SecurityGroupSetType struct {
	SecurityGroup []SecurityGroupItemType `json:"SecurityGroup"`
}

type DescribeSecurityGroupsResult struct {
	GlobalResult

	TotalCount     int                  `json:"TotalCount"`
	PageNumber     int                  `json:"PageNumber"`
	PageSize       int                  `json:"PageSize"`
	RegionId       string               `json:"RegionId"`
	SecurityGroups SecurityGroupSetType `json:"SecurityGroups"`
}

func (self *Client) DescribeSecurityGroups(args *DescribeSecurityGroupsArgs) (result *DescribeSecurityGroupsResult, errorResult *ErrorResult) {
	res := new(DescribeSecurityGroupsResult)
	errRes := self.CallAPI("DescribeSecurityGroups", args, res)

	return res, errRes
}

/**************************************************/

// Revoke `SecurityGroup`
type RevokeSecurityGroupArgs struct {
	SecurityGroupId string
	RegionId        string
	IpProtocol      string
	PortRange       string
	SourceGroupId   string
	SourceCidrIp    string
	Policy          string
	NicType         string
}

type RevokeSecurityGroupResult struct {
	GlobalResult
}

func (self *Client) RevokeSecurityGroup(args *RevokeSecurityGroupArgs) (result *RevokeSecurityGroupResult, errorResult *ErrorResult) {
	res := new(RevokeSecurityGroupResult)
	errRes := self.CallAPI("RevokeSecurityGroup", args, res)

	return res, errRes
}

/**************************************************/

// Delete `SecurityGroup`
type DeleteSecurityGroupArgs struct {
	SecurityGroupId string
	RegionId        string
}

type DeleteSecurityGroupResult struct {
	GlobalResult
}

func (self *Client) DeleteSecurityGroup(args *DeleteSecurityGroupArgs) (result *DeleteSecurityGroupResult, errorResult *ErrorResult) {
	res := new(DeleteSecurityGroupResult)
	errRes := self.CallAPI("DeleteSecurityGroup", args, res)

	return res, errRes
}
