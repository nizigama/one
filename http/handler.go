package http

import "net/http"

type Request struct {
	*http.Request
}

type Response struct {
	Content string
	Status  int
	Headers map[string]string
}

type Handler func(request *Request) *Response
