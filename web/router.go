package web

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	. "github.com/nizigama/one/rendering"
	"github.com/nizigama/one/types"
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

	http.FileServer(http.Dir("./public"))
	router.chiRouter.Handle("/public/*", http.StripPrefix("/public", http.FileServer(http.Dir("./public"))))

	router.View("/", "welcome")
}

func (r *Router) View(path string, name string) {

	r.chiRouter.Get(path, func(writer http.ResponseWriter, request *http.Request) {
		request.ParseForm()

		requestData := &types.Request{Request: request}

		View(name, requestData).SetWriter(writer).WriteStatus().Write()
	})
}

func (r *Router) Get(path string, handler types.Handler) {

	r.chiRouter.Get(path, func(writer http.ResponseWriter, request *http.Request) {
		response := handler(&types.Request{request})

		response.SetWriter(writer).WriteStatus().Write()
	})
}

func (r *Router) Post(path string, handler types.Handler) {

	r.chiRouter.Post(path, func(writer http.ResponseWriter, request *http.Request) {
		response := handler(&types.Request{request})

		response.SetWriter(writer).WriteStatus().Write()
	})
}

func (r *Router) Put(path string, handler types.Handler) {

	r.chiRouter.Put(path, func(writer http.ResponseWriter, request *http.Request) {
		response := handler(&types.Request{request})

		response.SetWriter(writer).WriteStatus().Write()
	})
}

func (r *Router) Delete(path string, handler types.Handler) {

	r.chiRouter.Delete(path, func(writer http.ResponseWriter, request *http.Request) {
		response := handler(&types.Request{request})

		response.SetWriter(writer).WriteStatus().Write()
	})
}
