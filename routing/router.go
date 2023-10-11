package routing

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/nizigama/one/rendering"
	"github.com/nizigama/one/types"
	"net/http"
)

type Router struct {
	*chi.Mux
}

type PathResolver func(router *Router)

func NewRouter(debugging bool) *Router {

	chiRouter := chi.NewRouter()

	if debugging {
		chiRouter.Use(middleware.Logger)
	}

	return &Router{
		chiRouter,
	}
}

func (r *Router) View(path string, name string) {

	r.Mux.Get(path, func(writer http.ResponseWriter, request *http.Request) {
		request.ParseForm()

		requestData := &types.Request{Request: request}

		rendering.View(name, requestData).SetWriter(writer).WriteStatus().Write()
	})
}

func (r *Router) Get(path string, handler types.Handler) {

	r.Mux.Get(path, func(writer http.ResponseWriter, request *http.Request) {
		response := handler(&types.Request{request})

		response.SetWriter(writer).WriteStatus().Write()
	})
}

func (r *Router) Post(path string, handler types.Handler) {

	r.Mux.Post(path, func(writer http.ResponseWriter, request *http.Request) {
		response := handler(&types.Request{request})

		response.SetWriter(writer).WriteStatus().Write()
	})
}

func (r *Router) Put(path string, handler types.Handler) {

	r.Mux.Put(path, func(writer http.ResponseWriter, request *http.Request) {
		response := handler(&types.Request{request})

		response.SetWriter(writer).WriteStatus().Write()
	})
}

func (r *Router) Delete(path string, handler types.Handler) {

	r.Mux.Delete(path, func(writer http.ResponseWriter, request *http.Request) {
		response := handler(&types.Request{request})

		response.SetWriter(writer).WriteStatus().Write()
	})
}
