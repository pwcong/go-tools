package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	. "translator/channel"
)

var channel string
var key string
var source string
var target string
var format string

func init() {

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "\nUsage: %s command [Options] <Text>\n\nThe commands are:\n"+
			"\ttranslate\n\tdetect\n\tgetlangs\n\nOptions:\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.StringVar(&channel, "c", "yandex", "translation service provider: yandex")
	flag.StringVar(&key, "k", "", "api key for translation service")
	flag.StringVar(&source, "s", "zh", "source language")
	flag.StringVar(&target, "t", "en", "target language")
	flag.StringVar(&format, "f", "json", "response format")

	flag.Parse()

}

func translate(text string) {
	if channel == "yandex" {
		res, err := YanDexTranslate(key, format, target, text)

		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("\n" + res + "\n")
		}
	} else {
		fmt.Println("\n****** Unknown Channel \"" + channel + "\" ******")
	}

}

func detect(text string) {

	if channel == "yandex" {
		res, err := YanDexDetect(key, format, text)

		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("\n" + res + "\n")

		}
	} else {
		fmt.Println("\n****** Unknown Channel \"" + channel + "\" ******")
	}

}

func getlangs() {
	if channel == "yandex" {
		res, err := YanDexGetLangs(key, format)

		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("\n" + res + "\n")
		}

	} else {
		fmt.Println("\n****** Unknown Channel \"" + channel + "\" ******")
	}

}

func main() {

	args := flag.Args()

	if len(args) < 1 {
		flag.Usage()
	} else {

		// 校验令牌
		if (channel == "yandex") && key == "" {
			key = os.Getenv("API_KEY_" + strings.ToUpper(channel))
			if key == "" {
				fmt.Println("\n****** Lack Of Key. " +
					"Please set special api key by \"-k\" or set environment variable \"API_KEY_" + strings.ToUpper(channel) + "\" ******")
				return
			}

		}

		command := flag.Arg(0)

		if command == "translate" {

			text := flag.Arg(1)
			if text == "" {
				fmt.Println("\n****** Lack Of Text ******")
			} else {
				translate(text)
			}
		} else if command == "detect" {
			text := flag.Arg(1)
			if text == "" {
				fmt.Println("\n****** Lack Of Text ******")
			} else {
				detect(text)
			}
		} else if command == "getlangs" {
			getlangs()
		} else {
			fmt.Println("\n****** Unknown Command \"" + command + "\" ******")
		}

	}
}
