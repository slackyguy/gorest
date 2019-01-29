package firebase

import (
	"fmt"

	firebasego "firebase.google.com/go"
	"github.com/slackyguy/gorest/lib"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
	//"firebase.google.com/go/auth"
	"firebase.google.com/go/db"
)

// App returns firebase app.
func App(
	ctx context.Context,
	appSettings *lib.AppSettings) (
	app *firebasego.App,
	err error) {

	conf := &firebasego.Config{
		DatabaseURL:  appSettings.DatabaseURL,
		AuthOverride: &map[string]interface{}{"uid": appSettings.ServiceUID},
	}

	opt := option.WithCredentialsFile(appSettings.CredentialsFile)
	app, err = firebasego.NewApp(ctx, conf, opt)

	return
}

// Client returns firebase client.
func Client(
	ctx context.Context,
	app *firebasego.App) (
	client *db.Client,
	err error) {

	client, err = app.Database(ctx)

	if err != nil {
		err = fmt.Errorf("error initializing app: %v", err)
	}

	return
}
