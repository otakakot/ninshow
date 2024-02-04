package e2e_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/otakakot/ninshow/pkg/api"
)

func TestIdP(t *testing.T) {
	t.Parallel()

	endpoint := os.Getenv("ENDPOINT")
	if endpoint == "" {
		endpoint = "http://localhost:8080"
	}

	cli, err := api.NewClient(endpoint, nil)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("signup_signin", func(t *testing.T) {
		t.Parallel()

		req := api.IdPSignupRequestSchema{
			Name:     uuid.NewString(),
			Email:    fmt.Sprintf("%s@example.com", uuid.NewString()),
			Password: uuid.NewString(),
		}

		signupRes, err := cli.IdpSignup(context.Background(), api.NewOptIdPSignupRequestSchema(
			req,
		))
		if err != nil {
			t.Error(err)
		}

		switch signupRes.(type) {
		case *api.IdpSignupOK:
		default:
			t.Errorf("unexpected type: %T", signupRes)
		}

		singinRes, err := cli.IdpSignin(context.Background(), api.NewOptIdPSigninRequestSchema(api.IdPSigninRequestSchema{
			Email:    req.Email,
			Password: req.Password,
		}))
		if err != nil {
			t.Error(err)
		}

		switch singinRes.(type) {
		case *api.IdpSigninOK:
			return
		default:
			t.Errorf("unexpected type: %T", singinRes)
		}
	})

	t.Run("unauthorized", func(t *testing.T) {
		t.Parallel()

		req := api.IdPSignupRequestSchema{
			Name:     uuid.NewString(),
			Email:    fmt.Sprintf("%s@example.com", uuid.NewString()),
			Password: uuid.NewString(),
		}

		signupRes, err := cli.IdpSignup(context.Background(), api.NewOptIdPSignupRequestSchema(
			req,
		))
		if err != nil {
			t.Error(err)
		}

		switch signupRes.(type) {
		case *api.IdpSignupOK:
		default:
			t.Errorf("unexpected type: %T", signupRes)
		}

		singinRes, err := cli.IdpSignin(context.Background(), api.NewOptIdPSigninRequestSchema(api.IdPSigninRequestSchema{
			Email:    req.Email,
			Password: "password",
		}))
		if err != nil {
			t.Error(err)
		}

		switch singinRes.(type) {
		case *api.IdpSigninUnauthorized:
			return
		default:
			t.Errorf("unexpected type: %T", singinRes)
		}
	})
}
