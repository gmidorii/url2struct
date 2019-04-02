package main

import (
	"flag"
	"log"

	"github.com/gmidorii/url2struct"
)

func main() {
	u := flag.String("u", "http://localhost:6666/example?hoge=2&fuga=1.0&foo=str", "url")
	flag.Parse()

	if err := url2struct.Generate(*u); err != nil {
		log.Printf("command error: %v\n", err)
	}
}
