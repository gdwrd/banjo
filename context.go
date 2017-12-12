package banjo

import (
	"encoding/json"
	"fmt"
)

// Context struct
//
// For easy cooperate with Requests & Response in action flow
//
type Context struct {
	Request  Request
	Response Response
}

// JSON function
//
// This func allows you to easy returning a JSON response
// Example usage:
// app.JSON({"foo" : "bar"})
//
// Params:
// - data {map[string]interface{}}
//
// Response:
// - None
//
func (ctx *Context) JSON(data map[string]interface{}) {
	logger := CreateLogger()
	body, err := json.Marshal(data)

	if err != nil {
		str := fmt.Sprintf("Error while parsing JSON object:\nError: %v\nData: %s", err, data)
		logger.Error(str)

		ctx.InternalServerError()
		return
	}

	if ctx.Response.Headers == nil {
		ctx.Response.Headers = make(map[string]string)
	}

	ctx.Response.Headers["Content-Type"] = "application/json; charset=utf-8"
	ctx.Response.Body = string(body)

	if ctx.Response.Status == 0 {
		ctx.Response.Status = 200
	}
}

// HTML function
//
// Prepared Headers to return HTML page
//
// Params:
// - data {string} HTML content
//
// Response:
// - None
//
func (ctx *Context) HTML(data string) {
	if ctx.Response.Headers == nil {
		ctx.Response.Headers = make(map[string]string)
	}

	ctx.Response.Headers["Content-Type"] = "text/html"
	ctx.Response.Body = data

	if ctx.Response.Status == 0 {
		ctx.Response.Status = 200
	}
}

// RedirectTo function
//
// Allows you to redirect to another page
//
// Params:
// - url {string} path to redirect
//
// Response:
// - None
//
func (ctx *Context) RedirectTo(url string) {
	if ctx.Response.Headers == nil {
		ctx.Response.Headers = make(map[string]string)
	}

	ctx.Response.Headers["Location"] = url
	ctx.Response.Status = 301
}

// InternalServerError function
//
// Modify Context struct with 500 Status error & Internal Server Error body
//
// Params:
// - None
//
// Response:
// - None
func (ctx *Context) InternalServerError() {
	ctx.Response.Status = 500
	ctx.Response.Body = "Internal Server Error"
}
