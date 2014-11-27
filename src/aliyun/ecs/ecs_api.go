package ecs

func (self *ECSClient) DescribeRegions() (result string, err error) {
    return self.CallAPI("DescribeRegions")
}