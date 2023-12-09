package usecase

import "context"

type Security interface {
	HandleBearer(ctx context.Context, input HandleBearerInput) (*HandleBearerOutput, error)
}

type HandleBearerInput struct {
	Sign  string
	Token string
}

type HandleBearerOutput struct {
	Ctx context.Context
}
