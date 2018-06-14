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

	flag.StringVar(&method, "f", "unicode_encode", `coding convert method. the optional values are as follows: 
		* unicode_encode
		* unicode_decode
		* url_encode
		* url_dncode
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
		if method == "unicode_encode" {
			res, err = utils.UnicodeEncode(source)
		} else if method == "unicode_decode" {
			res, err = utils.UnicodeDecode(source)
		} else if method == "url_encode" {
			res, err = utils.URLEncode(source)
		} else if method == "url_dncode" {
			res, err = utils.URLDecode(source)
		} else {
			err = errors.New("Unknown method name")
		}

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(res)

	}
}
