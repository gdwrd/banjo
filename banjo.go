// Package banjo is main package file
// Banjo Package
// Allows you to create your own simple web application
//
// See more in examples/example.go file
//
package banjo

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
)

// Banjo struct
// Main package Struct
//
type Banjo struct {
	config Config
	routes Routes
	parser Parser
	logger Logger
}

// Request struct using for passing as
// parameter to callback functions.
//
type Request struct {
	Headers     map[string]string
	MapParams   map[string]string
	Files       []map[string]string
	Params      string
	Method      string
	URL         string
	HTTPVersion string
}

// Response struct
// Using as returned value for callback function
//
type Response struct {
	Headers map[string]string
	Body    string
	Status  int
}

// M type is map[string]interface{} alias
//
type M map[string]interface{}

// Create Constructor:
//
// Initialize Banjo instance
//
// Params:
// - config {Config} Banjo configuration
//
// Response:
// - banjo {Banjo} Banjo configuration
//
func Create(config Config) Banjo {
	return Banjo{
		config: config,
		routes: CreateRoutes(),
		parser: Parser{},
		logger: CreateLogger(),
	}
}

// Get function
// For handling GET Requests
//
// Params:
// - url     {string} HTTP Request URL
// - closure {func(r Request) Response} Closure for handling HTTP Request
//
// Response:
// - None
//
func (banjo Banjo) Get(url string, closure func(ctx *Context)) {
	banjo.routes.Push("GET", url, closure)
}

// Post function
// For handling POST Requests
//
// Params:
// - url     {string} HTTP Request URL
// - closure {func(r Request) Response} Closure for handling HTTP Request
//
// Response:
// - None
//
func (banjo Banjo) Post(url string, closure func(ctx *Context)) {
	banjo.routes.Push("POST", url, closure)
}

// Put function
// For handling PUT Requests
//
// Params:
// - url     {string} HTTP Request URL
// - closure {func(r Request) Response} Closure for handling HTTP Request
//
// Response:
// - None
//
func (banjo Banjo) Put(url string, closure func(ctx *Context)) {
	banjo.routes.Push("PUT", url, closure)
}

// Patch function
// For handling PATCH Requests
//
// Params:
// - url     {string} HTTP Request URL
// - closure {func(r Request) Response} Closure for handling HTTP Request
//
// Response:
// - None
//
func (banjo Banjo) Patch(url string, closure func(ctx *Context)) {
	banjo.routes.Push("PATCH", url, closure)
}

// Options function
// For handling OPTIONS Requests
//
// Params:
// - url     {string} HTTP Request URL
// - closure {func(r Request) Response} Closure for handling HTTP Request
//
// Response:
// - None
//
func (banjo Banjo) Options(url string, closure func(ctx *Context)) {
	banjo.routes.Push("OPTIONS", url, closure)
}

// Head function
// For handling HEAD Requests
//
// Params:
// - url     {string} HTTP Request URL
// - closure {func(r Request) Response} Closure for handling HTTP Request
//
// Response:
// - None
//
func (banjo Banjo) Head(url string, closure func(ctx *Context)) {
	banjo.routes.Push("HEAD", url, closure)
}

// Delete function
// For handling DELETE Requests
//
// Params:
// - url     {string} HTTP Request URL
// - closure {func(r Request) Response} Closure for handling HTTP Request
//
// Response:
// - None
//
func (banjo Banjo) Delete(url string, closure func(ctx *Context)) {
	banjo.routes.Push("DELETE", url, closure)
}

// Run function
//
// Application starts listening for the requests,
// Parse them and runs code save id routes
// fields as {map[string]func(r Request) Response}
// All closures should return Response, you can
// use banjo.JSON, banjo.HTML etc methods
// All return your own Response struct
//
// This is last methods, that should called
// in the end of the application
//
// Params:
// - None
//
// Response:
// - None
//
func (banjo Banjo) Run() {
	banjo.logger.Info(fmt.Sprintf("BANjO.RUN Started PORT=%v", banjo.config.port))

	server, err := net.Listen("tcp", banjo.config.host+":"+banjo.config.port)

	if err != nil {
		banjo.logger.Critical("Error while trying to create connection")
		panic(err)
	}

	defer server.Close()

	for {
		conn, err := server.Accept()

		if err != nil {
			str := fmt.Sprintf("Error while trying to accept incomming connection:\nError: %v", err)
			banjo.logger.Error(str)
			continue
		}

		go banjo.handleRequest(conn)
	}
}

// handleRequest function
//
// Accept connection for future processing
// Methid handle each request, run closure if it
// available in routes table
// and return response for each reqeust
//
// Params:
// - conn {net.Conn} listener connection struct
//
// Response:
// - None
//
func (banjo Banjo) handleRequest(conn net.Conn) {
	data := make([]byte, 2048)

	_, err := conn.Read(data)

	if err != nil {
		str := fmt.Sprintf("Error while reading request data:\nError: %v", err)
		banjo.logger.Error(str)
		return
	}

	bytes := make([]byte, 2048)
	for i, v := range data {
		bytes[i] = byte(v)
	}

	ctx := Context{
		Request:  banjo.parser.Request(string(bytes)),
		Response: Response{},
	}

	action := banjo.routes.Block(ctx.Request.Method, ctx.Request.URL)
	action(&ctx)

	addRequiredHeaders(&ctx.Response)

	logLine := strings.Join([]string{ctx.Request.Method, "request to", ctx.Request.URL, strconv.Itoa(ctx.Response.Status)}, " ")
	banjo.logger.Info(logLine)

	responseRaw := banjo.parser.Response(ctx.Response)

	conn.Write([]byte(responseRaw))
	conn.Close()
}

// addRequiredHeaders function
//
// Added required headers for response {Response}
//
// Params:
// - data {*Response} pointer to Response struct
//
// Response:
// - None
//
func addRequiredHeaders(data *Response) {
	if data.Headers == nil {
		data.Headers = make(map[string]string)
	}

	data.Headers["Content-Length"] = strconv.Itoa(len(data.Body))
	data.Headers["Connection"] = "Closed"
	data.Headers["Data"] = time.Now().String()

	if data.Status == 0 {
		data.Status = 200
	}
}
