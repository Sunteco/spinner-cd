package mongo

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/mon"
)

type baseRepository[T any] interface {
	Insert(ctx context.Context, data *T) error
	FindOne(ctx context.Context, id string) (*T, error)
	Update(ctx context.Context, data *T) error
	Delete(ctx context.Context, id string) error
}

type repository struct {
	conn *mon.Model
}

func newRepository(conn *mon.Model) *repository {
	return &repository{conn: conn}
}
