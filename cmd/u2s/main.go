package main

import (
	"flag"
	"log"
	"os"

	"github.com/gmidorii/url2struct"
)

func main() {
	u := flag.String("u", "http://localhost:6666/example?hoge=2&fuga=1.0&foo=str", "url")
	q := flag.String("q", "", "query struct file")
	r := flag.String("r", "", "response struct file")
	flag.Parse()

	if *q != "" && *r != "" {
		qf, err := os.Create(*q)
		if err != nil {
			log.Println(err)
		}
		defer qf.Close()

		rf, err := os.Create(*r)
		if err != nil {
			log.Println(err)
		}
		defer rf.Close()

		if err := url2struct.Generate(*u, qf, rf); err != nil {
			log.Printf("command error: %v\n", err)
		}
	} else {
		if err := url2struct.Generate(*u, os.Stdout, os.Stdout); err != nil {
			log.Printf("command error: %v\n", err)
		}
	}
}
