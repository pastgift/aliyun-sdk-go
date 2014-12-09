// API on Monitor

package ecs

import (
)

// Describe `Instance` `Monitor` Data
type DescribeInstanceMonitorDataArgs struct {
    InstanceId  string
    StartTime   string
    EndTime     string
    Period      string
}

type InstanceMonitorDataST struct {
    InstanceId          string                  `json:"InstanceId"`
    CPU                 int                     `json:"CPU"`
    Memory              int                     `json:"Memory"`
    IntranetRX          int                     `json:"IntranetRX"`
    IntranetTX          int                     `json:"IntranetTX"`
    IntranetFlow        int                     `json:"IntranetFlow"`
    IntranetBandwidth   int                     `json:"IntranetBandwidth"`
    InternetRX          int                     `json:"InternetRX"`
    InternetTX          int                     `json:"InternetTX"`
    InternetFlow        int                     `json:"InternetFlow"`
    InternetBandwidth   int                     `json:"InternetBandwidth"`
    IOPSRead            int                     `json:"IOPSRead"`
    IOPSWrite           int                     `json:"IOPSWrite"`
    BPSRead             int                     `json:"BPSRead"`
    BPSWrite            int                     `json:"BPSWrite"`
    TimeStamp           string                  `json:"TimeStamp"`
}

type MonitorDataST struct {
    InstanceMonitorData []InstanceMonitorDataST `json:"InstanceMonitorData"`
}

type DescribeInstanceMonitorDataResult struct {
    GlobalResult

    MonitorData         MonitorDataST           `json:"MonitorData"`
}

func (self *Client) DescribeInstanceMonitorData(args *DescribeInstanceMonitorDataArgs) (result *DescribeInstanceMonitorDataResult, errorResult *ErrorResult) {
    res := new(DescribeInstanceMonitorDataResult)
    errRes := self.CallAPI("DescribeInstanceMonitorData", args, res)

    return res, errRes
}
