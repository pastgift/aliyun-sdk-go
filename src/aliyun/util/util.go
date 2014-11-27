package util

import (
    "time"
    "math/rand"
    "strconv"
    "encoding/base64"
    "crypto/sha1"
    "crypto/hmac"
)

func CreateTimestampString(layout string) (string) {
    tmsp_str := time.Now().UTC().Format(layout)

    return tmsp_str
}

func CreateRandomString() (string) {
    rand.Seed(time.Now().UnixNano())
    rand_int := rand.Int63()
    rand_str := strconv.FormatInt(rand_int, 36)

    return rand_str
}

func ComputeSignature(string_to_signature string, accesskey_secret string) (signature string) {
    // Crypto by HMAC-SHA1
    hmac_sha1 := hmac.New(sha1.New ,[]byte(accesskey_secret + "&"))
    hmac_sha1.Write([]byte(string_to_signature))
    sign := hmac_sha1.Sum(nil)

    // Encode to Base64
    sign_base64 := base64.StdEncoding.EncodeToString(sign)

    return sign_base64
}