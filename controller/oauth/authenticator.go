package oauth

//go get -u -v gopkg.in/oauth2.v3/...

import (
	"log"
	"net/http"

	"github.com/slackyguy/gorest/routing"
	"gopkg.in/oauth2.v3/errors"
	"gopkg.in/oauth2.v3/manage"
	"gopkg.in/oauth2.v3/models"
	"gopkg.in/oauth2.v3/server"
	"gopkg.in/oauth2.v3/store"
)

// Load is workaround to appengine call init
func Load() {}

func init() {
	manager := manage.NewDefaultManager()
	// token memory store
	manager.MustTokenStorage(store.NewMemoryTokenStore())

	// client memory store
	clientStore := store.NewClientStore()
	clientStore.Set("000000", &models.Client{
		ID:     "000000",
		Secret: "999999",
		Domain: "http://localhost",
	})

	manager.MapClientStorage(clientStore)

	srv := server.NewDefaultServer(manager)
	srv.SetAllowGetAccessRequest(true)
	srv.SetClientInfoHandler(server.ClientFormHandler)

	srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println("Internal Error:", err.Error())
		return
	})

	srv.SetResponseErrorHandler(func(re *errors.Response) {
		log.Println("Response Error:", re.Error.Error())
	})

	routing.RegisterHandler("/authorize", func(
		response http.ResponseWriter, request *http.Request) {
		err := srv.HandleAuthorizeRequest(response, request)
		//err = srv.HandleAuthorizeRequest(w, r)
		if err != nil {
			http.Error(response, err.Error(), http.StatusBadRequest)
		}
	})

	routing.RegisterHandler("/token", func(
		response http.ResponseWriter, request *http.Request) {
		srv.HandleTokenRequest(response, request)
	})
}
