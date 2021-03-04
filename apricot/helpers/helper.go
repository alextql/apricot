package helpers

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io"
	"net/url"
)

func Base64Encode(val interface{}) string {
	return base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%v", val)))
}

func UrlEscape(uri string) string {
	unEscape := UrlUnEscape(uri)
	return url.QueryEscape(unEscape)
}

func UrlUnEscape(uri string) string {
	unescapeUrl, _ := url.QueryUnescape(uri)
	return unescapeUrl
}

func Sha1Encode(text string) string {
	hash := sha1.New()
	_, _ = io.WriteString(hash, text)
	return fmt.Sprintf("%x", hash.Sum(nil))
}
