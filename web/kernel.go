package web

import (
	"fmt"
	. "github.com/nizigama/one/routing"
	"net/http"
)

// Kernel Will be in charge of loading router, sessions, middlewares before they are used by handlers or views
type Kernel struct {
	router    *Router
	port      int
	debugging bool
}

func NewKernel(debugging bool, serverPort int) *Kernel {

	return &Kernel{
		router:    NewRouter(debugging),
		port:      serverPort,
		debugging: debugging,
	}
}

func (k *Kernel) Boot(ps ...PathResolver) error {

	http.FileServer(http.Dir("./public"))
	k.router.Mux.Handle("/public/*", http.StripPrefix("/public", http.FileServer(http.Dir("./public"))))

	k.router.View("/", "welcome")

	for _, pathResolver := range ps {
		pathResolver(k.router)
	}

	return http.ListenAndServe(fmt.Sprintf(":%d", k.port), k.router)
}
