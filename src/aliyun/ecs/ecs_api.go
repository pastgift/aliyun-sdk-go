package ecs

func (self *ECSClient) DescribeRegions() (result string, err error) {
    ecs_req := NewECSRequest("DescribeRegions")
    return self.Do(ecs_req)
}