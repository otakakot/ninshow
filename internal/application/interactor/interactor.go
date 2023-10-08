package interactor

import (
	"context"
	"fmt"

	"github.com/otakakot/ninshow/internal/application/usecase"
	"github.com/otakakot/ninshow/internal/domain/model"
	"github.com/otakakot/ninshow/internal/domain/repository"
)

var _ usecase.Account = (*Accont)(nil)

type Accont struct {
	account repository.Account
}

func NewAcccount(
	acccount repository.Account,
) *Accont {
	return &Accont{
		account: acccount,
	}
}

// Signup implements usecase.Account.
func (ac *Accont) Signup(
	ctx context.Context,
	input usecase.AccountSignupInput,
) (*usecase.AccountSigninOutput, error) {
	account, err := model.SingupAccount(
		input.Username,
		input.Email,
		input.Password,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to signup account: %w", err)
	}

	if err := ac.account.Save(ctx, *account); err != nil {
		return nil, fmt.Errorf("failed to save account: %w", err)
	}

	return &usecase.AccountSigninOutput{}, nil
}

// Signin implements usecase.Account.
func (ac *Accont) Signin(
	ctx context.Context,
	input usecase.AccountSigninInput,
) (*usecase.AccountSigninOutput, error) {
	account, err := ac.account.Find(ctx, input.Username)
	if err != nil {
		return nil, fmt.Errorf("failed to find account: %w", err)
	}

	if err != account.ComparePassword(input.Password) {
		return nil, fmt.Errorf("failed to compare password: %w", err)
	}

	return &usecase.AccountSigninOutput{}, nil
}
