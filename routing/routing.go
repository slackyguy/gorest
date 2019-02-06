package routing

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/slackyguy/gorest/base"
	"github.com/slackyguy/gorest/controller"
)

// HTTP http router interface
var HTTP Router

// Router interface
type Router interface {
	RegisterRestfulHandlers(path string,
		factory func(*controller.Controller) controller.Interface) Router
	RegisterHandler(path string, handler func(
		response http.ResponseWriter, request *http.Request)) Router
	Handle()
}

// NewRequestHandler creates a request handler bound to an action
func NewRequestHandler(
	path string,
	router Router,
	factory func(*controller.Controller) controller.Interface,
	action func(controller controller.Interface),
) func(response http.ResponseWriter, request *http.Request) {
	return func(response http.ResponseWriter, request *http.Request) {
		validate(router, response, request, factory, action)
	}
}

func validate(
	router Router,
	response http.ResponseWriter,
	request *http.Request,
	factory func(*controller.Controller) controller.Interface,
	action func(controller.Interface)) {

	// TODO Validar autenticação do usuário
	if request.Header.Get("Content-Type") != "application/json" &&
		(request.Method == "PUT" || request.Method == "POST") {
		InvalidContentType(response, request)
	}

	contextFactory := base.ApplicationSettings.ContextFactory
	context := contextFactory(request)
	controller := &controller.Controller{
		Context:  context,
		Request:  request,
		Response: response}

	// TODO Mover para a action de "login" apenas.
	// Para os demais, a validação deve ser de OAuth
	// if !controller.BaseController().ValidateBasicAuthentication() {
	// 	return
	// }

	target := factory(controller)
	LoadParameters(request, controller)
	action(target)
}

// InvalidContentType - Action Filter for "Not Accepted ContentType".
func InvalidContentType(response http.ResponseWriter, request *http.Request) {
	// TODO Alterar o código de retorno
	fmt.Fprintln(response, "InvalidContentType")
}

// LoadParameters loads parameters to controller
func LoadParameters(request *http.Request, controller *controller.Controller) {

	if request.Method == "GET" ||
		request.Method == "PUT" ||
		request.Method == "DELETE" {

		path := strings.Split(
			strings.Trim(request.URL.Path, "/"), "/")
		if len(path) > 1 {
			controller.MessageID = path[1]
		}

	}
	if request.Method == "PUT" || request.Method == "POST" {

		bytes, _ := ioutil.ReadAll(request.Body)
		controller.Message = string(bytes)
	}
}
