package main

import (
	"flag"
	"fmt"
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

		rawQ, rawRes, err := url2struct.Generate(*u)
		if err != nil {
			log.Printf("command error: %v\n", err)
		}
		fmt.Fprintln(qf, rawQ)
		fmt.Fprintln(rf, rawRes)
	} else {
		rawQ, rawRes, err := url2struct.Generate(*u)
		if err != nil {
			log.Printf("command error: %v\n", err)
		}
		fmt.Fprintln(os.Stdout, rawQ)
		fmt.Fprintln(os.Stdout, rawRes)
	}
}
