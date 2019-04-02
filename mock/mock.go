package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = ":6666"

func main() {
	http.HandleFunc("/example", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `
			{
				"foo": 1,
				"hoge": "hoge"
			}
		`)
	})
	log.Printf("listen port=%v \n", port)
	http.ListenAndServe(port, nil)
}
