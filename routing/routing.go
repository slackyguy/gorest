package routing

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/slackyguy/gorest/base"
	"github.com/slackyguy/gorest/controller"
	"google.golang.org/appengine"
)

var appSettings *base.AppSettings
var router *mux.Router

// Register registers route for a given context path
func Register(
	path string,
	factory func(*controller.Controller) controller.Interface) {

	if appSettings == nil {
		appSettings = base.ReadFromFile("goapp.properties")
	}

	if router == nil {
		router = mux.NewRouter()
	}

	router.HandleFunc(
		fmt.Sprintf("/%s", path),
		RegisterAction(path, factory,
			func(controller controller.Interface) { controller.List() },
		)).Methods("GET")

	router.HandleFunc(
		fmt.Sprintf("/%s/{id}", path),
		RegisterAction(path, factory,
			func(controller controller.Interface) { controller.Get() },
		)).Methods("GET")

	router.HandleFunc(
		fmt.Sprintf("/%s", path),
		RegisterAction(path, factory,
			func(controller controller.Interface) { controller.Post() },
		)).Methods("POST")

	router.HandleFunc(
		fmt.Sprintf("/%s/{id}", path),
		RegisterAction(path, factory,
			func(controller controller.Interface) { controller.Put() },
		)).Methods("PUT")

	router.HandleFunc(
		fmt.Sprintf("/%s/{id}", path),
		RegisterAction(path, factory,
			func(controller controller.Interface) { controller.Delete() },
		)).Methods("DELETE")
}

// Routes return configured routes
func Routes() *mux.Router {
	return router
}

// RegisterAction - Generic action register
func RegisterAction(
	path string,
	factory func(*controller.Controller) controller.Interface,
	action func(controller controller.Interface),
) func(response http.ResponseWriter, request *http.Request) {
	return func(response http.ResponseWriter, request *http.Request) {
		validate(response, request, factory, action)
	}
}

func validate(
	response http.ResponseWriter,
	request *http.Request,
	factory func(*controller.Controller) controller.Interface,
	action func(controller.Interface)) {

	// TODO Validar autenticação do usuário
	if request.Header.Get("Content-Type") != "application/json" &&
		(request.Method == "PUT" || request.Method == "POST") {
		InvalidContentType(response, request)
	}

	context := requestContext(request)
	controller := &controller.Controller{
		Context:     context,
		AppSettings: appSettings,
		Request:     request,
		Response:    response}

	target := factory(controller)
	action(target)
}

// InvalidContentType - Action Filter for "Not Accepted ContentType".
func InvalidContentType(response http.ResponseWriter, request *http.Request) {
	// TODO Alterar o código de retorno
	fmt.Fprintln(response, "InvalidContentType")
}

func requestContext(request *http.Request) context.Context {
	context := appengine.NewContext(request)
	//context := context.Background()
	return context
}
