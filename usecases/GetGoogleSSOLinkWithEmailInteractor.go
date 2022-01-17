package usecases

import (
	"golang-clean-arch/usecases/in"
	"golang-clean-arch/usecases/out"
)

type getGoogleSsoLinkWithEmailInteractor struct {
	firebaseDataSource out.FirebaseDataSource
}

func NewGetGoogleSsoLinkWithEmailInteractor(
	firebaseDataSource *out.FirebaseDataSource,
) in.GetGoogleSSOLinkWithEmailinput {
	return &getGoogleSsoLinkWithEmailInteractor{
		*firebaseDataSource,
	}
}

func (g getGoogleSsoLinkWithEmailInteractor) Get(email string) string {
	return g.firebaseDataSource.GetGoogleSSOLinkWithEmailInput(email)
}
