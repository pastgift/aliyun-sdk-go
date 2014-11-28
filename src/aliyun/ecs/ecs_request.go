package ecs

import (
    "net/url"

    "aliyun/util"
)

const (
    REQ_SCHEME_HTTP  = "http"
    REQ_SCHEME_HTTPS = "https"
    REQ_HOST         = "ecs.aliyuncs.com"
    REQ_PATH         = "/"

    REQ_FMT_JSON     = "json"
    REQ_FMT_XML      = "xml"
    REQ_VER          = "2014-05-26"
    REQ_SIGN_METHOD  = "HMAC-SHA1"
    REQ_SIGN_VER     = "1.0"

    REQ_TMSP_LAYOUT  = "2006-01-02T15:04:05Z"
)

type ECSRequest struct {
    Method       string
    URL          url.URL
    Query        url.Values
}

func NewECSRequest(action string) (newECSRequest *ECSRequest) {
    req := new(ECSRequest)
    
    // All ECS APIs are using GET method
    req.Method = "GET"
    
    req.URL = url.URL{
        Scheme:REQ_SCHEME_HTTP,
        Host  :REQ_HOST,
        Path  :REQ_PATH,
    }

    req.Query = url.Values{
        "Action"              :{action},

        // Global Queries
        "Format"              :{REQ_FMT_JSON},
        "Version"             :{REQ_VER},

        // Optional Field
        //"ResourceOwnerAccount",
    }

    return req
}

func (self *ECSRequest) AddArgument(key string, value string) {
    self.Query.Set(key, value)
}

func (self *ECSRequest) Sign(accesskey_id string, accesskey_secret string) {
    self.AddArgument("SignatureMethod",     REQ_SIGN_METHOD)
    self.AddArgument("SignatureVersion",    REQ_SIGN_VER)
    self.AddArgument("AccessKeyId",         accesskey_id)

    // Each request need a new `Timestamp` and a new `SignatureNonce`
    self.AddArgument("Timestamp",           util.CreateTimestampString(REQ_TMSP_LAYOUT))
    self.AddArgument("SignatureNonce",      util.CreateRandomString())
    
    q := self.Query.Encode()
    q = util.PercentReplace(q)

    q = url.QueryEscape(q)
    q = util.PercentReplace(q)

    // string_to_sign = 
    //    "GET" + "&" + 
    //    url.QueryEscape("/") + "&" + q
    string_to_sign := "GET&%2F&" + q

    // Signature must be done after all the query terms a ready
    sign := util.ComputeSignature(string_to_sign, accesskey_secret)
    self.AddArgument("Signature", sign)
}