package banjo

import "testing"

func TestRoutesPushFunc(t *testing.T) {
	routes := CreateRoutes()
	routes.Push("GET", "/foo", func(ctx *Context) {
		ctx.Response = Response{}
	})

	if routes.GET == nil {
		t.Errorf("GET table should be different")
	}
}

func TestNotFoundResponseFunc(t *testing.T) {
	routes := CreateRoutes()
	action := routes.Block("GET", "/foo")
	ctx := &Context{}
	action(ctx)

	if ctx.Response.Status != 404 {
		t.Errorf("Response Status should be 404")
	}
}

func TestFullRoutesUsageStruct(t *testing.T) {
	routes := CreateRoutes()
	routes.Push("GET", "/foo", func(ctx *Context) {
		ctx.Response.Status = 200
	})

	action := routes.Block("GET", "/foo")
	ctx := &Context{}
	action(ctx)

	if ctx.Response.Status != 200 {
		t.Errorf("Response Status should be 200")
	}
}
