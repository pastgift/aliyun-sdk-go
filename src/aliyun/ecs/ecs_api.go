package ecs

func (self *ECSClient) DescribeRegions() (result string, err error) {
    return self.Do(NewECSRequest("DescribeRegions"))
}