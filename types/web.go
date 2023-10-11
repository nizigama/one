package types

import (
	"net/http"
)

type Request struct {
	*http.Request
}

type Handler func(request *Request) WebWriter

type WebWriter interface {
	Write() (int, error)
	WriteStatus() WebWriter
	SetHeaders(map[string]string) WebWriter
	SetWriter(http.ResponseWriter) WebWriter
	SetStatus(code int)
}
