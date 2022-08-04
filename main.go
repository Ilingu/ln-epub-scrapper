package main

import (
	"flag"
	"log"
	"strings"
)

var appFailure bool

func main() {
	output := flag.String("output", ".", "Where the Epub file will be saved.")
	url := flag.String("url", "", "Perpetualdaydreams's Light Novel Info Page.")
	flag.Parse()

	if len(strings.TrimSpace(*url)) <= 0 {
		log.Println("❌ Please give me an Perpetualdaydreams url to scrap")
		return
	}

	RawChapters, config, ok := Scrap(*url)
	if appFailure || !ok {
		log.Fatal("❌ app failure")
		return
	}
	config.output = *output

	epubPath, succeed := RawChapters.ToEpub(config)
	if !succeed {
		log.Println("❌ Couldn't Convert To Epub")
		return
	}
	log.Printf("✅ Converted to Epub Successfully, epub file path: %s", epubPath)
}
