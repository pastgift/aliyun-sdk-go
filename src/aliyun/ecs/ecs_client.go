package ecs

import (
    "log"
    "time"
    "net/http"
    "io/ioutil"
)

type ECSClient struct {
    AccessKeyId     string  "AccessKey Id"
    AccessKeySecret string  "AccessKey Secret"
}

func NewECSClient(accesskey_id string, accesskey_secret string) (newECSClient *ECSClient) {
    client := &ECSClient{
        AccessKeyId     : accesskey_id,
        AccessKeySecret : accesskey_secret,
    }

    return client
}

func (self *ECSClient) Do(ecs_req *ECSRequest) (resp ECSResponse, err error) {
    resp = ECSResponse{}

    // Sign request
    ecs_req.Sign(self.AccessKeyId, self.AccessKeySecret)

    http_client := &http.Client{}

    ecs_req.URL.RawQuery = ecs_req.Query.Encode()
    http_req, err := http.NewRequest(ecs_req.Method, ecs_req.URL.String(), nil)
    if err != nil {
        return resp, err
    }

    t0 := time.Now()
    http_resp, err := http_client.Do(http_req)
    t1 := time.Now()
    log.Printf("The API call took %v to run.\n", t1.Sub(t0))

    if err != nil {
        return resp, err
    }

    body, err := ioutil.ReadAll(http_resp.Body)
    if err != nil {
        return resp, err
    }
    defer http_resp.Body.Close()

    resp.StatusCode = http_resp.StatusCode
    resp.RawBody    = body

    return resp, nil
}