package routing

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/slackyguy/gorest/controller"
)

// MuxRouter router using gorilla/mux
type MuxRouter struct {
	router *mux.Router
}

// NewMuxRouter creates a new MuxRouter
func NewMuxRouter() *MuxRouter {
	return &MuxRouter{router: mux.NewRouter()}
}

// RegisterRestfulHandlers registers route for a given context path
func (mux *MuxRouter) RegisterRestfulHandlers(
	path string,
	factory func(*controller.Controller) controller.Interface) Router {

	mux.router.HandleFunc(
		fmt.Sprintf("/%s", path),
		NewRequestHandler(path, mux, factory,
			func(controller controller.Interface) { controller.List() },
		)).Methods("GET")

	mux.router.HandleFunc(
		fmt.Sprintf("/%s/{id}", path),
		NewRequestHandler(path, mux, factory,
			func(controller controller.Interface) { controller.Get() },
		)).Methods("GET")

	mux.router.HandleFunc(
		fmt.Sprintf("/%s", path),
		NewRequestHandler(path, mux, factory,
			func(controller controller.Interface) { controller.Post() },
		)).Methods("POST")

	mux.router.HandleFunc(
		fmt.Sprintf("/%s/{id}", path),
		NewRequestHandler(path, mux, factory,
			func(controller controller.Interface) { controller.Put() },
		)).Methods("PUT")

	mux.router.HandleFunc(
		fmt.Sprintf("/%s/{id}", path),
		NewRequestHandler(path, mux, factory,
			func(controller controller.Interface) { controller.Delete() },
		)).Methods("DELETE")

	return mux
}

// RegisterHandler - register a simple handler
func (mux *MuxRouter) RegisterHandler(path string, handler func(
	response http.ResponseWriter, request *http.Request)) Router {
	mux.router.HandleFunc(path, handler)

	return mux
}

// Handle routes
func (mux *MuxRouter) Handle() {
	http.Handle("/", mux.router)
}
