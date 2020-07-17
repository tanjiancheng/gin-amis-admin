package model

import (
	"context"
)

type IApp interface {
	Init(ctx context.Context, appId string) error
	Query(ctx context.Context, appId string) (bool, error)
}
