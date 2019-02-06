package main

import (
	"context"
	"log"
	"net/http"

	"github.com/slackyguy/gorest/base"
	"github.com/slackyguy/gorest/controller/examples"
	"github.com/slackyguy/gorest/routing"
	"google.golang.org/appengine"

	"github.com/slackyguy/gorest/controller/oauth"
	"github.com/slackyguy/gorest/controller/websocket"
)

func main() {

	// OBS: The references of packages with
	// "DoNothing" only ensures that the respective
	// package is loaded (calling the init() mehod).
	// The order is also important:
	// The "Router" is currently initialized
	// by examples.init()
	examples.DoNothing()
	oauth.DoNothing()
	websocket.DoNothing()

	routing.HTTP.Handle()

	// Option 1: Using http.ListenAndServe:
	SetupBasicHTTPSettings()
	log.Fatal(http.ListenAndServe(":8080", nil))

	// Option 2: Using AppEngine
	// SetupAppEngineSettings()
	// appengine.Main()

	// "goapp.properties" example:
	// databaseURL = https://[base].firebaseio.com
	// serviceUID = [service-uid]
	// credentialsFile = [credentials-file].json
}

// SetupBasicHTTPSettings setup http.ListenAndServe method
func SetupBasicHTTPSettings() {

	base.ApplicationSettings = base.GetFromFile("goapp.properties").SetContextFactory(
		func(request *http.Request) context.Context {
			return context.Background()
		})
}

// SetupAppEngineSettings setup AppEngine method
// running: dev_appserver.py app.yaml
func SetupAppEngineSettings() {

	base.ApplicationSettings = base.GetFromFile("goapp.properties").SetContextFactory(
		func(request *http.Request) context.Context {
			return appengine.NewContext(request)
		})
}
