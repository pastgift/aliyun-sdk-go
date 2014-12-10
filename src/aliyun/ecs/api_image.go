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

type DiskDeviceMappingST struct {
    Device              string                  `json:"Device"`
    Size                string                  `json:"Size"`
    SnapshotId          string                  `json:"SnapshotId"`
}

type DiskDeviceMappingsST struct {
    DiskDeviceMapping   []DiskDeviceMappingST   `json:"DiskDeviceMapping"`
}

type ImageST struct {
    Architecture        string                  `json:"Architecture"`
    CreationTime        string                  `json:"CreationTime"`
    Description         string                  `json:"Description"`
    DiskDeviceMappings  DiskDeviceMappingsST    `json:"DiskDeviceMappings"`
    ImageId             string                  `json:"ImageId"`
    ImageName           string                  `json:"ImageName"`
    ImageOwnerAlias     string                  `json:"ImageOwnerAlias"`
    ImageVersion        string                  `json:"ImageVersion"`
    IsSubscribed        string                  `json:"IsSubscribed"`
    OSName              string                  `json:"OSName"`
    ProductCode         string                  `json:"ProductCode"`
    Size                string                  `json:"Size"`
}

type ImagesST struct {
    Image               []ImageST               `json:"Image"`
}

type DescribeImagesResult struct {
    GlobalResult

    RegionId            string                  `json:"RegionId"`
    TotalCount          int                     `json:"TotalCount"`
    PageNumber          int                     `json:"PageNumber"`
    PageSize            int                     `json:"PageSize"`
    Images              ImagesST                `json:"Images"`
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
    ClientToken     strings
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
