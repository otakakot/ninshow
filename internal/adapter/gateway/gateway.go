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
	tx, err := ac.rdb.Client.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	query := `INSERT INTO accounts (id, email, name) VALUES ($1, $2, $3)`

	stmt, err := tx.Prepare(query)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return fmt.Errorf("failed to rollback transaction: %w", err)
		}

		return fmt.Errorf("failed to prepare statement: %w", err)
	}

	if _, err := stmt.ExecContext(ctx, account.ID, account.Email, account.Name); err != nil {
		if err := tx.Rollback(); err != nil {
			return fmt.Errorf("failed to rollback transaction: %w", err)
		}

		return fmt.Errorf("failed to execute statement: %w", err)
	}

	query = `INSERT INTO password_authns (value, account_id) VALUES ($1, $2)`

	stmt, err = tx.Prepare(query)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return fmt.Errorf("failed to rollback transaction: %w", err)
		}

		return fmt.Errorf("failed to prepare statement: %w", err)
	}

	if _, err := stmt.ExecContext(ctx, account.HashPass, account.ID); err != nil {
		if err := tx.Rollback(); err != nil {
			return fmt.Errorf("failed to rollback transaction: %w", err)
		}

		return fmt.Errorf("failed to execute statement: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

// Find implements repository.Account.
func (ac *Account) Find(
	ctx context.Context,
	id string,
) (*model.Account, error) {
	query := `SELECT id, email, name FROM accounts WHERE id = $1`

	row := ac.rdb.Client.QueryRowContext(ctx, query, id)

	var account model.Account
	if err := row.Scan(&account.ID, &account.Email, &account.Name); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("account not found")
		}

		return nil, fmt.Errorf("failed to scan row: %w", err)
	}

	return &account, nil
}

// FindByEmail implements repository.Account.
func (ac *Account) FindByEmail(
	ctx context.Context,
	email string,
) (*model.Account, error) {
	query := `SELECT id, email, name FROM accounts WHERE email = $1`

	row := ac.rdb.Client.QueryRowContext(ctx, query, email)

	var account model.Account
	if err := row.Scan(&account.ID, &account.Email, &account.Name); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("account not found")
		}

		return nil, fmt.Errorf("failed to scan row: %w", err)
	}

	return &account, nil
}

// FindPassword implements repository.Account.
func (ac *Account) FindPassword(
	ctx context.Context,
	accountID string,
) ([]byte, error) {
	query := `SELECT value FROM password_authns WHERE account_id = $1`

	row := ac.rdb.Client.QueryRowContext(ctx, query, accountID)

	var hashPass []byte
	if err := row.Scan(&hashPass); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("account not found")
		}

		return nil, fmt.Errorf("failed to scan row: %w", err)
	}

	return hashPass, nil
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
	query := `SELECT id, name, redirect_uri FROM oidc_clients WHERE id = $1`

	row := oc.rdb.Client.QueryRowContext(ctx, query, id)

	var oidcClient model.OIDCClient
	if err := row.Scan(&oidcClient.ID, &oidcClient.Name, &oidcClient.RedirectURI); err != nil {
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
	tx, err := oc.rdb.Client.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	query := `INSERT INTO oidc_clients (id, name, redirect_uri) VALUES ($1, $2, $3)`

	stmt, err := tx.Prepare(query)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return fmt.Errorf("failed to rollback transaction: %w", err)
		}

		return fmt.Errorf("failed to prepare statement: %w", err)
	}

	if _, err := stmt.ExecContext(ctx, client.ID, client.Name, client.RedirectURI); err != nil {
		if err := tx.Rollback(); err != nil {
			return fmt.Errorf("failed to rollback transaction: %w", err)
		}

		return fmt.Errorf("failed to execute statement: %w", err)
	}

	query = `INSERT INTO oidc_secrets (value, client_id) VALUES ($1, $2)`

	stmt, err = tx.Prepare(query)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return fmt.Errorf("failed to rollback transaction: %w", err)
		}

		return fmt.Errorf("failed to prepare statement: %w", err)
	}

	if _, err := stmt.ExecContext(ctx, client.HashSec, client.ID); err != nil {
		if err := tx.Rollback(); err != nil {
			return fmt.Errorf("failed to rollback transaction: %w", err)
		}

		return fmt.Errorf("failed to execute statement: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

// FindSecret implements repository.OIDCClient.
func (oc *OIDCClient) FindSecret(
	ctx context.Context,
	clientID string,
) ([]byte, error) {
	query := `SELECT value FROM oidc_secrets WHERE client_id = $1`

	row := oc.rdb.Client.QueryRowContext(ctx, query, clientID)

	var hashSec []byte
	if err := row.Scan(&hashSec); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("oidc client not found")
		}

		return nil, fmt.Errorf("failed to scan row: %w", err)
	}

	return hashSec, nil
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
