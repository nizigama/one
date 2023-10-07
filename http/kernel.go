package http

import (
	"fmt"
	"net/http"
)

type Kernel struct {
	router    *Router
	port      int
	debugging bool
}

func NewKernel(debugging bool, serverPort int, router *Router) *Kernel {

	return &Kernel{
		router:    router,
		port:      serverPort,
		debugging: debugging,
	}
}

func (k *Kernel) Start() error {

	return http.ListenAndServe(fmt.Sprintf(":%d", k.port), k.router.chiRouter)
}
