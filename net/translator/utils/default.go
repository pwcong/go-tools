package utils

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
)

const (
	JSONFORMAT_PREFIX = ""
	JSONFORMAT_INDENT = "  "

	XML_PREFIX = ""
	XML_INDENT = "  "
)

func FormatJSON(bytes []byte) (string, error) {

	var data map[string]interface{}
	if err := json.Unmarshal(bytes, &data); err != nil {
		return "", err
	}

	res, err := json.MarshalIndent(data, JSONFORMAT_PREFIX, JSONFORMAT_INDENT)
	if err != nil {
		return "", err
	}

	return string(res), nil

}

func FormatXML(bytes []byte, object interface{}) (string, error) {

	if err := xml.Unmarshal(bytes, object); err != nil {
		return "", err
	}

	res, err := xml.MarshalIndent(object, XML_PREFIX, XML_INDENT)
	if err != nil {
		return "", err
	}

	return string(res), nil
}

func Get(url string) ([]byte, error) {
	t, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}

	if t.StatusCode != 200 {
		return []byte{}, errors.New(t.Status)
	}

	defer t.Body.Close()

	body, err := ioutil.ReadAll(t.Body)

	return body, err
}

func GetJSON(url string) (string, error) {

	body, err := Get(url)
	if err != nil {
		return "", err
	}

	return FormatJSON(body)

}
