package usecase

import "context"

type Account interface {
	Signup(context.Context, AccountSignupInput) (*AccountSigninOutput, error)
	Signin(context.Context, AccountSigninInput) (*AccountSigninOutput, error)
}

type AccountSignupInput struct {
	Email    string
	Username string
	Password string
}

type AccountSignupOutput struct{}

type AccountSigninInput struct {
	Username string
	Password string
}

type AccountSigninOutput struct{}
