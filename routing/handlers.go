package routing

import (
	"fmt"
	"net/http"

	"github.com/slackyguy/gorest/controller"
)

// BasicRouter basic router based on http.handler
type BasicRouter struct{}

// NewBasicRouter returns a BasicRouter ()
func NewBasicRouter() *BasicRouter {
	return new(BasicRouter)
}

// RegisterRestfulHandlers registers route for a given context path
func (router *BasicRouter) RegisterRestfulHandlers(
	path string,
	factory func(*controller.Controller) controller.Interface) Router {
	contextPath := fmt.Sprintf("/%s", path)

	router.RegisterHandler(fmt.Sprintf("/%s/", path), func(response http.ResponseWriter, request *http.Request) {

		var action func(response http.ResponseWriter, request *http.Request)

		switch request.Method {
		case "GET":
			action = NewRequestHandler(path, router, factory,
				func(controller controller.Interface) { controller.Get() })
			break
		case "PUT":
			action = NewRequestHandler(path, router, factory,
				func(controller controller.Interface) { controller.Put() })
			break
		case "DELETE":
			action = NewRequestHandler(path, router, factory,
				func(controller controller.Interface) { controller.Delete() })
			break
		}

		action(response, request)
	})

	router.RegisterHandler(contextPath, func(response http.ResponseWriter, request *http.Request) {

		var action func(response http.ResponseWriter, request *http.Request)

		switch request.Method {
		case "GET":
			action = NewRequestHandler(path, router, factory,
				func(controller controller.Interface) { controller.List() })
			break
		case "POST":
			action = NewRequestHandler(path, router, factory,
				func(controller controller.Interface) { controller.Post() })
			break
		}

		action(response, request)
	})

	return router
}

// RegisterHandler - register a simple handler
func (router *BasicRouter) RegisterHandler(path string, handler func(
	response http.ResponseWriter, request *http.Request)) Router {
	http.HandleFunc(path, handler)
	return router
}

// Handle register all routes (nothing to do in this case)
func (router *BasicRouter) Handle() {}
