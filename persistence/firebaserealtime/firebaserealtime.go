package firebaserealtime

import (
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
	"github.com/slackyguy/gorest/base"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
	//"firebase.google.com/go/auth"
)

// App returns firebase app.
func App(
	ctx context.Context,
	appSettings *base.AppSettings) *firebase.App {

	conf, opt := firebaseSettings(appSettings)
	app, err := firebase.NewApp(ctx, conf, opt)

	check(err)

	return app
}

// Client returns firebase client.
func Client(
	ctx context.Context,
	appSettings *base.AppSettings) *db.Client {

	conf, opt := firebaseSettings(appSettings)
	app, errApp := firebase.NewApp(ctx, conf, opt)

	check(errApp)

	client := client(ctx, app)
	return client
}

// Client returns firebase client.
func client(
	ctx context.Context,
	app *firebase.App) (
	client *db.Client) {

	client, err := app.Database(ctx)

	check(err)

	return
}

func firebaseSettings(appSettings *base.AppSettings) (
	conf *firebase.Config, opt option.ClientOption) {

	conf = &firebase.Config{
		DatabaseURL:  appSettings.DatabaseURL,
		AuthOverride: &map[string]interface{}{"uid": appSettings.ServiceUID},
	}

	opt = option.WithCredentialsFile(appSettings.CredentialsFile)

	return
}

func check(error ...error) {
	for _, err := range error {
		if err != nil {
			log.Fatalf("%v", error)
			return
		}
	}
}
