package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/slackyguy/gorest/base"
	"github.com/slackyguy/gorest/persistence"
	"github.com/slackyguy/gorest/persistence/firebaserealtime"
)

// Controller base type
type Controller struct {
	Context  context.Context
	Response http.ResponseWriter
	Request  *http.Request

	AppSettings *base.AppSettings
	MessageID   string
	Message     string

	Repository persistence.Interface
	Model      interface{}
	MapModel   interface{}
}

func (ctrl *Controller) load() {

	if ctrl.Request.Method == "GET" ||
		ctrl.Request.Method == "PUT" ||
		ctrl.Request.Method == "DELETE" {

		// for x := range ctrl.Request.URL.Query() {
		// 	log.Println(x)
		// }

		path := strings.Split(
			strings.Trim(ctrl.Request.URL.Path, "/"), "/")
		if len(path) > 1 {
			ctrl.MessageID = path[1]
		}

	}
	if ctrl.Request.Method == "PUT" || ctrl.Request.Method == "POST" {

		bytes, _ := ioutil.ReadAll(ctrl.Request.Body)
		ctrl.Message = string(bytes)

	}
}

// Setup binds parameters and load resources
func (ctrl *Controller) Setup(collection string) {
	ctrl.load()

	ctrl.Repository = &firebaserealtime.Repository{
		Repository: persistence.Repository{
			AppSettings: ctrl.AppSettings,
			Context:     ctrl.Context,
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
