// Common functions for ECS

package ecs

import (
	"log"
)

const (
	MSG_API_ERR  = "API Calling Error"
	MSG_REQ_ERR  = "Do Request Error"
	MSG_JSON_ERR = "Load Json Error"
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

func (self *Client) CallAPI(action string, args interface{}, res interface{}) (errorResult *ErrorResult) {
	errRes := new(ErrorResult)

	req := NewRequest(action)
	req.SetArgs(args)

	resp, err := self.Do(req)
	if err != nil {
		// Do request Error is ususally unexcepted.
		log.Panicln(MSG_REQ_ERR)
	}

	shellColor := "\033[0;33m"
	shellNormal := "\033[m"
	log.Printf("%s[API RAW RESPONSE] Status Code: %d %s", shellColor, resp.StatusCode, shellNormal)
	log.Printf("%s[API RAW RESPONSE] Received:\n\t%s %s", shellColor, resp.String(), shellNormal)

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
