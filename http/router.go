package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

type Router struct {
	chiRouter *chi.Mux
}

func NewRouter(debugging bool) *Router {

	chiRouter := chi.NewRouter()

	if debugging {
		chiRouter.Use(middleware.Logger)
	}

	return &Router{
		chiRouter: chiRouter,
	}
}

func DefaultRoutes(router *Router) {

	router.Get("/", func(request *Request) *Response {

		return &Response{
			Status: http.StatusBadRequest,
			Headers: map[string]string{
				"Content-Type": "text/html",
			},
			Content: "<h1>Welcome to the default One & Only</h1>",
		}
	})
}

func (r *Router) Get(path string, handler Handler) {

	r.chiRouter.Get(path, func(writer http.ResponseWriter, request *http.Request) {
		response := handler(&Request{request})

		for key, value := range response.Headers {

			writer.Header().Set(key, value)
		}

		writer.WriteHeader(response.Status)

		writer.Write([]byte(response.Content))
	})
}

func (r *Router) Post(path string, handler Handler) {

	r.chiRouter.Post(path, func(writer http.ResponseWriter, request *http.Request) {
		response := handler(&Request{request})

		for key, value := range response.Headers {

			writer.Header().Set(key, value)
		}

		writer.WriteHeader(response.Status)

		writer.Write([]byte(response.Content))
	})
}

func (r *Router) Put(path string, handler Handler) {

	r.chiRouter.Put(path, func(writer http.ResponseWriter, request *http.Request) {
		response := handler(&Request{request})

		for key, value := range response.Headers {

			writer.Header().Set(key, value)
		}

		writer.WriteHeader(response.Status)

		writer.Write([]byte(response.Content))
	})
}

func (r *Router) Delete(path string, handler Handler) {

	r.chiRouter.Delete(path, func(writer http.ResponseWriter, request *http.Request) {
		response := handler(&Request{request})

		for key, value := range response.Headers {

			writer.Header().Set(key, value)
		}

		writer.WriteHeader(response.Status)

		writer.Write([]byte(response.Content))
	})
}
