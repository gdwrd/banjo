package banjo

import (
	"bytes"
	"net/http"
	"testing"
	"time"
)

func TestBanjoGetRequestFunc(t *testing.T) {
	cnf := DefaultConfig()
	app := Create(cnf)

	app.Get("/foo", func(ctx *Context) {
		ctx.Response.Status = 201
	})

	ctx := &Context{}
	app.routes.GET["/foo"](ctx)

	if ctx.Response.Status != 201 {
		t.Errorf("Response Status should be 201")
	}
}

func TestBanjoPostRequestFunc(t *testing.T) {
	app := Create(DefaultConfig())

	app.Post("/foo", func(ctx *Context) {
		ctx.Response.Status = 201
	})

	ctx := &Context{Response: Response{}}
	app.routes.Block("POST", "/foo")(ctx)

	if ctx.Response.Status != 201 {
		t.Errorf("Response Status should be 201")
	}
}

func TestBanjoPutRequestFunc(t *testing.T) {
	cnf := DefaultConfig()
	app := Create(cnf)

	app.Put("/foo", func(ctx *Context) {
		ctx.Response.Status = 201
	})

	ctx := &Context{}
	app.routes.Block("PUT", "/foo")(ctx)

	if ctx.Response.Status != 201 {
		t.Errorf("Response Status should be 201")
	}
}

func TestBanjoPatchRequestFunc(t *testing.T) {
	cnf := DefaultConfig()
	app := Create(cnf)

	app.Patch("/foo", func(ctx *Context) {
		ctx.Response.Status = 201
	})

	ctx := &Context{}
	app.routes.Block("PATCH", "/foo")(ctx)

	if ctx.Response.Status != 201 {
		t.Errorf("Response Status should be 201")
	}
}

func TestBanjoOptionsRequestFunc(t *testing.T) {
	cnf := DefaultConfig()
	app := Create(cnf)

	app.Options("/foo", func(ctx *Context) {
		ctx.Response.Status = 201
	})

	ctx := &Context{}
	app.routes.Block("OPTIONS", "/foo")(ctx)

	if ctx.Response.Status != 201 {
		t.Errorf("Response Status should be 201")
	}
}

func TestBanjoHeadRequestFunc(t *testing.T) {
	cnf := DefaultConfig()
	app := Create(cnf)

	app.Head("/foo", func(ctx *Context) {
		ctx.Response.Status = 201
	})

	ctx := &Context{}
	app.routes.Block("HEAD", "/foo")(ctx)

	if ctx.Response.Status != 201 {
		t.Errorf("Response Status should be 201")
	}
}

func TestBanjoDeleteRequestFunc(t *testing.T) {
	cnf := DefaultConfig()
	app := Create(cnf)

	app.Delete("/foo", func(ctx *Context) {
		ctx.Response.Status = 201
	})

	ctx := &Context{}
	app.routes.Block("DELETE", "/foo")(ctx)

	if ctx.Response.Status != 201 {
		t.Errorf("Response Status should be 201")
	}
}

func TestAddingRequiredHeadersToResponse(t *testing.T) {
	response := &Response{Headers: make(map[string]string)}
	addRequiredHeaders(response)

	if response.Status != 200 {
		t.Errorf("Default status should be 200")
	}
}

func TestJSONPreparingToResponseFunction(t *testing.T) {
	ctx := &Context{}
	ctx.JSON(M{"foo": "bar"})

	if ctx.Response.Status != 200 {
		t.Errorf("Status should be 200")
	}
	if ctx.Response.Body != "{\"foo\":\"bar\"}" {
		t.Errorf("Body should be {\"foo\":\"bar\"}")
	}
}

func TestJSONParsingErrorFunction(t *testing.T) {
	ctx := Context{}
	ctx.JSON(M{"foo": func(i int) int { return i }})

	if ctx.Response.Status != 500 {
		t.Errorf("Status should be 500")
	}

	if ctx.Response.Body != "Internal Server Error" {
		t.Errorf("Body should be `Internal Server Error`")
	}
}

func TestHTMLPreparingToResponseFunction(t *testing.T) {
	ctx := &Context{}
	ctx.HTML("<h1>Hello from Banjo</h1>")

	if ctx.Response.Headers["Content-Type"] != "text/html" {
		t.Errorf("Content-Type should be text/html")
	}
}

func TestRedirectToPreparingToResponseFunction(t *testing.T) {
	ctx := &Context{}
	ctx.RedirectTo("/admin")

	if ctx.Response.Headers["Location"] != "/admin" {
		t.Errorf("Location should be /admin")
	}

	if ctx.Response.Status != 301 {
		t.Errorf("Status should be 301")
	}
}

func TestBanjoRun(t *testing.T) {
	banjo := Create(DefaultConfig())
	banjo.Get("/foo", func(ctx *Context) {
		ctx.JSON(M{"message": "HELLO!"})
	})

	go banjo.Run()

	time.Sleep(time.Second * 1)

	req, err := http.NewRequest("GET", "http://localhost:4321/foo", bytes.NewBuffer([]byte("")))
	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		t.Errorf("NESMOGLA")
	}

	if resp.Status != "200" {
		t.Errorf("Status should be 200")
	}
}
