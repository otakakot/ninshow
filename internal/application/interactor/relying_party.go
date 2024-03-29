package interactor

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"log/slog"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/otakakot/ninshow/internal/application/usecase"
	"github.com/otakakot/ninshow/internal/domain/model"
	"github.com/otakakot/ninshow/pkg/api"
	"github.com/otakakot/ninshow/pkg/log"
)

var _ usecase.RelyingParty = (*RelyingParty)(nil)

type RelyingParty struct {
	verifieis map[string]string
}

func NewRelyingParty() *RelyingParty {
	return &RelyingParty{
		verifieis: make(map[string]string),
	}
}

// Login implements usecase.RelyingParty.
func (rp *RelyingParty) Login(
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

	velifier := model.GenerateCodeVerifier()

	challenge := model.GenerateCodeChallenge(velifier)

	rp.verifieis[state] = velifier

	values := url.Values{
		"response_type":         {"code"},                         // Authorization Flow なので code を指定
		"client_id":             {input.ClientID},                 // RPを識別するためのID OPに登録しておく必要がある
		"redirect_uri":          {input.RedirectURI},              // ログイン後にリダイレクトさせるURL OPに登録しておく必要がある
		"scope":                 {strings.Join(input.Scope, " ")}, // RPが要求するスコープ OPに登録しておく必要がある
		"state":                 {state},                          // CSRF対策のためのstate
		"nonce":                 {uuid.NewString()},               // CSRF対策のためのnonce
		"code_challenge":        {challenge.Challenge},            // PKCEのためのcode_challenge
		"code_challenge_method": {string(challenge.Method)},
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
func (rp *RelyingParty) Callback(
	ctx context.Context,
	input usecase.RelyingPartyCallbackInput,
) (*usecase.RelyingPartyCallbackOutput, error) {
	end := log.StartEnd(ctx)
	defer end()

	// OpenID Provider へ code を送信して ID Token を取得する
	cli, err := api.NewClient(input.OIDCEndpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %w", err)
	}

	redirectURI, _ := url.Parse("")

	verifier := rp.verifieis[input.State]

	res, err := cli.OpToken(ctx, &api.OPTokenRequestSchema{
		GrantType:    api.OPTokenRequestSchemaGrantTypeAuthorizationCode,
		Code:         input.Code,
		RedirectURI:  *redirectURI,
		ClientID:     api.NewOptString(input.ClientID),
		ClientSecret: api.NewOptString(input.ClientSecret),
		Scope:        api.NewOptString("openid profile email"),
		CodeVerifier: api.NewOptString(verifier),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to request token: %w", err)
	}

	v, ok := res.(*api.OPTokenResponseSchemaHeaders)
	if !ok {
		return nil, fmt.Errorf("failed to assert response: %T", v)
	}

	// TODO: ID Token を検証する

	if res, err := cli.OpToken(ctx, &api.OPTokenRequestSchema{
		GrantType:    api.OPTokenRequestSchemaGrantTypeRefreshToken,
		RefreshToken: api.NewOptString(v.Response.RefreshToken),
		ClientID:     api.NewOptString(input.ClientID),
		Scope:        api.NewOptString("openid profile email"),
	}); err != nil {
		return nil, fmt.Errorf("failed to request token: %w", err)
	} else {
		switch res.(type) {
		case *api.OPTokenResponseSchemaHeaders:
			slog.Info(fmt.Sprintf("%#v", res))
		default:
			slog.Warn(fmt.Sprintf("%T", res))
		}
	}

	// Accsess Token を利用し OpenID Provider へ UserInfo Request を送信する
	cl, err := api.NewClient(input.OIDCEndpoint, &security{token: v.Response.AccessToken})
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %w", err)
	}

	rs, err := cl.OpUserinfo(ctx)
	if err != nil {
		slog.Warn(fmt.Sprintf("failed to request userinfo: %v", err))

		return nil, fmt.Errorf("failed to request userinfo: %w", err)
	}

	vv, ok := rs.(*api.OPUserInfoResponseSchema)
	if !ok {
		return nil, fmt.Errorf("failed to assert response: %T", vv)
	}

	const tmp = `<!DOCTYPE html>
	<html>
		<head>
			<meta charset="UTF-8">
			<title>Callback</title>
		</head>
		<style>
			.token {
				word-break: break-all;
			}
		</style>
		<body>
			<h1>Relying Party Callback</h1>
			<h2>token type</h2>
			<p>{{.TokenType}}</p>
			<h2>access token</h2>
			<p class="token">{{.AccessToken}}</p>
			<h2>refresh token</h2>
			<p class="token">{{.RefreshToken}}</p>
			<h2>id token</h2>
			<p class="token">{{.IDToken}}</p>
			<h2>sub</h2>
			<p>{{.Sub}}</p>
		</body>
	</html>
	`

	tmpl, _ := template.New("callback").Parse(tmp)

	data := &struct {
		TokenType    string
		AccessToken  string
		RefreshToken string
		IDToken      string
		Sub          string
	}{
		TokenType:    v.Response.TokenType,
		AccessToken:  v.Response.AccessToken,
		RefreshToken: v.Response.RefreshToken,
		IDToken:      v.Response.IDToken,
		Sub:          vv.Sub,
	}

	buf := new(bytes.Buffer)

	if err := tmpl.Execute(buf, data); err != nil {
		return nil, fmt.Errorf("failed to execute template: %w", err)
	}

	return &usecase.RelyingPartyCallbackOutput{
		Data: buf,
	}, nil
}

var _ api.SecuritySource = (*security)(nil)

type security struct {
	token string
}

// Bearer implements api.SecuritySource.
func (sec *security) Bearer(_ context.Context, _ string) (api.Bearer, error) {
	return api.Bearer{Token: sec.token}, nil
}
