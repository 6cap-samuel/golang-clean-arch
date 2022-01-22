package out

type FirebaseDataSource interface {
	GetGoogleSSOLinkWithEmailInput(email string) string
}
