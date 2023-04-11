package utility

import (
	"crypto/md5"
	"fmt"
	url2 "net/url"
	"strings"
)

func PrepareUrl(url string) string {
	if url[:4] != "http" {
		url = "http://" + url
	}
	return url
}

func ValidUrl(url string) bool {
	if _, err := url2.ParseRequestURI(url); err != nil {
		return false
	}
	if u, err := url2.Parse(url); err != nil || u.Scheme == "" || u.Host == "" || len(strings.Split(u.Host, ".")) < 2 {
		return false
	}
	return true
}

func Hash(bytes []byte) string {
	md5Str := md5.Sum(bytes)
	return fmt.Sprintf("%x", md5Str)
}
