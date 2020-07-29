package main

import (
	"flag"
	"log"

	"github.com/webability-go/xamboo"
)

const VERSION = "0.0.2"

func main() {
	// *** system Language !!! preload

	var file string
	flag.StringVar(&file, "config", "", "configuration file")
	flag.Parse()

	if file == "" {
		log.Fatalln("The configuration file is missing as argument --config=file")
		return
	}

	xamboo.Run(file)
}
