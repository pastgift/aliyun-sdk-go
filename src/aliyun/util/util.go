package util

import (
    "time"
    "math/rand"
    "strconv"
    "strings"
    "encoding/base64"
    "crypto/sha1"
    "crypto/hmac"
)

const (
    ECS_TMSP_LAYOUT  = "2006-01-02T15:04:05Z"
)

func CreateSpecifiedTimestampString(year, month, day, hour, min, sec int) string {
    t := time.Date(year, time.Month(month), day, hour, min, sec, 0, time.UTC)
    tmspStr := t.Format(ECS_TMSP_LAYOUT)

    return tmspStr
}

func CreateTimestampString() string {
    tmspStr := time.Now().UTC().Format(ECS_TMSP_LAYOUT)

    return tmspStr
}

func CreateRandomString() string {
    rand.Seed(time.Now().UnixNano())
    randInt := rand.Int63()
    randStr := strconv.FormatInt(randInt, 36)

    return randStr
}

func CreateSignature(stringToSignature, accesskeySecret string) string {
    // Crypto by HMAC-SHA1
    hmacSha1 := hmac.New(sha1.New ,[]byte(accesskeySecret + "&"))
    hmacSha1.Write([]byte(stringToSignature))
    sign := hmacSha1.Sum(nil)

    // Encode to Base64
    base64Sign := base64.StdEncoding.EncodeToString(sign)

    return base64Sign
}

func PercentReplace(str string) string {
    str = strings.Replace(str, "+", "%20", -1)
    str = strings.Replace(str, "*", "%2A", -1)
    str = strings.Replace(str, "%7E", "~", -1)

    return str
}