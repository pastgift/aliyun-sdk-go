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

func CreateTimestampString(layout string) (string) {
    tmspStr := time.Now().UTC().Format(layout)

    return tmspStr
}

func CreateRandomString() (string) {
    rand.Seed(time.Now().UnixNano())
    randInt := rand.Int63()
    randStr := strconv.FormatInt(randInt, 36)

    return randStr
}

func ComputeSignature(stringToSignature string, accesskeySecret string) (signature string) {
    // Crypto by HMAC-SHA1
    hmacSha1 := hmac.New(sha1.New ,[]byte(accesskeySecret + "&"))
    hmacSha1.Write([]byte(stringToSignature))
    sign := hmacSha1.Sum(nil)

    // Encode to Base64
    base64Sign := base64.StdEncoding.EncodeToString(sign)

    return base64Sign
}

func PercentReplace(s string) string {
    s = strings.Replace(s, "+", "%20", -1)
    s = strings.Replace(s, "*", "%2A", -1)
    s = strings.Replace(s, "%7E", "~", -1)

    return s
}
