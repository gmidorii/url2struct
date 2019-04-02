# url2struct

## Overview

generate Go query/response struct from raw url.

## Usage

### Command

build

```sh
cd cmd/u2s
go build
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
url2struct.Generate("http://example.com/sample", queryFileWriter, responseFileWriter)
```

## Example

mock server
```sh
go run mock/mock.go
```

command
```sh
# output to stdout
./u2s -u "http://localhost:6666/example?hoge=2&fuga=1.0&foo=str"

# output to query/response file.
./u2s -u "http://localhost:6666/example?hoge=2&fuga=1.0&foo=str" -q ./query.go -r ./res.go
```