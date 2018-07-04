package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/pwcong/go-tools/text/coding/utils"
)

var method string

func init() {

	flag.StringVar(&method, "m", "unicode_encode", `coding convert method. the optional values are as follows: 
		* base64_encode
		* base64_decode
		* unicode_encode
		* unicode_decode
		* url_encode
		* url_decode
		* md5
		* sha1
		* sha256
		* sha512
		`)

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "\nUsage: %s [Options] <Source>\n\nOptions:\n\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

}

func main() {
	args := flag.Args()

	if len(args) != 1 {
		flag.Usage()
	} else {

		source := flag.Arg(0)

		var res string
		var err error

		if method == "base64_encode" {
			res, err = utils.Base64Encode(source)
		} else if method == "base64_decode" {
			res, err = utils.Base64Decode(source)
		} else if method == "unicode_encode" {
			res, err = utils.UnicodeEncode(source)
		} else if method == "unicode_decode" {
			res, err = utils.UnicodeDecode(source)
		} else if method == "url_encode" {
			res, err = utils.URLEncode(source)
		} else if method == "url_decode" {
			res, err = utils.URLDecode(source)
		} else if method == "md5" {
			res, err = utils.MD5Encrypt(source)
		} else if method == "sha1" {
			res, err = utils.SHA1Encrypt(source)
		} else if method == "sha256" {
			res, err = utils.SHA256Encrypt(source)
		} else if method == "sha512" {
			res, err = utils.SHA512Encrypt(source)
		} else {
			err = errors.New("Unknown method name")
		}

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(res)

	}
}
