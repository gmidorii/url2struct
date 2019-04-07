# url2struct

## Overview

generate Go query/response struct from raw url.

## Usage

### Install

```sh
go get github.com/gmidorii/url2struct
```

### Command

build

```sh
make build
```

run

```sh
./u2s -h
Usage of ./u2s:
  -q string
        query struct file
  -r string
        response struct file
  -u string
        url (default "http://localhost:6666/example?hoge=2&fuga=1.0&foo=str")
```

### Module

```go
if err := url2struct.Generate("http://example.com/sample", queryFileWriter, responseFileWriter); err != nil {
}
```

## Example

start mock server

```sh
# localhost:6666
go run mock/mock.go
```

command

```sh
# output to stdout
./u2s -u "http://localhost:6666/example?hoge=2&fuga=1.0&foo=str"

# output to query/response file.
./u2s -u "http://localhost:6666/example?hoge=2&fuga=1.0&foo=str" -q ./query.go -r ./res.go
```
