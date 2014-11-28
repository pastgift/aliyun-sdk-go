package ecs

import (
    "net/http"
    "io/ioutil"
)

type ECSClient struct {
    AccessKeyId     string
    AccessKeySecret string
}

func NewECSClient(accesskey_id string, accesskey_secret string) (newECSClient *ECSClient) {
    client := new(ECSClient)

    client.AccessKeyId     = accesskey_id
    client.AccessKeySecret = accesskey_secret

    return client
}

func (self *ECSClient) Do(ecs_req *ECSRequest) (resp string, err error) {
    // Sign request
    ecs_req.Sign(self.AccessKeyId, self.AccessKeySecret)

    http_client := &http.Client{}

    ecs_req.URL.RawQuery = ecs_req.Query.Encode()
    http_req, err := http.NewRequest(ecs_req.Method, ecs_req.URL.String(), nil)
    if err != nil {
        return "", err
    }

    http_resp, err := http_client.Do(http_req)
    defer http_resp.Body.Close()

    if err != nil {
        return "", err
    }

    body, err := ioutil.ReadAll(http_resp.Body)
    if err != nil {
        return "", err
    }

    return string(body), nil
}