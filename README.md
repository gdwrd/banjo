# banjo

[![Build Status](https://travis-ci.org/nsheremet/banjo.svg?branch=master)](https://travis-ci.org/nsheremet/banjo)

'banjo' it's a simple web server for building simple Sinatra-like applications

## Install

```bash 
$ go get github.com/nsheremet/banjo
```

## Example Usage

Example file - `main.go`

```go
package main

import "banjo"

func main() {
  app := banjo.Create(banjo.DefaultConfig())
  app.Get("/", func(r banjo.Request) Response {
    return app.JSON(banjo.M{"foo":"bar"})
  })

  app.Run()
}
```

## License

`banjo` is primarily distributed under the terms of Mozilla Public License 2.0.

See LICENSE for details.