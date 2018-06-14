package utils

import (
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
