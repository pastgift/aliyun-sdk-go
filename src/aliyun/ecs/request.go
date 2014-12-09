package ecs

import (
    "net/url"
    "reflect"

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
)

type Request struct {
    Method       string     "Request Method"
    URL          url.URL    "URL without Query"
    Query        url.Values "Queries"
}

func NewRequest(action string) *Request {
    req := new(Request)

    // All ECS APIs are using GET method
    req.Method = "GET"

    req.URL = url.URL{
        Scheme:REQ_SCHEME_HTTP,
        Host  :REQ_HOST,
        Path  :REQ_PATH,
    }

    req.Query = url.Values{
        "Action"            :{action},

        // Global Queries
        "Format"            :{REQ_FMT_JSON},
        "Version"           :{REQ_VER},

        "Timestamp"         :{""},
        "SignatureNonce"    :{""},
        "SignatureMethod"   :{""},
        "SignatureVersion"  :{""},
        "AccessKeyId"       :{""},

        // Optional Field
        //"ResourceOwnerAccount",

        // Signature must be done after all other query terms are ready
        //"Signature"         :{""},
    }

    return req
}

func (self *Request) SetArg(key, value string) {
    // For deleting key(value == ""), Getting value
    // from inexisting key also return an empty string.
    if value == self.Query.Get(key) {
        return
    }

    if value == "" {
        self.Query.Del(key)
    }

    self.Query.Set(key, value)
}

func (self *Request) SetArgs(args interface{}) {
    // Since args from user is a pointer,
    // need reflect.Indirect to get it's value.
    v := reflect.Indirect(reflect.ValueOf(args))
    t := v.Type()

    for i := 0; i < v.NumField(); i++ {
       self.SetArg(t.Field(i).Name, v.Field(i).String())
    }
}

func (self *Request) Sign(accesskeyId, accesskeySecret string) {
    // Each request need a new `Timestamp` and a new `SignatureNonce`
    self.SetArg("Timestamp",           util.CreateTimestampString())
    self.SetArg("SignatureNonce",      util.CreateRandomString())

    self.SetArg("SignatureMethod",     REQ_SIGN_METHOD)
    self.SetArg("SignatureVersion",    REQ_SIGN_VER)
    self.SetArg("AccessKeyId",         accesskeyId)

    q := self.Query.Encode()
    q = util.PercentReplace(q)

    q = url.QueryEscape(q)
    q = util.PercentReplace(q)

    // stringToSign =
    //    "GET" + "&" +
    //    url.QueryEscape("/") + "&" + q
    stringToSign := "GET&%2F&" + q

    sign := util.CreateSignature(stringToSign, accesskeySecret)
    self.SetArg("Signature", sign)
}
