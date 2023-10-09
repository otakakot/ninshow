package model

type AuthorizeParam struct {
	RedirectURI string
	ClientID    string
	State       string
	Nonce       string
}
