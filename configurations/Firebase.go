package configurations

import (
	"context"
	"encoding/base64"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"golang-clean-arch/exceptions"
	"google.golang.org/api/option"
	"os"
)

func NewFirebaseApp() *firebase.App {

	d, _ := base64.StdEncoding.DecodeString(os.Getenv("GCP_SERVICE_ACCOUNT"))

	opt := option.WithCredentialsJSON(d)
	config := &firebase.Config{ProjectID: "myfoodblog-c6e9c"}
	app, err := firebase.NewApp(context.Background(), config, opt)
	exceptions.PanicIfNeeded(err)

	return app
}

func NewFirebaseAuth() *auth.Client {
	authClient, err := NewFirebaseApp().Auth(context.Background())
	exceptions.PanicIfNeeded(err)

	return authClient
}
