package repository

import (
	"fmt"
	"github.com/mgufrone/forex/internal/shared/criteria"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type dbCriteria struct {
	mods []qm.QueryMod
}

func newDbCriteria() criteria.ICriteriaBuilder {
	return dbCriteria{mods: []qm.QueryMod{}}
}

func (d dbCriteria) Copy() criteria.ICriteriaBuilder {
	e := newDbCriteria()
	e.(*dbCriteria).mods = d.mods
	return e
}

func (d dbCriteria) Select(fields ...string) criteria.ICriteriaBuilder {
	d.mods = append(d.mods, qm.Select(fields...))
	return d
}

func (d dbCriteria) Paginate(page int, perPage int) criteria.ICriteriaBuilder {
	d.mods = append(d.mods, qm.Limit(perPage), qm.Offset((page-1)*perPage))
	return d
}

func (d dbCriteria) Order(field string, direction string) criteria.ICriteriaBuilder {
	d.mods = append(d.mods, qm.OrderBy(fmt.Sprintf("%s %s", field, direction)))
	return d
}

func (d dbCriteria) Where(condition ...criteria.ICondition) criteria.ICriteriaBuilder {
	for _, c := range condition {
		var q qm.QueryMod
		switch c.Operator() {
		case criteria.In:
			v := c.Value().([]interface{})
			q = qm.WhereIn(fmt.Sprintf("%s in ?", c.Field()), v...)
			break
		case criteria.NotIn:
			v := c.Value().([]interface{})
			q = qm.WhereNotIn(fmt.Sprintf("%s not in ?", c.Field()), v...)
			break
		case criteria.Eq:
			q = qm.Where(fmt.Sprintf("%s = ?", c.Field()), c.Value())
			break
		case criteria.Not:
			q = qm.Where(fmt.Sprintf("%s != ?", c.Field()), c.Value())
			break
		case criteria.Gt:
			q = qm.Where(fmt.Sprintf("%s > ?", c.Field()), c.Value())
			break
		case criteria.Gte:
			q = qm.Where(fmt.Sprintf("%s >= ?", c.Field()), c.Value())
			break
		case criteria.Lte:
			q = qm.Where(fmt.Sprintf("%s <= ?", c.Field()), c.Value())
			break
		case criteria.Lt:
			q = qm.Where(fmt.Sprintf("%s < ?", c.Field()), c.Value())
			break
		case criteria.Between:
			v := c.Value().([]interface{})
			q = qm.Where(fmt.Sprintf("%s between ? and ?", c.Field()), v[0], v[1])
			break
		case criteria.Like:
			q = qm.Where(fmt.Sprintf("%s ilike ?", c.Field()), fmt.Sprintf(`%%%s%%`, c.Value()))
			break
		case criteria.NotLike:
			q = qm.Where(fmt.Sprintf("%s not ilike ?", c.Field()), fmt.Sprintf(`%%%s%%`, c.Value()))
			break
		}
		d.mods = append(d.mods, q)
	}
	return d
}

func (d dbCriteria) And(other ...criteria.ICriteriaBuilder) criteria.ICriteriaBuilder {
	for _, o := range other {
		mds := o.(dbCriteria).mods
		expr := qm.Expr(mds...)
		d.mods = append(d.mods, expr)
	}
	return d
}

func (d dbCriteria) Or(other ...criteria.ICriteriaBuilder) criteria.ICriteriaBuilder {
	for _, o := range other {
		mds := o.(dbCriteria).mods
		expr := qm.Or2(qm.Expr(mds...))
		d.mods = append(d.mods, expr)
	}
	return d
}

func (d dbCriteria) ToString() string {
	panic("implement me")
}
