package gateway

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/otakakot/ninshow/internal/domain/model"
	"github.com/otakakot/ninshow/internal/domain/repository"
)

var _ repository.Account = (*Account)(nil)

type Account struct {
	rdb *RDB
}

func NewAcccount(
	rdb *RDB,
) *Account {
	return &Account{
		rdb: rdb,
	}
}

// Save implements repository.Account.
func (ac *Account) Save(
	ctx context.Context,
	account model.Account,
) error {
	query := `INSERT INTO accounts (id, email, username, password) VALUES ($1, $2, $3, $4)`

	stmt, err := ac.rdb.Client.Prepare(query)
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %w", err)
	}

	if _, err := stmt.ExecContext(ctx, account.ID, account.Email, account.Username, account.HashPass); err != nil {
		return fmt.Errorf("failed to execute statement: %w", err)
	}

	return nil
}

// Find implements repository.Account.
func (ac *Account) Find(
	ctx context.Context,
	id string,
) (*model.Account, error) {
	query := `SELECT id, email, username, password FROM accounts WHERE id = $1`

	row := ac.rdb.Client.QueryRowContext(ctx, query, id)

	var account model.Account
	if err := row.Scan(&account.ID, &account.Email, &account.Username, &account.HashPass); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("account not found")
		}

		return nil, fmt.Errorf("failed to scan row: %w", err)
	}

	return &account, nil
}

// FindByUsername implements repository.Account.
func (ac *Account) FindByUsername(
	ctx context.Context,
	username string,
) (*model.Account, error) {
	query := `SELECT id, email, username, password FROM accounts WHERE username = $1`

	row := ac.rdb.Client.QueryRowContext(ctx, query, username)

	var account model.Account
	if err := row.Scan(&account.ID, &account.Email, &account.Username, &account.HashPass); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("account not found")
		}

		return nil, fmt.Errorf("failed to scan row: %w", err)
	}

	return &account, nil
}

var _ repository.OIDCClient = (*OIDCClient)(nil)

type OIDCClient struct {
	rdb *RDB
}

func NewOIDCClient(
	rdb *RDB,
) *OIDCClient {
	return &OIDCClient{
		rdb: rdb,
	}
}

// Find implements repository.OIDCClient.
func (oc *OIDCClient) Find(
	ctx context.Context,
	id string,
) (*model.OIDCClient, error) {
	query := `SELECT id, secret, name, redirect_uri FROM oidc_clients WHERE id = $1`

	row := oc.rdb.Client.QueryRowContext(ctx, query, id)

	var oidcClient model.OIDCClient
	if err := row.Scan(&oidcClient.ID, &oidcClient.HashSec, &oidcClient.Name, &oidcClient.RedirectURI); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("oidc client not found")
		}

		return nil, fmt.Errorf("failed to scan row: %w", err)
	}

	return &oidcClient, nil
}

// Save implements repository.OIDCClient.
func (oc *OIDCClient) Save(
	ctx context.Context,
	client model.OIDCClient,
) error {
	query := `INSERT INTO oidc_clients (id, secret, name, redirect_uri) VALUES ($1, $2, $3, $4)`

	stmt, err := oc.rdb.Client.Prepare(query)
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %w", err)
	}

	if _, err := stmt.ExecContext(ctx, client.ID, client.HashSec, client.Name, client.RedirectURI); err != nil {
		return fmt.Errorf("failed to execute statement: %w", err)
	}

	return nil
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

func NewRefreshTokenCache() *KVS[string] {
	return NewKVS[string]()
}
