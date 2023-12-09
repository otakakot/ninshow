package usecase

import (
	"context"
	"io"
	"net/http"
	"net/url"
)

type RelyingParty interface {
	Login(ctx context.Context, input RelyingPartyLoginInput) (*RelyingPartyLoginOutput, error)
	Callback(ctx context.Context, input RelyingPartyCallbackInput) (*RelyingPartyCallbackOutput, error)
}

type RelyingPartyLoginInput struct {
	OIDCEndpoint string   // OP の認証エンドポイント
	ClientID     string   // OP に登録してあるクライアントID
	RedirectURI  string   // OP に登録してある認証成功後にリダイレクトさせるURI
	Scope        []string // OP に登録してあるスコープ
}

type RelyingPartyLoginOutput struct {
	Cookie      *http.Cookie
	RedirectURI url.URL
}

type RelyingPartyCallbackInput struct {
	Code         string
	OIDCEndpoint string
	ClientID     string
	ClientSecret string
}

type RelyingPartyCallbackOutput struct {
	Data io.Reader
}
