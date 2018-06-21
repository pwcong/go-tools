package channel

import (
	"encoding/xml"
	"net/url"

	"github.com/pwcong/go-tools/net/translator/utils"
)

const (
	YANDEX_API_GETLANGS_JSON = "https://translate.yandex.net/api/v1.5/tr.json/getLangs"
	YANDEX_API_GETLANGS_XML  = "https://translate.yandex.net/api/v1.5/tr/getLangs"

	YANDEX_API_DETECT_JSON = "https://translate.yandex.net/api/v1.5/tr.json/detect"
	YANDEX_API_DETECT_XML  = "https://translate.yandex.net/api/v1.5/tr/detect"

	YANDEX_API_TRANSLATE_JSON = "https://translate.yandex.net/api/v1.5/tr.json/translate"
	YANDEX_API_TRANSLATE_XML  = "https://translate.yandex.net/api/v1.5/tr/translate"
)

func YanDexTranslate(key string, format string, lang string, text string) (string, error) {

	v := url.Values{}
	v.Set("key", key)
	v.Set("text", text)
	v.Set("lang", lang)
	query := v.Encode()

	if format == "json" {

		return utils.GetJSON(YANDEX_API_TRANSLATE_JSON + "?" + query)

	} else if format == "xml" {
		body, err := utils.Get(YANDEX_API_TRANSLATE_XML + "?" + query)

		if err != nil {
			return "", err
		}

		type Translation struct {
			XMLName xml.Name `xml:"Translation"`
			Code    int      `xml:"code,attr"`
			Lang    string   `xml:"lang,attr"`
			Text    string   `xml:"text"`
		}
		t := new(Translation)

		return utils.FormatXML(body, t)
	}

	return "****** Unsupport Format \"" + format + "\" ******", nil
}

func YanDexDetect(key string, format string, text string) (string, error) {

	v := url.Values{}
	v.Set("key", key)
	v.Set("text", text)
	query := v.Encode()

	if format == "json" {

		return utils.GetJSON(YANDEX_API_DETECT_JSON + "?" + query)

	} else if format == "xml" {

		body, err := utils.Get(YANDEX_API_DETECT_XML + "?" + query)
		if err != nil {
			return "", err
		}

		type DetectedLang struct {
			XMLName xml.Name `xml:"DetectedLang"`
			Code    int      `xml:"code,attr"`
			Lang    string   `xml:"lang,attr"`
		}
		t := new(DetectedLang)

		return utils.FormatXML(body, t)

	}

	return "****** Unsupport Format \"" + format + "\" ******", nil

}

func YanDexGetLangs(key string, format string) (string, error) {

	if format == "json" {

		return utils.GetJSON(YANDEX_API_GETLANGS_JSON + "?key=" + key)

	} else if format == "xml" {

		body, err := utils.Get(YANDEX_API_GETLANGS_XML + "?key=" + key)
		if err != nil {
			return "", err
		}

		type Langs struct {
			XMLName xml.Name `xml:"Langs"`
			Dirs    []string `xml:"dirs>string"`
		}
		t := new(Langs)

		return utils.FormatXML(body, t)

	}

	return "****** Unsupport Format \"" + format + "\" ******", nil

}
