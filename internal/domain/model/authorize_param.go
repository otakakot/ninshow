package model

type AuthorizeParam struct {
	RedirectURI string
	State       string
	Nonce       string
}
