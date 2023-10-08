package model

import (
	"fmt"

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