package utils

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

func UnicodeEncode(source string) (string, error) {
	res := strconv.QuoteToASCII(source)
	return res[1 : len(res)-1], nil
}
func UnicodeDecode(source string) (string, error) {

	words := strings.Split(source, "\\u")
	var res string

	for _, v := range words {
		if len(v) < 1 {
			continue
		}

		t, err := strconv.ParseInt(v, 16, 32)
		if err != nil {
			return "", err
		}

		res += fmt.Sprintf("%c", t)

	}

	return res, nil

}

func URLEncode(source string) (string, error) {
	return url.QueryEscape(source), nil
}

func URLDecode(source string) (string, error) {
	return url.QueryUnescape(source)
}

func MD5Encrypt(source string) (string, error) {
	res := fmt.Sprintf("%x", md5.Sum([]byte(source)))
	return res, nil
}

func SHA1Encrypt(source string) (string, error) {
	res := fmt.Sprintf("%x", sha1.Sum([]byte(source)))
	return res, nil
}

func SHA256Encrypt(source string) (string, error) {
	res := fmt.Sprintf("%x", sha256.Sum256([]byte(source)))
	return res, nil
}

func SHA512Encrypt(source string) (string, error) {
	res := fmt.Sprintf("%x", sha512.Sum512([]byte(source)))
	return res, nil
}

func Base64Encode(source string) (string, error) {
	return base64.StdEncoding.EncodeToString([]byte(source)), nil
}

func Base64Decode(source string) (string, error) {

	bytes, err := base64.StdEncoding.DecodeString(source)
	if err != nil {
		return "", nil
	}
	return string(bytes), nil
}
