package rate

import (
	"context"
	"github.com/mgufrone/forex/internal/shared/criteria"
	"time"
)

type IQuery interface {
	CriteriaBuilder() criteria.ICriteriaBuilder
	Apply(cb criteria.ICriteriaBuilder) IQuery
	// expose-able
	Latest(ctx context.Context, date time.Time) (out []*Rate, err error)
	History(ctx context.Context, span TimeSpan, start, end time.Time) (out []*Rate, err error)
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
