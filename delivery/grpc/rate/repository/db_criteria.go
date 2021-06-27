package repository

import (
	"context"
	"fmt"
	"github.com/mgufrone/forex/internal/shared/criteria"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"strings"
)

type dbCriteriaBuilder struct {
	conditions []criteria.ICondition
	ands       []criteria.ICriteriaBuilder
	ors        []criteria.ICriteriaBuilder
	pagination []int
	selects    []string
	sort       [][]string
	group      []string
}

func (d dbCriteriaBuilder) Group(field string) criteria.ICriteriaBuilder {
	d.group = append(d.group, field)
	return d
}

func (d dbCriteriaBuilder) Paginate(page int, perPage int) criteria.ICriteriaBuilder {
	d.pagination = []int{page, perPage}
	return d
}

func (d dbCriteriaBuilder) Order(field string, direction string) criteria.ICriteriaBuilder {
	d.sort = append(d.sort, []string{field, direction})
	return d
}

func (d dbCriteriaBuilder) Copy() criteria.ICriteriaBuilder {
	return dbCriteriaBuilder{}
}

func (d dbCriteriaBuilder) Select(fields ...string) criteria.ICriteriaBuilder {
	d.selects = append(d.selects, fields...)
	return d
}

func (d dbCriteriaBuilder) Where(condition ...criteria.ICondition) criteria.ICriteriaBuilder {
	for _, r := range condition {
		if r == nil {
			continue
		}
		d.conditions = append(d.conditions, r)
	}
	return d
}

func (d dbCriteriaBuilder) And(other ...criteria.ICriteriaBuilder) criteria.ICriteriaBuilder {
	d.ands = append(d.ands, other...)
	return d
}

func (d dbCriteriaBuilder) Or(other ...criteria.ICriteriaBuilder) criteria.ICriteriaBuilder {
	d.ors = append(d.ors, other...)
	return d
}

func (d dbCriteriaBuilder) ToString() (res string) {
	concats := make([]string, 0)
	if len(d.pagination) > 0 {
		concats = append(concats, fmt.Sprintf("page:%d;per_page:%d;", d.pagination[0], d.pagination[1]))
	}
	for _, s := range d.sort {
		concats = append(concats, "sort:%s-%s", s[0], s[1])
	}
	for _, s := range d.selects {
		concats = append(concats, "select:%s", s)
	}
	for _, s := range d.group {
		concats = append(concats, "group:%s", s)
	}
	for _, r := range d.conditions {
		concats = append(concats, r.ToString())
	}
	if len(d.ands) > 0 {
		var ands []string
		for _, a := range d.ands {
			ands = append(ands, a.ToString())
		}
		concats = append(concats, fmt.Sprintf("ands(%s)", strings.Join(ands, ",")))
	}
	if len(d.ors) > 0 {
		var ands []string
		for _, a := range d.ors {
			ands = append(ands, a.ToString())
		}
		concats = append(concats, fmt.Sprintf("ors(%s)", strings.Join(ands, ",")))
	}
	res = strings.Join(concats, ";")
	return
}

func operator(condition criteria.ICondition) string {
	switch condition.Operator() {
	case criteria.Eq:
		return "="
	case criteria.Not:
		return "!="
	case criteria.Like:
		return "like"
	case criteria.NotLike:
		return "not like"
	case criteria.In:
		return "in"
	case criteria.NotIn:
		return "not in"
	}
	return ""
}
func value(operator criteria.ICondition) interface{} {
	if operator.Operator() == criteria.Like || operator.Operator() == criteria.NotLike {
		return fmt.Sprintf("%%%s%%", operator.Value())
	}
	return operator.Value()
}

func (d dbCriteriaBuilder) apply(db *gorm.DB) *gorm.DB {
	ori := db
	if len(d.pagination) > 0 {
		db = db.
			Limit(d.pagination[1]).
			Offset((d.pagination[0] - 1) * d.pagination[1])
	}
	if len(d.sort) > 0 {
		for _, srt := range d.sort {
			isDesc := true
			if srt[0] == "asc" {
				isDesc = false
			}
			db = db.Order(clause.OrderByColumn{
				Column: clause.Column{Name: srt[0]},
				Desc:   isDesc,
			})
		}
	}
	if len(d.group) > 0 {
		for _, g := range d.group {
			db = db.Group(g)
		}
	}
	if len(d.selects) > 0 {
		db = db.Select(d.selects)
	}
	if len(d.conditions) > 0 {
		ses := db.WithContext(context.TODO())
		for _, r := range d.conditions {
			ses = ses.Where(fmt.Sprintf("%s %s ?", r.Field(), operator(r)), value(r))
		}
		db = db.Where(ses)
	}
	if len(d.ands) > 0 {
		tx := ori.WithContext(context.TODO())
		for _, a := range d.ands {
			ses := a.(dbCriteriaBuilder).apply(ori.WithContext(context.TODO()))
			tx = tx.Where(ses)
		}
		db = db.Where(tx)
	}
	if len(d.ors) > 0 {
		tx := ori.WithContext(context.TODO())
		for idx, a := range d.ors {
			ses := a.(dbCriteriaBuilder).apply(ori.WithContext(context.TODO()))
			if idx == 0 {
				tx = tx.Where(ses)
			} else {
				tx = tx.Or(ses)
			}
		}
		db = db.Where(tx)
	}
	return db
}
