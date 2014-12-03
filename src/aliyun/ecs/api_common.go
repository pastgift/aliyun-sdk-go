package ecs

import (
    "log"
)

const (
    MSG_API_ERR     = "API Calling Error"
    MSG_REQ_ERR     = "Do Request Error"
    MSG_JSON_ERR    = "Load Json Error"
)

type GlobalResult struct {
    RequestId string `json:"RequestId"`
}

type ErrorResult struct {
    GlobalResult
    HostId  string `json:"HostId"`
    Code    string `json:"Code"`
    Message string `json:"Message"`
}

func (self *Client) CallAPI(req *Request, res interface{}) (errorResult *ErrorResult) {
    errRes := new(ErrorResult)

    resp, err := self.Do(req)
    if err != nil {
        // Do request Error is ususally unexcepted.
        log.Panicln(MSG_REQ_ERR)
    }

    log.Println("Status Code:", resp.StatusCode)
    log.Println("Received:\n\t", resp.String())

    // Struct of Returning ErrorData is fixed.
    if resp.StatusCode >= 400 {
        // Let user to handle the API error result
        resp.FillTo(errRes)

        log.Println(MSG_API_ERR)
        return errRes
    }

    err = resp.FillTo(res)
    if err != nil {
        // Do request Error is ususally unexcepted.
        log.Panicln(MSG_JSON_ERR)
    }

    return nil
}
