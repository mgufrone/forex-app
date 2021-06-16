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
	return newDbCriteria()
}

func (d *dbQuery) Apply(cb criteria.ICriteriaBuilder) rate.IQuery {
	d.mods = cb.(dbCriteria).mods
	return d
}

func (d *dbQuery) GetAll(ctx context.Context) (out []*rate.Rate, err error) {
	rates, err := models.Rates(d.mods...).All(ctx, d.db)
	if err != nil {
		return
	}
	out = make([]*rate.Rate, len(rates))
	for idx, c := range rates {
		rt := rate.NewRate(c.Base, c.Symbol, c.Source, c.SourceType, c.Sell, c.Buy, c.UpdatedAt)
		rt.SetID(c.ID)
		out[idx] = rt
	}
	return
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
