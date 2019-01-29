package lib

import (
	"context"
	"net/http"
)

// Controller base struct type
type Controller struct {
	Context     context.Context
	AppSettings *AppSettings
}

// ControllerInterface provides common interface for rest constroller
type ControllerInterface interface {
	// Get API handler
	Get(w http.ResponseWriter, r *http.Request)
	// List API handler
	List(w http.ResponseWriter, r *http.Request)
	// Post API handler
	Post(w http.ResponseWriter, r *http.Request)
	// Delete API handler
	Delete(w http.ResponseWriter, r *http.Request)

	//put(w http.ResponseWriter, r *http.Request)
}
