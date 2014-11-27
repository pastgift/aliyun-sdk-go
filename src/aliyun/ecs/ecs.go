package ecs

/*
 * ECS-API-Reference 2014-05-26
 */

import (
    "fmt"

    "net/http"
    "net/url"
    "io/ioutil"

    "aliyun/util"
)

type EcsClient struct {
    AccessKeyId     string
    AccessKeySecret string

    ApiUrl          url.URL
    ApiQuery        url.Values
}

func NewEcsClient(accesskey_id string, accesskey_secret string) (newEcsClient *EcsClient) {
    client := new(EcsClient)

    client.AccessKeyId     = accesskey_id
    client.AccessKeySecret = accesskey_secret

    client.ApiUrl = url.URL{
        Scheme:"http",
        Host  :"ecs.aliyuncs.com",
        Path  :"/",
    }

    client.ApiQuery = url.Values{
        // Set by API function
        //"Action"

        // Set when API called
        //"Timestamp"           :{util.CreateTimestampString("2006-01-02T15:04:05Z")},
        //"SignatureNonce"      :{util.CreateRandomString()},
        //"Signature"

        // Optional Field
        //"ResourceOwnerAccount":{""},

        "Format"              :{"json"},
        "Version"             :{"2014-05-26"},
        "AccessKeyId"         :{client.AccessKeyId},
        "SignatureMethod"     :{"HMAC-SHA1"},
        "SignatureVersion"    :{"1.0"},
    }

    return client
}

func (self *EcsClient) getStringToSign() (string_to_sign string) {
    s := "GET" + "&" +
        url.QueryEscape("/") + "&" +
        url.QueryEscape(self.ApiQuery.Encode())

    return s
}

func (self *EcsClient) addSignature() {
    //TODO Check query is prepared

    string_to_sign := self.getStringToSign()
    sign := util.ComputeSignature(string_to_sign, self.AccessKeySecret)

    self.ApiQuery.Set("Signature", sign)
}

func (self *EcsClient) prepareAction(action string) {
    self.ApiQuery.Set("Action", action)
    self.ApiQuery.Set("Timestamp", util.CreateTimestampString("2006-01-02T15:04:05Z"))
    self.ApiQuery.Set("SignatureNonce", util.CreateRandomString())
}

func (self *EcsClient) prepareUrl() {
    self.ApiUrl.RawQuery = self.ApiQuery.Encode()
}

func (self *EcsClient) callApi() (result string, err error) {
    // All ECS APIs are using GET method
    method := "GET"

    http_client := new(http.Client)
    req, err := http.NewRequest(method, self.ApiUrl.String(), nil)
    if err != nil {
        return "", err
    }

    resp, err := http_client.Do(req)
    defer resp.Body.Close()
    if err != nil {
        return "", err
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }

    result = string(body)
    return result, nil
}

func (self *EcsClient) DescribeRegions() (result string, err error) {
    self.prepareAction("DescribeRegions")
    self.addSignature()
    self.prepareUrl()

    fmt.Println(">>>", self.ApiUrl.String())
    fmt.Println(">>>", self.ApiQuery.Get("Signature"))

    return self.callApi()
}
