package client

import (
	"context"
	"firebase.google.com/go/v4/auth"
	"golang-clean-arch/exceptions"
	"golang-clean-arch/usecases/out"
)

type firebaseClient struct {
	authConfig *auth.Client
}

func NewFirebaseClient(authConfig *auth.Client) out.FirebaseDataSource {
	return &firebaseClient{
		authConfig: authConfig,
	}
}

func (f firebaseClient) GetGoogleSSOLinkWithEmailInput(email string) string {
	action := auth.ActionCodeSettings{
		URL:                   "https://henhaochi.io/auth",
		HandleCodeInApp:       false,
		IOSBundleID:           "",
		AndroidPackageName:    "",
		AndroidMinimumVersion: "",
		AndroidInstallApp:     false,
		DynamicLinkDomain:     "",
	}

	link, err := f.authConfig.EmailSignInLink(context.Background(), email, &action)
	exceptions.PanicIfNeeded(err)
	return link
}
