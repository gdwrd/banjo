package banjo

import "testing"

func TestHTTPRequestHeadersParsing(t *testing.T) {
	p := Parser{}
	rawRequest := "GET /foo HTTP/1.1\r\nContent-Type: application/json; charset=utf-8\r\nAccept: application/json\r\n\r\n"
	request := p.Request(rawRequest)

	if request.Method != "GET" {
		t.Errorf("Request should be GET")
	}
	if request.URL != "/foo" {
		t.Errorf("Request should be GET")
	}
	if request.Headers["Content-Type"] != "application/json; charset=utf-8" {
		t.Errorf("`Content-Type` value should be `application/json`")
	}
}

func TestHTTPRequestJSONParamsParsing(t *testing.T) {
	p := Parser{}
	rawRequest := "POST /foo HTTP/1.1\r\nContent-Type: application/json\r\n\r\n{\"foo\":\"bar\"}"
	request := p.Request(rawRequest)

	if request.Params != "{\"foo\":\"bar\"}" {
		t.Errorf("Param value should be {\"foo\":\"bar\"}")
	}
}

func TestHTTPRequestFormDataParamsParsing(t *testing.T) {
	p := Parser{}
	rawRequest := "POST /foo HTTP/1.1\r\nContent-Type: application/x-www-form-urlencoded\r\n\r\nfoo=bar&bar=foo"
	request := p.Request(rawRequest)

	if request.MapParams["foo"] != "bar" {
		t.Errorf("Param `foo` value should be `bar`")
	}
	if request.MapParams["bar"] != "foo" {
		t.Errorf("Param `bar` value should be `foo`")
	}
}

func TestHTTPRequestFormDataParamParsing(t *testing.T) {
	p := Parser{}
	rawRequest := "POST /foo HTTP/1.1\r\nContent-Type: application/x-www-form-urlencoded\r\n\r\nfoo=bar"
	request := p.Request(rawRequest)

	if request.MapParams["foo"] != "bar" {
		t.Errorf("Param `foo` value should be `bar`")
	}
}

func TestHTTPRequestMultipartParamParsing(t *testing.T) {
	p := Parser{}
	rawRequest := "POST /foo HTTP/1.1\r\nContent-Type: multipart/form-data; boundary=----11111\r\n\r\n----11111Content-Disposition: form-data; name=\"foo\"\r\n\r\nbar\r\n----11111\r\nContent-Disposition: form-data; name=\"file\"; filename=\"bar.txt\"\r\n\r\nThis is textfile content\r\n----11111--"
	request := p.Request(rawRequest)

	if request.MapParams["foo"] != "bar" {
		t.Errorf("Param `foo` value should be `bar`")
	}
}

func TestHTTPResponseJSONBodyParsing(t *testing.T) {
	p := Parser{}
	str := "HTTP/1.1 200\r\nContent-Type: application/json\r\n\r\n{\"foo\":\"bar\"}"
	rawResponse := p.Response(Response{
		Headers: map[string]string{"Content-Type": "application/json"},
		Status:  200,
		Body:    "{\"foo\":\"bar\"}",
	})

	if str != rawResponse {
		t.Errorf("Requests should be same")
	}
}
