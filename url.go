package url2struct

import (
	"bytes"
	"fmt"
	"go/format"
	"net/http"
	"net/url"
	"strconv"

	"github.com/ChimeraCoder/gojson"
	"github.com/iancoleman/strcase"
)

func Generate(rawurl string) (rawQueryStruct, rawResStruct string, err error) {
	rawQueryStruct, err = query(rawurl)
	if err != nil {
		return
	}
	rawResStruct, err = response(rawurl)
	if err != nil {
		return
	}
	return
}

func query(rawurl string) (string, error) {
	u, err := url.Parse(rawurl)
	if err != nil {
		return "", err
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
		return "", err
	}

	return string(src), nil
}

func field(key, val string) string {
	switch {
	case isInt(val):
		return fieldFormat("int", key, val)
	case isFloat64(val):
		return fieldFormat("float64", key, val)
	case isBool(val):
		return fieldFormat("bool", key, val)
	default:
		return fieldFormat("string", key, val)
	}
}

func fieldFormat(vType, key, val string) string {
	return fmt.Sprintf("%v %v`url:\"%v\"`;", strcase.ToCamel(key), vType, strcase.ToSnake(key))
}

func isBool(v string) bool {
	_, err := strconv.ParseBool(v)
	if err != nil {
		return false
	}
	return true
}

func isInt(v string) bool {
	_, err := strconv.Atoi(v)
	if err != nil {
		return false
	}
	return true
}

func isFloat64(v string) bool {
	// check int value.
	if isInt(v) {
		return false
	}

	_, err := strconv.ParseFloat(v, 64)
	if err != nil {
		return false
	}
	return true
}

func response(rawurl string) (string, error) {
	resp, err := http.Get(rawurl)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	b, err := gojson.Generate(resp.Body, gojson.ParseJson, "AutoResponse", "response", []string{"json"}, true, true)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
