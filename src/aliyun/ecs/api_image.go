// API on Images

package ecs

import (
)

// Describe `Image`s
type DescribeImagesArgs struct {
    RegionId        string
    PageNumber      int
    PageSize        int
    ImageId         string
    SnapshotId      string
    ImageName       string
    ImageOwnerAlias string
}

type DiskDeviceMapping_Type struct {
    SnapshotId          string                      `json:"SnapshotId"`
    //Note: According API doc, `Size` Field is string-type.
    Size                int                         `json:"Size"`
    Device              string                      `json:"Device"`
}

type DiskDeviceMapping_Array struct {
    DiskDeviceMapping   []DiskDeviceMapping_Type    `json:"DiskDeviceMapping"`
}

type ImageType struct {
    ImageId             string                      `json:"ImageId"`
    ImageVersion        string                      `json:"ImageVersion"`
    Architecture        string                      `json:"Architecture"`
    ImageName           string                      `json:"ImageName"`
    Description         string                      `json:"Description"`
    Size                int                         `json:"Size"`
    ImageOwnerAlias     string                      `json:"ImageOwnerAlias"`
    OSName              string                      `json:"OSName"`
    DiskDeviceMappings  DiskDeviceMapping_Array     `json:"DiskDeviceMappings"`
    ProductCode         string                      `json:"ProductCode"`
    IsSubscribed        string                      `json:"IsSubscribed"`
    CreationTime        string                      `json:"CreationTime"`
}

type ImageType_Array struct {
    Images              []ImageType                 `json:"Images"`
}

type DescribeImagesResult struct {
    GlobalResult

    RegionId            string                      `json:"RegionId"`
    TotalCount          int                         `json:"TotalCount"`
    PageNumber          int                         `json:"PageNumber"`
    PageSize            int                         `json:"PageSize"`
    Images              ImageType_Array             `json:"Images"`
}

func (self *Client) DescribeImages(args *DescribeImagesArgs) (result *DescribeImagesResult, errorResult *ErrorResult) {
    res := new(DescribeImagesResult)
    errRes := self.CallAPI("DescribeImages", args, res)

    return res, errRes
}

/**************************************************/

// Create `Image`
type CreateImageArgs struct {
    RegionId        string
    SnapshotId      string
    ImageName       string
    ImageVersion    string
    Description     string
    ClientToken     string
}

type CreateImageResult struct {
    GlobalResult

    ImageId         string  `json:"ImageId"`
}

func (self *Client) CreateImage(args *CreateImageArgs) (result *CreateImageResult, errorResult *ErrorResult) {
    res := new(CreateImageResult)
    errRes := self.CallAPI("CreateImage", args, res)

    return res, errRes
}

/**************************************************/

// DeleteImage
type DeleteImageArgs struct {
    RegionId    string
    ImageId     string
}

type DeleteImageResult struct {
    GlobalResult
}

func (self *Client) DeleteImage(args *DeleteImageArgs) (result *DeleteImageResult, errorResult *ErrorResult) {
    res := new(DeleteImageResult)
    errRes := self.CallAPI("DeleteImage", args, res)

    return res, errRes
}
