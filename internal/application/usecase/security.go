package usecase

import "context"

type Security interface {
	HandleBearer(context.Context, HandleBearerInput) (*HandleBearerOutput, error)
}

type HandleBearerInput struct {
	Sign  string
	Token string
}

type HandleBearerOutput struct {
	Ctx context.Context
}
