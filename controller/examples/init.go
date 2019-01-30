package examples

import (
	"net/http"

	"github.com/slackyguy/gorest/routing"
)

// Redirect (redirect code example)
func Redirect(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
}

func init() {
	routing.Register("retailers", newRetailersController("retailers"))
}

// Load callback (this is a workaround to force google cloud appEngine call init)
func Load() {}
