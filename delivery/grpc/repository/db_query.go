package repository

import (
	"context"
	"github.com/mgufrone/forex/delivery/grpc/models"
	"github.com/mgufrone/forex/internal/domains/rate"
	"github.com/mgufrone/forex/internal/shared/criteria"
	"gorm.io/gorm"
)

type dbQuery struct {
	db *gorm.DB
	cb criteria.ICriteriaBuilder
}

func NewQuery(db *gorm.DB) rate.IQuery {
	return &dbQuery{db: db}
}

func (d *dbQuery) CriteriaBuilder() criteria.ICriteriaBuilder {
	return dbCriteriaBuilder{}
}

func (d *dbQuery) Apply(cb criteria.ICriteriaBuilder) rate.IQuery {
	d.cb = cb
	return d
}

func (d *dbQuery) apply(ctx context.Context) *gorm.DB {
	db := d.db
	if d.cb != nil {
		cr := d.cb.(dbCriteriaBuilder)
		db = cr.apply(d.db)
	}
	db = db.WithContext(ctx).Model(&models.Rate{})
	d.cb = nil
	return db
}

func (d *dbQuery) GetAll(ctx context.Context) (out []*rate.Rate, err error) {
	res, total, err := d.GetAndCount(ctx)
	if err != nil || total == 0 {
		return
	}
	for _, r := range res {
		out = append(out, r)
	}
	return
}

func (d *dbQuery) Count(ctx context.Context) (total int64, err error) {
	db := d.apply(ctx)
	return d.count(ctx, db)
}
func (d *dbQuery) count(ctx context.Context, db *gorm.DB) (total int64, err error) {
	err = db.WithContext(ctx).Count(&total).Error
	return
}

func (d *dbQuery) GetAndCount(ctx context.Context) (out []*rate.Rate, total int64, err error) {
	db := d.apply(ctx)
	total, err = d.count(context.Background(), db)
	if err != nil || total == 0 {
		return
	}
	res := make([]*models.Rate, 0)
	err = db.Find(&res).Error
	if err == nil {
		for _, m := range res {
			mdl, _ := m.ToDomain()
			out = append(out, mdl)
		}
	}
	return
}

func (d *dbQuery) FindByID(ctx context.Context, id string) (out *rate.Rate, err error) {
	db := d.db
	var res *models.Rate
	err = db.WithContext(ctx).Where("id = ?", id).First(&res).Error
	if err == nil {
		out, err = res.ToDomain()
	}
	return
}
