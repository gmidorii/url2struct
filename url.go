package url2struct

import (
	"bytes"
	"fmt"
	"go/format"
	"net/http"
	"net/url"

	"github.com/ChimeraCoder/gojson"
	"github.com/iancoleman/strcase"
)

func Field(key string, val interface{}) string {
	switch val.(type) {
	case float64:
		return fmt.Sprintf("%v float64`url:\"%v\"`;", strcase.ToCamel(key), strcase.ToSnake(key))
	case int:
		return fmt.Sprintf("%v int`url:\"%v\"`;", strcase.ToCamel(key), strcase.ToSnake(key))
	default:
		return fmt.Sprintf("%v string`url:\"%v\"`;", strcase.ToCamel(key), strcase.ToSnake(key))
	}
}

func Generate(rawurl string) error {
	u, err := url.Parse(rawurl)
	if err != nil {
		return err
	}

	q := u.Query()

	buf := bytes.Buffer{}
	fmt.Fprint(&buf, "type AutoQuery struct {")
	for k, v := range q {
		// use array values only lead.
		fmt.Fprint(&buf, Field(k, v[0]))
	}
	fmt.Fprint(&buf, "}")

	src, err := format.Source(buf.Bytes())
	if err != nil {
		return err
	}

	fmt.Println(string(src))

	if err := response(rawurl); err != nil {
		fmt.Println(err)
	}

	return nil
}

func response(rawurl string) error {
	resp, err := http.Get(rawurl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	b, err := gojson.Generate(resp.Body, gojson.ParseJson, "AutoResponse", "response", []string{"json"}, true, true)
	if err != nil {
		return err
	}
	fmt.Println(string(b))
	return nil
}
