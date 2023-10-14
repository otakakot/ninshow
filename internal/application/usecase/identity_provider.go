package usecase

import (
	"context"
)

type IdentityProvider interface {
	Signup(context.Context, IdentityProviderSignupInput) (*IdentityProviderSignupOutput, error)
	Signin(context.Context, IdentityProviderSigninInput) (*IdentityProviderSigninOutput, error)
}

type IdentityProviderSignupInput struct {
	Email    string
	Username string
	Password string
}

type IdentityProviderSignupOutput struct{}

type IdentityProviderSigninInput struct {
	Username string
	Password string
}

type IdentityProviderSigninOutput struct{}