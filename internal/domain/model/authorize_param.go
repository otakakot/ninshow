package model

type AuthorizeParam struct {
	RedirectURI   string
	ClientID      string
	ResponseType  string
	Scope         []string
	State         string
	Nonce         string
	CodeChallenge *CodeChallenge
}
