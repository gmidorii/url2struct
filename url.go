package url2struct

import (
	"bytes"
	"fmt"
	"go/format"
	"io"
	"net/http"
	"net/url"

	"github.com/ChimeraCoder/gojson"
	"github.com/iancoleman/strcase"
)

func Generate(rawurl string, queryWriter, responseWriter io.Writer) error {

	if err := query(rawurl, queryWriter); err != nil {
		return err
	}
	if err := response(rawurl, responseWriter); err != nil {
		return err
	}

	return nil
}

func query(rawurl string, w io.Writer) error {
	u, err := url.Parse(rawurl)
	if err != nil {
		return err
	}

	q := u.Query()

	buf := bytes.Buffer{}
	fmt.Fprint(&buf, "type AutoQuery struct {")
	for k, v := range q {
		// use array values only lead.
		fmt.Fprint(&buf, field(k, v[0]))
	}
	fmt.Fprint(&buf, "}\n")

	src, err := format.Source(buf.Bytes())
	if err != nil {
		return err
	}

	fmt.Fprint(w, string(src))

	return nil
}

func field(key string, val interface{}) string {
	switch val.(type) {
	case float64:
		return fmt.Sprintf("%v float64`url:\"%v\"`;", strcase.ToCamel(key), strcase.ToSnake(key))
	case int:
		return fmt.Sprintf("%v int`url:\"%v\"`;", strcase.ToCamel(key), strcase.ToSnake(key))
	default:
		return fmt.Sprintf("%v string`url:\"%v\"`;", strcase.ToCamel(key), strcase.ToSnake(key))
	}
}

func response(rawurl string, w io.Writer) error {
	resp, err := http.Get(rawurl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	b, err := gojson.Generate(resp.Body, gojson.ParseJson, "AutoResponse", "response", []string{"json"}, true, true)
	if err != nil {
		return err
	}

	fmt.Fprint(w, string(b))
	return nil
}
