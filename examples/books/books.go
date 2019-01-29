package books

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/slackyguy/gorest/lib"
	"github.com/slackyguy/gorest/lib/appengine"
)

// Book is a json-serializable type.
type Book struct {
	ISBN   string `json:"isbn,omitempty"`
	Author string `json:"author,omitempty"`
	Title  string `json:"title,omitempty"`
}

// (no public constructor)
type controller struct {
	lib.Controller
}

// New returns a new stores controller
func New(context context.Context, settings *lib.AppSettings) lib.ControllerInterface {
	controller := new(controller)
	controller.Context = context
	controller.AppSettings = settings
	return controller
}

// Get API handler (stores)
func (controller) Get(w http.ResponseWriter, r *http.Request) {

}

// List API handler (stores)
func (controller) List(w http.ResponseWriter, r *http.Request) {
	// return func(w http.ResponseWriter, r *http.Request) {
	//
	// 	app, _ := firebase.FirebaseApp(ctx, databaseURL, serviceUID)
	// 	client, _ := firebase.FirebaseClient(ctx, app)
	// 	data, error := firebase.ReadData(ctx, client, "users")
	//
	// 	if error != nil {
	// 		fmt.Fprintln(w, error)
	// 	} else {
	// 		fmt.Fprintln(w, data)
	// 	}
	// }
}

// Post API handler (stores)
func (controller) Post(w http.ResponseWriter, r *http.Request) {
}

// Delete API handler (stores)
func (controller) Delete(w http.ResponseWriter, r *http.Request) {

}

// TestingInit (just a test over )
func TestingInit() {}
func init() {
	router := mux.NewRouter()
	appSettings := lib.ReadFromFile("goapp.properties")
	appengine.Register("books", appSettings, router, New)
}
