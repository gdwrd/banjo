package banjo

import (
	"bytes"
	"fmt"
	"log"
	"strings"
)

// Parser struct
//
// Used for easy access to ParserIntr
// functions, this is empty struct
//
type Parser struct{}

// Separator is a default separator in HTTP Requests
const Separator = "\r\n"

// DubSeparator is a default duble separator in HTTP Requests
const DubSeparator = "\r\n\r\n"

// HTTPVersion it'a default HTTP version
const HTTPVersion = "HTTP/1.1"

// Request function for parsing Raw
// HTTP Request to banjo.Request struct
//
// Params:
// - data {string} Raw HTTP Request
//
// Response:
// - request {banjo.Request}
//
func (p Parser) Request(rawData string) Request {
	var rawH, rawB, method, url, httpVersion string

	if strings.Contains(rawData, DubSeparator) {
		data := strings.Split(rawData, DubSeparator)

		if len(data) == 2 {
			rawH, rawB = data[0], data[1]
		} else {
			rawH, rawB = data[0], strings.Join(data[1:len(data)], DubSeparator)
		}
	} else {
		rawH = rawData
	}

	arrH := strings.Split(rawH, Separator)

	if strings.Contains(arrH[0], "HTTP/1") {
		str := strings.Split(arrH[0], " ")
		method, url, httpVersion = str[0], str[1], str[2]
	}

	headers := parseHeaders(arrH)
	params, files := parseParams(rawB, headers["Content-Type"])

	return Request{
		Headers:     headers,
		Params:      rawB,
		Files:       files,
		MapParams:   params,
		Method:      method,
		URL:         url,
		HTTPVersion: httpVersion,
	}
}

// Response prepared banjo.Response struct to
// Raw HTTP Response string
//
// Params:
// - data {banjo.Response} prepared banjo.Response struct
//
// Response:
// - response {string} Raw HTTP Response string
//
func (p Parser) Response(data Response) string {
	var buffer bytes.Buffer

	buffer.WriteString(fmt.Sprintf("%s %d\r\n", HTTPVersion, data.Status))

	for k, v := range data.Headers {
		buffer.WriteString(strings.Join([]string{k, v}, ": "))
		buffer.WriteString(Separator)
	}

	buffer.WriteString(Separator)
	buffer.WriteString(data.Body)

	return buffer.String()
}

// parseHeaders function
//
// Allows you to parse []string with request
// headers to map[string]string
//
// Params:
// - data {[]string} headers as array of strings
//
// Response:
// - data {map[string]string}
//
func parseHeaders(data []string) map[string]string {
	headers := make(map[string]string)

	for _, str := range data {
		if strings.Contains(str, ": ") {
			headerArr := strings.Split(str, ": ")
			k, v := headerArr[0], headerArr[1]

			if _, ok := headers[k]; ok {
				headers[k] = strings.Join([]string{headers[k], v}, "; ")
			} else {
				headers[k] = v
			}
		}
	}

	return headers
}

// parseParams function
//
// Allows you to parse Request params
// depending on Content-Type header
//
// Params:
// - data   {string} HTTP Request Body
// - cType {string} Content-Type header
//
// Response:
// - data   {map[string]string} Parsed body
//
func parseParams(data string, cType string) (map[string]string, []map[string]string) {
	params := make(map[string]string)
	files := []map[string]string{}
	var boundary string

	if strings.Contains(cType, "application/json") {
		return params, files
	} else if strings.Contains(cType, "application/x-www-form-urlencoded") {
		params = parseFormParams(data)
	} else if strings.Contains(cType, "multipart/form-data") {
		types := strings.Split(cType, "; ")

		if str := types[1]; str != "" {
			array := strings.Split(str, "=")

			if boundary = array[1]; boundary == "" {
				log.Fatal("Error while parsing `multipart/form-data` headers, boundary required")
			}
		} else {
			log.Fatal("Error while parsing `multipart/form-data` headers, boundary required")
		}

		params, files = parseMultipartParams(data, boundary)
	}

	return params, files
}

// parseFormParams function
//
// Function parse params for form-data content-type
//
// Params:
// - data {string}
//
// Response:
// - response {map[string]string}
//
func parseFormParams(data string) map[string]string {
	params := make(map[string]string)

	if strings.Contains(data, "&") {
		elements := strings.Split(data, "&")

		for _, item := range elements {
			if strings.Contains(item, "=") {
				itemA := strings.Split(item, "=")
				k, v := itemA[0], itemA[1]
				params[k] = v
			}
		}
	} else {
		if strings.Contains(data, "=") {
			itemA := strings.Split(data, "=")
			k, v := itemA[0], itemA[1]
			params[k] = v
		}
	}

	return params
}

// parseMultipartParams function
//
// Function parse `multipart/form-data` params
// to map[string]string and []map[string]string,
// First attribute is params map, Second is sended files
//
// Params:
// - data     {string} HTTP Request Body string
// - boundary {string} Multipart Boundary
//
// Response:
// - params {map[string]string}   parsed params
// - files  {[]map[string]string} parsed files struct
//
func parseMultipartParams(data string, boundary string) (map[string]string, []map[string]string) {
	params := make(map[string]string)
	files := []map[string]string{}
	items := strings.Split(data, boundary)

	for _, item := range items {
		if item == "" || item == "--" {
			continue
		}
		file := make(map[string]string)

		array := strings.Split(item, DubSeparator)
		fieldHeader, fieldContent := array[0], array[1]
		array = strings.Split(fieldHeader, "; ")
		array = array[1:len(array)]

		if strings.Contains(fieldHeader, "filename") {
			for _, i := range array {
				a := strings.Split(i, "=")
				k, v := strings.Trim(a[0], "\""), a[1]
				file[k] = v
			}

			file["content"] = fieldContent
		} else {
			for _, i := range array {
				a := strings.Split(i, "=")
				v := strings.Trim(a[1], "\"")
				params[v] = strings.Trim(fieldContent, Separator)
			}
		}

		files = append(files, file)
	}

	return params, files
}
