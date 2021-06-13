package rate

import (
	"context"
	"github.com/mgufrone/forex/internal/shared/criteria"
)

type IQuery interface {
	CriteriaBuilder() criteria.ICriteriaBuilder
	Apply(cb criteria.ICriteriaBuilder) IQuery
	// expose-able
	GetAll(ctx context.Context) (out []*Rate, err error)
	Count(ctx context.Context) (total int64, err error)
	GetAndCount(ctx context.Context) (out []*Rate, total int64, err error)
	FindByID(ctx context.Context, id string) (out *Rate, err error)
}

type ICommand interface {
	Create(ctx context.Context, in *Rate) (err error)
	Update(ctx context.Context, in *Rate) (err error)
	Delete(ctx context.Context, in *Rate) (err error)
}
