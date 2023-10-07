package controller

import (
	"context"
	"log/slog"

	"github.com/otakakot/ninshow/internal/model"
	"github.com/otakakot/ninshow/pkg/api"
)

// IdpSignin implements api.Handler.
func (*Controller) IdpSignin(ctx context.Context, req api.OptIdPSigninRequestSchema) (api.IdpSigninRes, error) {
	slog.Info("start idp signin controller")
	defer slog.Info("end idp signin controller")

	account, err := model.FindAccount(req.Value.Username)
	if err != nil {
		return &api.IdpSigninUnauthorized{}, nil
	}

	if err := account.ComparePassword(req.Value.Password); err != nil {
		return &api.IdpSigninUnauthorized{}, nil
	}

	return &api.IdpSigninOK{}, nil
}

// IdpSignup implements api.Handler.
func (*Controller) IdpSignup(ctx context.Context, req api.OptIdPSignupRequestSchema) (api.IdpSignupRes, error) {
	slog.Info("start idp signup controller")
	defer slog.Info("end idp signup controller")

	account, err := model.SingupAccount(req.Value.Username, req.Value.Email, req.Value.Password)
	if err != nil {
		return &api.IdpSignupInternalServerError{}, err
	}

	if err := model.SaveAccount(*account); err != nil {
		return &api.IdpSignupInternalServerError{}, err
	}

	return &api.IdpSignupOK{}, nil
}
