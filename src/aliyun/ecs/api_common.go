package ecs

import (
    "log"
)

type GlobalResult struct {
    RequestId string `json:"RequestId"`
}

func (self *Client) CallAPI(req *Request, res interface{}) (err error) {
    resp, err := self.Do(req)
    if err != nil {
        log.Println("Do Request Error")
        return err
    }

    err = resp.FillTo(res)
    if err != nil {
        log.Println("Load Json Error")
        return err
    }

    log.Println("Received:\n\t", resp.String())
    return nil
}
