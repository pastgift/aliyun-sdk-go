package ecs

import (
	"encoding/json"
)

type Response struct {
	StatusCode int    "HTTP Status Code"
	RawBody    []byte "Raw byte data"
}

func NewResponse() *Response {
	resp := &Response{
		StatusCode: 0,
	}

	return resp
}

func (self *Response) String() string {
	return string(self.RawBody)
}

func (self *Response) FillTo(target interface{}) error {
	return json.Unmarshal(self.RawBody, target)
}
