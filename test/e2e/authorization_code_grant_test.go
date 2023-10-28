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

	t.Run("認可コードグラント", func(t *testing.T) {
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

		clientID := "e4110264-ca70-4179-8958-195542ddc9bd"

		baseUrl, _ := url.Parse(fmt.Sprintf("%s/op/authorize", endpoint))
		params := url.Values{}
		params.Add("response_type", "code")
		params.Add("scope", "openid profile email")
		params.Add("client_id", clientID)
		params.Add("redirect_uri", "http://localhost:8080")
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

	t.Run("許可されていないscopeにより失敗", func(t *testing.T) {
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

		clientID := "e4110264-ca70-4179-8958-195542ddc9bd"

		baseUrl, _ := url.Parse(fmt.Sprintf("%s/op/authorize", endpoint))
		params := url.Values{}
		params.Add("response_type", "code")
		params.Add("scope", "openid phone") // 許可されていないscope phone を含める
		params.Add("client_id", clientID)
		params.Add("redirect_uri", "http://localhost:8080")
		params.Add("state", state)
		params.Add("nonce", nonce)

		baseUrl.RawQuery = params.Encode()

		resp, err := http.Get(baseUrl.String())
		if err != nil {
			t.Fatal(err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusBadRequest {
			t.Errorf("unexpected status code: %d", resp.StatusCode)
		}
	})

	t.Run("登録されていないredirect_uriにより失敗", func(t *testing.T) {
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

		clientID := "e4110264-ca70-4179-8958-195542ddc9bd"

		baseUrl, _ := url.Parse(fmt.Sprintf("%s/op/authorize", endpoint))
		params := url.Values{}
		params.Add("response_type", "code")
		params.Add("scope", "openid profile email")
		params.Add("client_id", clientID)
		params.Add("redirect_uri", "http://localhost:5050") // 登録されていないredirect_uri
		params.Add("state", state)
		params.Add("nonce", nonce)

		baseUrl.RawQuery = params.Encode()

		resp, err := http.Get(baseUrl.String())
		if err != nil {
			t.Fatal(err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusBadRequest {
			t.Errorf("unexpected status code: %d", resp.StatusCode)
		}
	})

	t.Run("許可されていないclient_idにより失敗", func(t *testing.T) {
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

		// 登録していないclient_idを指定
		clientID := "0e0c59ee-7a05-4f23-a902-433c7f29a12e"

		baseUrl, _ := url.Parse(fmt.Sprintf("%s/op/authorize", endpoint))
		params := url.Values{}
		params.Add("response_type", "code")
		params.Add("scope", "openid profile email")
		params.Add("client_id", clientID)
		params.Add("redirect_uri", "http://localhost:8080")
		params.Add("state", state)
		params.Add("nonce", nonce)

		baseUrl.RawQuery = params.Encode()

		resp, err := http.Get(baseUrl.String())
		if err != nil {
			t.Fatal(err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusUnauthorized {
			t.Errorf("unexpected status code: %d", resp.StatusCode)
		}
	})
}
