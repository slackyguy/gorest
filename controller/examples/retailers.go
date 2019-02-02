package examples

import (
	"github.com/slackyguy/gopay/dao"
	"github.com/slackyguy/gorest/controller"
)

type retailers struct {
	controller.Controller
}

func newRetailersController(collection string) func(
	baseController *controller.Controller) controller.Interface {
	return func(baseController *controller.Controller) controller.Interface {
		ctrl := new(retailers)
		ctrl.Controller = *baseController
		ctrl.Model = dao.Retailer{}
		ctrl.MapModel = make(map[string]dao.Retailer)
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
