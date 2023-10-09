package model

type AuthorizeParam struct {
	RedirectURI string
	ClientID    string
	Scope       []string
	State       string
	Nonce       string
}
