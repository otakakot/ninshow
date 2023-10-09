package gateway

import (
	"context"
	"fmt"
	"sync"
	"time"

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

var _ repository.Cache[any] = (*KVS[any])(nil)

func NewKVS[T any]() *KVS[T] {
	return &KVS[T]{
		values: make(map[string]T),
	}
}

type KVS[T any] struct {
	mu     sync.Mutex
	values map[string]T
}

func (kvs *KVS[T]) Set(
	_ context.Context,
	key string,
	val T,
	_ time.Duration,
) error {
	kvs.mu.Lock()
	defer kvs.mu.Unlock()

	kvs.values[key] = val

	return nil
}

func (kvs *KVS[T]) Get(
	_ context.Context,
	key string,
) (T, error) {
	var val T

	kvs.mu.Lock()
	defer kvs.mu.Unlock()

	if _, ok := kvs.values[key]; !ok {
		return val, fmt.Errorf("key %s not found", key)
	}

	return kvs.values[key], nil
}

func (kvs *KVS[T]) Del(
	_ context.Context,
	key string,
) error {
	kvs.mu.Lock()
	defer kvs.mu.Unlock()

	delete(kvs.values, key)

	return nil
}

func NewParamCache() *KVS[model.AuthorizeParam] {
	return NewKVS[model.AuthorizeParam]()
}

func NewLoggedInCache() *KVS[model.LoggedIn] {
	return NewKVS[model.LoggedIn]()
}

func NewAccessTokenCache() *KVS[struct{}] {
	return NewKVS[struct{}]()
}

func NewRefreshTokenCache() *KVS[struct{}] {
	return NewKVS[struct{}]()
}
