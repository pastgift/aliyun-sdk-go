// API on Images

package ecs

import (
)

// Describe `Image`s
type DescribeImagesArgs struct {

}

type DescribeImagesResult struct {
    GlobalResult
    // TODO
}

func (self *Client) DescribeImages(args *DescribeImagesArgs) (result *DescribeImagesResult, errorResult *ErrorResult) {
    res := new(DescribeImagesResult)
    errRes := self.CallAPI("DescribeImages", args, res)

    return res, errRes
}

// Create `Image`
type CreateImageArgs struct {

}

type CreateImageResult struct {
    GlobalResult
    // TODO
}

func (self *Client) CreateImage(args *CreateImageArgs) (result *CreateImageResult, errorResult *ErrorResult) {
    res := new(CreateImageResult)
    errRes := self.CallAPI("CreateImage", args, res)

    return res, errRes
}

// DeleteImage
type DeleteImageArgs struct {

}

type DeleteImageResult struct {
    GlobalResult
    // TODO
}

func (self *Client) DeleteImage(args *DeleteImageArgs) (result *DeleteImageResult, errorResult *ErrorResult) {
    res := new(DeleteImageResult)
    errRes := self.CallAPI("DeleteImage", args, res)

    return res, errRes
}