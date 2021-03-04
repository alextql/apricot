package helpers

import (
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"net/url"
	"os"
	"path/filepath"
)

func FileAbsPath(file string) (string, error) {
	var err error
	var info os.FileInfo
	if info, err = os.Stat(file); err != nil {
		if os.IsNotExist(err) {
			return "", errors.New(fmt.Sprintf("%s not a file", file))
		}

		return "", err
	}

	return filepath.Abs(info.Name())
}

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
