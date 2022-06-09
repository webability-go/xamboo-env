package main

import (
	"flag"
	"log"

	"golang.org/x/text/language"

	"github.com/webability-go/xamboo"
)

const VERSION = "0.1.0"

func main() {
	// *** system Language !!! preload

	var file string
	var slang string
	flag.StringVar(&file, "config", "", "Configuration file")
	flag.StringVar(&slang, "language", "", "Systam language")
	flag.Parse()

	if file == "" {
		log.Fatalln("The configuration file is missing as argument --config=filepath")
		return
	}
	lang := language.English
	if slang != "" {
		tlang, err := language.Parse(slang)
		if err != nil {
			log.Fatalln("The language is not parseable (known languages: es, en, fr) --language=en")
			return
		}
		lang = tlang
	}

	xamboo.Run(file, lang)
}
