package main

import (
	"context"
	"net/http"

	"github.com/slackyguy/gorest/base"
	"github.com/slackyguy/gorest/routing"

	"github.com/slackyguy/gorest/controller/examples"
	"github.com/slackyguy/gorest/controller/oauth"
	"github.com/slackyguy/gorest/controller/websocket"
	"google.golang.org/appengine"
)

func main() {

	routing.HTTP.RegisterHandler("/", func(response http.ResponseWriter, request *http.Request) {
		if request.Method != "GET" {
			http.Error(response, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		http.ServeFile(response, request, "home.html")
	})

	// OBS.:
	// 1 - Assuming that controllers.init() is called first, the initialization
	// of the router was delegated to that method.
	// 2 - If running outside appengine, replace the Load() calls bellow for
	// something with the corresponding init().
	examples.Load()
	oauth.Load()
	websocket.Load()
	routing.HTTP.Handle()

	// running: dev_appserver.py app.yaml
	base.ApplicationSettings = base.GetFromFile("goapp.properties").SetContextFactory(
		func(request *http.Request) context.Context {
			return appengine.NewContext(request)
		})
	appengine.Main()

	// base.ApplicationSettings = base.GetFromFile("goapp.properties").SetContextFactory(
	// 	func(request *http.Request) context.Context {
	// 		return context.Background()
	// 	})
	//log.Fatal(http.ListenAndServe(":8080", nil))
}
