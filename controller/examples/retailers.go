package examples

import (
	"net/http"

	"github.com/slackyguy/gorest/controller"
	"github.com/slackyguy/gorest/routing"
)

// Retailer is a json-serializable type.
type Retailer struct {
	BusinessID  string `json:"business_id,omitempty"`
	Description string `json:"description,omitempty"`
	GeoCode     string `json:"geocode,omitempty"`
}

type retailers struct {
	controller.Controller
}

func newRetailersController(collection string) func(
	baseController *controller.Controller) controller.Interface {
	return func(baseController *controller.Controller) controller.Interface {
		ctrl := new(retailers)
		ctrl.Controller = *baseController
		ctrl.Model = Retailer{}
		ctrl.MapModel = make(map[string]Retailer)
		ctrl.Setup(collection)

		return ctrl
	}
}

// Get API handler
func (ctrl *retailers) Get() {
	ctrl.Controller.Get()
}

// List API handler
func (ctrl *retailers) List() {
	ctrl.Controller.List()
}

// Post API handler
func (ctrl *retailers) Post() {
	ctrl.Controller.Post()
}

// Put API handler
func (ctrl *retailers) Put() {
	ctrl.Controller.Put()
}

// Delete API handler
func (ctrl *retailers) Delete() {
	ctrl.Controller.Delete()
}

// BaseController instance
func (ctrl *retailers) BaseController() *controller.Controller {
	return &ctrl.Controller
}

func init() {
	routing.HTTP = routing.NewBasicRouter()
	//routing.HTTP = routing.NewMuxRouter()

	routing.HTTP.RegisterHandler(
		"/", func(response http.ResponseWriter, request *http.Request) {
			if request.Method != "GET" {
				http.Error(response, "Method not allowed", http.StatusMethodNotAllowed)
				return
			}

			http.ServeFile(response, request, "home.html")
		}).RegisterRestfulHandlers(
		"retailers", newRetailersController("retailers"))
}

// DoNothing (NOP)
func DoNothing() {}
