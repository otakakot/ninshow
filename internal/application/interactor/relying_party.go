package interactor

import (
	"bytes"
	"context"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/otakakot/ninshow/internal/application/usecase"
	"github.com/otakakot/ninshow/pkg/log"
)

var _ usecase.RelyingParty = (*RelyingParty)(nil)

type RelyingParty struct{}

func NewRelyingParty() *RelyingParty {
	return &RelyingParty{}
}

// Login implements usecase.RelyingParty.
func (*RelyingParty) Login(
	ctx context.Context,
	input usecase.RelyingPartyLoginInput,
) (*usecase.RelyingPartyLoginOutput, error) {
	end := log.StartEnd(ctx)
	defer end()

	var buf bytes.Buffer

	buf.WriteString(input.OIDCEndpoint)

	state := uuid.NewString()

	cookie := &http.Cookie{
		Name:     "state",
		Value:    state,
		Path:     "/",
		Domain:   "localhost",
		Expires:  time.Now().Add(time.Hour),
		Secure:   true,
		HttpOnly: true,
	}

	values := url.Values{
		"response_type": {"code"},                         // Authorization Flow なので code を指定
		"client_id":     {input.ClientID},                 // RPを識別するためのID OPに登録しておく必要がある
		"redirect_uri":  {input.RedirectURI},              // ログイン後にリダイレクトさせるURL OPに登録しておく必要がある
		"scope":         {strings.Join(input.Scope, " ")}, // RPが要求するスコープ OPに登録しておく必要がある
		"state":         {state},                          // CSRF対策のためのstate
		"nonce":         {uuid.NewString()},               // CSRF対策のためのnonce
	}

	buf.WriteByte('?')

	buf.WriteString(values.Encode())

	redirect, _ := url.ParseRequestURI(buf.String())

	return &usecase.RelyingPartyLoginOutput{
		Cookie:      cookie,
		RedirectURI: *redirect,
	}, nil
}

// Callback implements usecase.RelyingParty.
func (*RelyingParty) Callback(
	ctx context.Context,
	input usecase.RelyingPartyCallbackInput,
) (*usecase.RelyingPartyCallbackOutput, error) {
	end := log.StartEnd(ctx)
	defer end()

	// OpenID Provider へ code を送信して ID Token を取得する

	// ID Token を検証する

	// Accsess Token を利用し OpenID Provider へ UserInfo Request を送信する

	panic("unimplemented")
}
