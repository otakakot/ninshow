package interactor

import (
	"context"
	"fmt"

	"github.com/otakakot/ninshow/internal/application/usecase"
	"github.com/otakakot/ninshow/internal/domain/model"
	"github.com/otakakot/ninshow/internal/domain/repository"
	"github.com/otakakot/ninshow/pkg/log"
)

var _ usecase.IdentityProvider = (*IdentityProvider)(nil)

type IdentityProvider struct {
	account repository.Account
}

func NewIdentityProvider(
	acccount repository.Account,
) *IdentityProvider {
	return &IdentityProvider{
		account: acccount,
	}
}

// Signup implements usecase.IdentityProvider.
func (idp *IdentityProvider) Signup(
	ctx context.Context,
	input usecase.IdentityProviderSignupInput,
) (*usecase.IdentityProviderSignupOutput, error) {
	end := log.StartEnd(ctx)
	defer end()

	account, err := model.SingupAccount(
		input.Username,
		input.Email,
		input.Password,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to signup account: %w", err)
	}

	if err := idp.account.Save(ctx, *account); err != nil {
		return nil, fmt.Errorf("failed to save account: %w", err)
	}

	return &usecase.IdentityProviderSignupOutput{}, nil
}

// Signin implements usecase.IdentityProvider.
func (idp *IdentityProvider) Signin(
	ctx context.Context,
	input usecase.IdentityProviderSigninInput,
) (*usecase.IdentityProviderSigninOutput, error) {
	end := log.StartEnd(ctx)
	defer end()

	account, err := idp.account.Find(ctx, input.Username)
	if err != nil {
		return nil, fmt.Errorf("failed to find account: %w", err)
	}

	if err != account.ComparePassword(input.Password) {
		return nil, fmt.Errorf("failed to compare password: %w", err)
	}

	return &usecase.IdentityProviderSigninOutput{}, nil
}
