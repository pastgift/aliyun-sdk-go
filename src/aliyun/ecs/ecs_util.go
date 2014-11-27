/*
 * ECS-API-Reference 2014-05-26
 */

package ecs

import (
    "net/http"
    "net/url"
    "io/ioutil"

    "aliyun/util"
)

const (
    API_SCHEME_HTTP  = "http"
    API_SCHEME_HTTPS = "https"
    API_HOST         = "ecs.aliyuncs.com"
    API_PATH         = "/"

    API_FMT_JSON     = "json"
    API_FMT_XML      = "xml"
    API_VER          = "2014-05-26"
    API_SIGN_METHOD  = "HMAC-SHA1"
    API_SIGN_VER     = "1.0"

    API_TMSP_LAYOUT  = "2006-01-02T15:04:05Z"
)

// Handler to access ECS API
type ECSClient struct {
    AccessKeyId     string
    AccessKeySecret string

    APIUrl          url.URL
    APIQuery        url.Values
}

func NewECSClient(accesskey_id string, accesskey_secret string) (newECSClient *ECSClient) {
    client := new(ECSClient)

    client.AccessKeyId     = accesskey_id
    client.AccessKeySecret = accesskey_secret

    client.APIUrl = url.URL{
        Scheme:API_SCHEME_HTTP,
        Host  :API_HOST,
        Path  :API_PATH,
    }

    client.APIQuery = url.Values{
        // Global Queries
        //"Timestamp"
        //"SignatureNonce"
        //"Signature"

        "Format"              :{API_FMT_JSON},
        "Version"             :{API_VER},
        "AccessKeyId"         :{client.AccessKeyId},
        "SignatureMethod"     :{API_SIGN_METHOD},
        "SignatureVersion"    :{API_SIGN_VER},

        // Optional Field
        //"ResourceOwnerAccount",
    }

    return client
}

func (self *ECSClient) setAction(action string) {
    self.APIQuery.Set("Action", action)
}

func (self *ECSClient) getStringToSign() (string_to_sign string) {
    s := "GET" + "&" +
        url.QueryEscape("/") + "&" +
        url.QueryEscape(self.APIQuery.Encode())

    return s
}

func (self *ECSClient) completeSignature() {
    // Each request need a new `Timestamp` and a new `SignatureNonce`
    self.APIQuery.Set("Timestamp", util.CreateTimestampString(API_TMSP_LAYOUT))
    self.APIQuery.Set("SignatureNonce", util.CreateRandomString())

    // Signature must be done after all the query terms a ready
    string_to_sign := self.getStringToSign()
    sign := util.ComputeSignature(string_to_sign, self.AccessKeySecret)
    self.APIQuery.Set("Signature", sign)
}

func (self *ECSClient) doRequest(result string, err error) {
    http_client := &http.Client{}

    // All ECS APIs are using GET method
    method := "GET"
    self.APIUrl.RawQuery = self.APIQuery.Encode()
    req, err := http.NewRequest(method, self.APIUrl.String(), nil)
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

    return string(body), nil
}

// Easy way to call API
func (self *ECSClient) CallAPI(action string) (result string, err error){
    self.setAction(action)
    self.completeSignature()
    return self.doRequest()
}