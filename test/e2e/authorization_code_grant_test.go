package e2e_test

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/otakakot/ninshow/pkg/api"
)

func TestAuthorizationCodeGrant(t *testing.T) {
	t.Parallel()

	endpoint := os.Getenv("ENDPOINT")
	if endpoint == "" {
		endpoint = "http://localhost:8080"
	}

	cli, err := api.NewClient(endpoint, nil)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("authorization_code_grant", func(t *testing.T) {
		t.Parallel()

		username := uuid.NewString()

		email := fmt.Sprintf("%s@example.com", uuid.NewString())

		password := uuid.NewString()

		req := api.IdPSignupRequestSchema{
			Username: username,
			Email:    email,
			Password: password,
		}

		if res, err := cli.IdpSignup(context.Background(), api.NewOptIdPSignupRequestSchema(
			req,
		)); err != nil {
			t.Fatal(err)
		} else {
			switch res.(type) {
			case *api.IdpSignupOK:
			default:
				t.Fatalf("unexpected type: %T", res)
			}
		}

		state := uuid.NewString()

		nonce := uuid.NewString()

		// NOTE: http.Get は 302 が返ってくるとリダイレクトしてくれるらしい
		// NOTE: ogen は上記想定が考慮できていないためうまくレスポンスをデコードできない
		// redirectURI, _ := url.ParseRequestURI("http://localhost:8080")
		// res, err := cli.OpAuthorize(context.Background(), api.OpAuthorizeParams{
		// 	ResponseType: "code",
		// 	Scope:        "openid",
		// 	ClientID:     *redirectURI,
		// 	RedirectURI:  *redirectURI,
		// 	State:        api.NewOptString(state),
		// 	Nonce:        api.NewOptString(nonce),
		// })
		// if err != nil {
		// 	t.Fatal(err)
		// }

		baseUrl, _ := url.Parse(fmt.Sprintf("%s/op/authorize", endpoint))
		params := url.Values{}
		params.Add("response_type", "code")
		params.Add("scope", "openid profile email")
		params.Add("client_id", "26bf8924-c1d9-484d-8a72-db1df2b05ccd")
		params.Add("redirect_uri", "http://localhost:8080/rp/callback")
		params.Add("state", state)
		params.Add("nonce", nonce)

		baseUrl.RawQuery = params.Encode()

		resp, err := http.Get(baseUrl.String())
		if err != nil {
			t.Fatal(err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			t.Fatalf("unexpected status code: %d", resp.StatusCode)
		}

		fmt.Println("Response body:", resp.Body)
	})
}
