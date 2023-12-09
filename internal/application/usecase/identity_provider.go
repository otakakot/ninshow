package usecase

import (
	"context"
)

type IdentityProvider interface {
	Signup(ctx context.Context, input IdentityProviderSignupInput) (*IdentityProviderSignupOutput, error)
	Signin(ctx context.Context, input IdentityProviderSigninInput) (*IdentityProviderSigninOutput, error)
}

type IdentityProviderSignupInput struct {
	Email    string
	Name     string
	Password string
}

type IdentityProviderSignupOutput struct{}

type IdentityProviderSigninInput struct {
	Email    string
	Password string
}

type IdentityProviderSigninOutput struct{}
