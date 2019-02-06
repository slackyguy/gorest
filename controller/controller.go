package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/slackyguy/gorest/persistence"
	"github.com/slackyguy/gorest/persistence/firebaserealtime"
)

// Controller base type
type Controller struct {
	Context  context.Context
	Response http.ResponseWriter
	Request  *http.Request

	MessageID string
	Message   string

	Repository persistence.Interface
	Model      interface{}
	MapModel   interface{}
}

// Setup binds parameters and load resources
func (ctrl *Controller) Setup(collection string) {

	ctrl.Repository = &firebaserealtime.Repository{
		Repository: persistence.Repository{
			Context: ctrl.Context,
		},
	}

	ctrl.Repository.SetCollectionName(collection)
}

// Interface provides common interface for rest constroller
type Interface interface {
	// Get API handler
	Get()
	// List API handler
	List()
	// Post API handler
	Post()
	// Put API handler
	Put()
	// Delete API handler
	Delete()
	// BaseController instance
	BaseController() *Controller
}

// Get API handler
func (ctrl *Controller) Get() {
	ctrl.Repository.Find(ctrl.MessageID, &ctrl.Model)
	json.NewEncoder(ctrl.Response).Encode(&ctrl.Model)
}

// List API handler
func (ctrl *Controller) List() {
	ctrl.Repository.List(&ctrl.MapModel)
	json.NewEncoder(ctrl.Response).Encode(&ctrl.MapModel)
}

// Post API handler
func (ctrl *Controller) Post() {
	json.Unmarshal([]byte(ctrl.Message), &ctrl.Model)
	key := ctrl.Repository.Create(&ctrl.Model)
	fmt.Fprintln(ctrl.Response, key)
}

// Put API handler
func (ctrl *Controller) Put() {
	json.Unmarshal([]byte(ctrl.Message), &ctrl.Model)
	ctrl.Repository.Update(ctrl.MessageID, &ctrl.Model)
	fmt.Fprintln(ctrl.Response, ctrl.MessageID)
}

// Delete API handler
func (ctrl *Controller) Delete() {
	ctrl.Repository.Delete(ctrl.MessageID)
}

// BaseController instance
func (ctrl *Controller) BaseController() *Controller {
	return ctrl
}

// ValidateBasicAuthentication provides basic authentication method
func (ctrl *Controller) ValidateBasicAuthentication() bool {

	username, password, authOK := ctrl.Request.BasicAuth()

	if !authOK {
		//WWW-Authenticate: Basic
		//WWW-Authenticate: Basic realm="Access to ...", charset="UTF-8"
		ctrl.Response.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
		http.Error(ctrl.Response, "Unauthorized", http.StatusUnauthorized)

	} else if !validate(username, password) {
		http.Error(ctrl.Response, "Invalid Credentials", http.StatusUnauthorized)
	}

	return true
}

func validate(username string, password string) (sucess bool) {
	log.Println(username, password)

	return true
}
