package repository

import (
	"context"
	"fmt"
	models2 "github.com/mgufrone/forex/delivery/grpc/rate/models"
	"github.com/mgufrone/forex/internal/domains/rate"
	"github.com/mgufrone/forex/internal/shared/criteria"
	"gorm.io/gorm"
	"math"
	"time"
)

type dbQuery struct {
	db *gorm.DB
	cb criteria.ICriteriaBuilder
}

func (d *dbQuery) Latest(ctx context.Context, date time.Time) (out []*rate.Rate, err error) {
	ori := d.db
	db := d.apply(ctx)
	if !date.IsZero() {
		db = db.Where(fmt.Sprintf("DATE(%s) = ?", rate.UpdatedAtColumn), date.Format(rate.DateLayout))
	}
	subQuery := db.Session(&gorm.Session{}).Order(fmt.Sprintf("%s %s", rate.UpdatedAtColumn, "DESC"))
	var res []*models2.Rate
	err = ori.Table("(?) as u", subQuery).
		Group(rate.SourceColumn).
		Group(rate.SourceTypeColumn).
		Group(rate.SymbolColumn).
		Group(rate.BaseColumn).
		Order(fmt.Sprintf("%s %s", rate.BaseColumn, "ASC")).
		Order(fmt.Sprintf("%s %s", rate.SymbolColumn, "ASC")).
		Find(&res).Error
	if err != nil {
		return
	}
	for _, data := range res {
		mdl, err := data.ToDomain()
		if err != nil {
			continue
		}
		out = append(out, mdl)
	}
	return
}

func (d *dbQuery) resolveRate(rates []*models2.Rate, prev *rate.Rate, span rate.TimeSpan, tm time.Time) *rate.Rate {
	for _, rt := range rates {
		if prev == nil && tm.After(rt.UpdatedAt) {
			prev, _ = rt.ToDomain()
		}
		if tm.Format(span.Format()) == rt.UpdatedAt.Format(span.Format()) {
			mdl, _ := rt.ToDomain()
			return mdl
		}
	}
	cp := prev.Copy()
	cp.SetUpdatedAt(tm)
	return cp
}
func (d *dbQuery) History(ctx context.Context, span rate.TimeSpan, start, end time.Time) (out []*rate.Rate, err error) {
	db := d.apply(ctx)
	if !(start.IsZero() || end.IsZero()) {
		db = db.Where(rate.SavedBetween(start, end))
	}
	var res []*models2.Rate
	err = db.Order(fmt.Sprintf("%s %s", rate.UpdatedAtColumn, "ASC")).Find(&res).Error
	if err != nil || len(res) == 0 {
		return
	}
	// generate based on the records
	total := 0

	diff := start.Sub(end)
	switch span {
	case rate.Hourly:
		total = int(diff.Hours())
	case rate.Daily:
		total = int(math.Ceil(diff.Hours() / 24))
	case rate.FiveMinute:
		total = int(math.Ceil(diff.Minutes() / 5))
	case rate.Minute:
		total = int(math.Ceil(diff.Minutes()))
	}
	out = make([]*rate.Rate, total)
	for i := 0; i < total; i++ {
		var prev *rate.Rate
		if i != 0 {
			prev = out[(total)-i]
		}
		revIdx := (total-1)-i
		out[revIdx] = d.resolveRate(res, prev, span, start.Add(-(time.Duration(revIdx) * span.ToDuration())))
	}
	return
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
	db = db.WithContext(ctx).Model(&models2.Rate{})
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
	res := make([]*models2.Rate, 0)
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
	var res *models2.Rate
	err = db.WithContext(ctx).Where("id = ?", id).First(&res).Error
	if err == nil {
		out, err = res.ToDomain()
	}
	return
}
