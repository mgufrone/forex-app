package repository

import (
	"context"
	"database/sql"
	"github.com/mgufrone/forex/delivery/grpc/models"
	"github.com/mgufrone/forex/internal/domains/rate"
	"github.com/mgufrone/forex/internal/shared/criteria"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type dbQuery struct {
	db *sql.DB
	mdl *models.Rate
	mods []qm.QueryMod
}

func (d *dbQuery) CriteriaBuilder() criteria.ICriteriaBuilder {
	panic("implement me")
}

func (d *dbQuery) Apply(cb criteria.ICriteriaBuilder) rate.IQuery {
	panic("implement me")
}

func (d *dbQuery) GetAll(ctx context.Context) (out []*rate.Rate, err error) {
	models.Rates().All()
}

func (d *dbQuery) Count(ctx context.Context) (total int64, err error) {
	panic("implement me")
}

func (d *dbQuery) GetAndCount(ctx context.Context) (out []*rate.Rate, total int64, err error) {
	panic("implement me")
}

func (d *dbQuery) FindByID(ctx context.Context, id string) (out *rate.Rate, err error) {
	panic("implement me")
}
