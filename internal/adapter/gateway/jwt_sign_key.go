package gateway

import (
	"context"
	"crypto/rand"
	"crypto/rsa"

	"github.com/google/uuid"
	"github.com/otakakot/ninshow/internal/domain/model"
	"github.com/otakakot/ninshow/internal/domain/repository"
)

var _ repository.JWTSignKey = (*JWTSignKey)(nil)

type JWTSignKey struct {
	rdb *RDB
	id  string
	key *rsa.PrivateKey
}

func NewJWTSignKey(rdb *RDB) *JWTSignKey {
	id := uuid.New().String()

	reader := rand.Reader

	bitSize := 2048

	key, _ := rsa.GenerateKey(reader, bitSize)

	return &JWTSignKey{
		rdb: rdb,
		id:  id,
		key: key,
	}
}

// Find implements repository.JWTSignKey.
func (*JWTSignKey) Find(
	ctx context.Context,
	id string,
) (*model.JWTSignKey, error) {
	panic("unimplemented")
}

// List implements repository.JWTSignKey.
func (gw *JWTSignKey) List(
	ctx context.Context,
) ([]model.JWTSignKey, error) {

	return []model.JWTSignKey{
		{
			ID:  gw.id,
			Key: gw.key,
		},
	}, nil
}

// Save implements repository.JWTSignKey.
func (*JWTSignKey) Save(
	ctx context.Context,
	key model.JWTSignKey,
) error {
	panic("unimplemented")
}
