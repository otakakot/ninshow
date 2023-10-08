package gateway

import (
	"context"
	"fmt"
	"sync"

	"github.com/otakakot/ninshow/internal/domain/model"
	"github.com/otakakot/ninshow/internal/domain/repository"
)

var _ repository.Account = (*Account)(nil)

type Account struct {
	mu       sync.Mutex
	accounts map[string]model.Account
}

func NewAcccount() *Account {
	return &Account{
		accounts: make(map[string]model.Account),
	}
}

// Save implements repository.Account.
func (ac *Account) Save(
	_ context.Context,
	account model.Account,
) error {
	ac.mu.Lock()
	defer ac.mu.Unlock()

	if _, ok := ac.accounts[account.Username]; ok {
		return fmt.Errorf("account already exists")
	}

	ac.accounts[account.Username] = account

	return nil
}

// Find implements repository.Account.
func (ac *Account) Find(
	_ context.Context,
	username string,
) (*model.Account, error) {
	ac.mu.Lock()
	defer ac.mu.Unlock()

	account, ok := ac.accounts[username]
	if !ok {
		return nil, fmt.Errorf("account not found")
	}

	return &account, nil
}
