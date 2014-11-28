package ecs

import (
    "encoding/json"
)

type ECSResponse struct {
    StatusCode int      "HTTP Status Code"
    RawBody    []byte   "Raw byte data"
}

func NewECSResponse() *ECSResponse {
    ecs_resp := &ECSResponse{
        StatusCode:0,
    }

    return ecs_resp
}

func (self *ECSResponse) String() string {
    return string(self.RawBody)
}

func (self *ECSResponse) FillJson(v interface{}) error {
    return json.Unmarshal(self.RawBody, v)
}