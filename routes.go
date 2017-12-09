package banjo

import (
	"reflect"
)

// Routes struct
//
// Used for storing users closures
// The struct have one map for each method table
// table type: map[string]func(ctx *Context)
//
type Routes struct {
	GET     map[string]func(ctx *Context)
	POST    map[string]func(ctx *Context)
	PUT     map[string]func(ctx *Context)
	PATCH   map[string]func(ctx *Context)
	OPTIONS map[string]func(ctx *Context)
	HEAD    map[string]func(ctx *Context)
	DELETE  map[string]func(ctx *Context)
}

// CreateRoutes function
//
// Create Routes with empty method tables
//
// Params:
// - None
//
// Response:
// - routes {Routes} Routes struct with empty map fields
//
func CreateRoutes() Routes {
	return Routes{
		GET:     make(map[string]func(ctx *Context)),
		POST:    make(map[string]func(ctx *Context)),
		PUT:     make(map[string]func(ctx *Context)),
		PATCH:   make(map[string]func(ctx *Context)),
		OPTIONS: make(map[string]func(ctx *Context)),
		HEAD:    make(map[string]func(ctx *Context)),
		DELETE:  make(map[string]func(ctx *Context)),
	}
}

// Block function
//
// Returns func(request Request) Response closure
//
// Params:
// - method {string} HTTP Request Method
// - url    {string} HTTP Request URL
//
// Response:
// - closure {func(request Request) Response} Returns closure with user
//
func (routes Routes) Block(method string, url string) func(ctx *Context) {
	object := reflect.ValueOf(routes)
	value := reflect.Indirect(object).FieldByName(method)
	intf := value.Interface()
	table := intf.(map[string]func(ctx *Context))

	block, ok := table[url]
	if ok {
		return block
	}

	return notFound()
}

// Push function
//
// Adding new element to one of fields in Routes struct
// Element should be passed by url {string},
// value type of {func(request Request) Response}
//
// Params:
// - method  {string} HTTP Request Method
// - url		 {string} HTTP Request URL
// - closure {func(request Request) Response}
//
// Response:
// - None
//
func (routes Routes) Push(method string, url string, closure func(ctx *Context)) {
	object := reflect.ValueOf(routes)
	object = reflect.Indirect(object)
	value := object.FieldByName(method)
	value.SetMapIndex(reflect.ValueOf(url), reflect.ValueOf(closure))
}

// notFound function
//
// Returns default error page Response if
// closure with given method & url are not available
//
// Params:
// - None
//
// Response:
// - closure {func(request Request) Response} Default error closure
//
func notFound() func(ctx *Context) {
	return func(ctx *Context) {
		ctx.Response.Body = "Page Not Found"
		ctx.Response.Status = 404
	}
}
