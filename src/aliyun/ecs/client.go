package ecs

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Client struct {
	AccessKeyId     string "AccessKey Id"
	AccessKeySecret string "AccessKey Secret"
}

func NewClient(accesskey_id, accesskey_secret string) *Client {
	client := &Client{
		AccessKeyId:     accesskey_id,
		AccessKeySecret: accesskey_secret,
	}

	return client
}

func (self *Client) Do(req *Request) (resp Response, err error) {
	resp = Response{}

	// Sign request
	req.Sign(self.AccessKeyId, self.AccessKeySecret)

	httpClient := &http.Client{}

	req.URL.RawQuery = req.Query.Encode()
	httpReq, err := http.NewRequest(req.Method, req.URL.String(), nil)
	if err != nil {
		return resp, err
	}

	t0 := time.Now()
	httpResp, err := httpClient.Do(httpReq)
	t1 := time.Now()
	log.Printf("The API call took %v to run.\n", t1.Sub(t0))

	if err != nil {
		return resp, err
	}

	body, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return resp, err
	}
	defer httpResp.Body.Close()

	resp.StatusCode = httpResp.StatusCode
	resp.RawBody = body

	return resp, nil
}
