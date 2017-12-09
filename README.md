# BANjO

[![Build Status](https://travis-ci.org/nsheremet/banjo.svg?branch=master)](https://travis-ci.org/nsheremet/banjo)
[![Software License](https://img.shields.io/badge/License-MPL--2.0-green.svg)](https://github.com/nsheremet/banjo/blob/master/LICENSE.md)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/nsheremet/banjo)
[![Coverage Status](http://codecov.io/github/nsheremet/banjo/coverage.svg?branch=master)](http://codecov.io/github/nsheremet/banjo?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/nsheremet/banjo)](https://goreportcard.com/report/github.com/nsheremet/banjo)

**banjo** it's a simple web framework for building simple web applications

## Install

```bash 
$ go get github.com/nsheremet/banjo
```

## Example Usage

Simple Web App - `main.go`

```go
package main

import "banjo"

func main() {
  app := banjo.Create(banjo.DefaultConfig())
  
  app.Get("/", func(ctx *banjo.Context) {
    ctx.JSON(banjo.M{"foo":"bar"})
  })

  app.Run()
}
```

Example responses:

```go
// ... Redirect To
  app.Get("/admin", func(ctx *banjo.Context) {
    ctx.RedirectTo("/")
  })
// ... HTML
  app.Get("/foo", func(ctx *banjo.Context) {
    ctx.HTML("<h1>Hello from BONjO!</h1>")
  })
// ... Return Params as JSON
  app.Post("/bar", func(ctx *banjo.Context) {
    ctx.JSON(banjo.M{
      "params": ctx.Request.Params
    })
    ctx.Response.Status = 201
  })
```

## License

`banjo` is primarily distributed under the terms of Mozilla Public License 2.0.

See LICENSE for details.