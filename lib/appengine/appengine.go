package appengine

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/slackyguy/gorest/lib"
	"google.golang.org/appengine"
)

// Register registers route for a given context path
func Register(
	path string,
	settings *lib.AppSettings,
	router *mux.Router,
	factory func(context context.Context, settings *lib.AppSettings) lib.ControllerInterface) {

	router.HandleFunc(fmt.Sprintf("/%s", path), RegisterGetList(path, settings, router, factory)).Methods("GET")
	router.HandleFunc(fmt.Sprintf("/%s/{id}", path), RegisterGet(path, settings, router, factory)).Methods("GET")
	router.HandleFunc(fmt.Sprintf("/%s/{id}", path), RegisterPost(path, settings, router, factory)).Methods("POST")
	router.HandleFunc(fmt.Sprintf("/%s/{id}", path), RegisterDelete(path, settings, router, factory)).Methods("DELETE")
}

// RegisterGet registers Get route
func RegisterGet(
	path string,
	settings *lib.AppSettings,
	router *mux.Router,
	factory func(context context.Context, settings *lib.AppSettings) lib.ControllerInterface,
) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		factory(requestContext(r), settings).Get(w, r)
	}
}

// RegisterGetList registers Get(List) route
func RegisterGetList(
	path string,
	settings *lib.AppSettings,
	router *mux.Router,
	factory func(context context.Context, settings *lib.AppSettings) lib.ControllerInterface,
) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		factory(requestContext(r), settings).List(w, r)
	}
}

// RegisterPost registers Post route
func RegisterPost(
	path string,
	settings *lib.AppSettings,
	router *mux.Router,
	factory func(context context.Context, settings *lib.AppSettings) lib.ControllerInterface,
) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		factory(requestContext(r), settings).Post(w, r)
	}
}

// RegisterDelete registers Delete route
func RegisterDelete(
	path string,
	settings *lib.AppSettings,
	router *mux.Router,
	factory func(context context.Context, settings *lib.AppSettings) lib.ControllerInterface,
) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		factory(requestContext(r), settings).Delete(w, r)
	}
}

func requestContext(r *http.Request) context.Context {
	ctx := appengine.NewContext(r)
	//ctx := context.Background()
	return ctx
}
