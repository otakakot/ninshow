package model

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Account struct {
	ID       string
	Username string
	Email    string
	Password string
	HashPass []byte
}

func SingupAccount(
	username string,
	email string,
	password string,
) (*Account, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	return &Account{
		ID:       uuid.NewString(),
		Username: username,
		Email:    email,
		Password: password,
		HashPass: hash,
	}, nil
}

func (ac *Account) ComparePassword(
	password string,
) error {
	return bcrypt.CompareHashAndPassword(ac.HashPass, []byte(password))
}

var Accounts = map[string]Account{}
var mu sync.Mutex

func SaveAccount(
	account Account,
) error {
	mu.Lock()
	defer mu.Unlock()

	if _, ok := Accounts[account.Username]; ok {
		return fmt.Errorf("account already exists")
	}

	Accounts[account.Username] = account

	return nil
}

func FindAccount(
	username string,
) (*Account, error) {
	mu.Lock()
	defer mu.Unlock()

	account, ok := Accounts[username]
	if !ok {
		return nil, fmt.Errorf("account not found")
	}

	return &account, nil
}
