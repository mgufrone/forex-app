package rate

import (
	"encoding/json"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/mgufrone/forex/internal/shared/criteria"
	"github.com/mgufrone/forex/internal/shared/infrastructure/grpc/rate_service"
	"strconv"
	"time"
)

type GrpcCriteria struct {
	Filter *rate_service.RateFilter
}

func (g GrpcCriteria) Group(field string) criteria.ICriteriaBuilder {
	g.Filter.Group = append(g.Filter.Group, field)
	return g
}

func (g GrpcCriteria) Copy() criteria.ICriteriaBuilder {
	// TODO: make duplicate of filter
	return GrpcCriteria{Filter: &rate_service.RateFilter{}}
}

func (g GrpcCriteria) Select(fields ...string) criteria.ICriteriaBuilder {
	g.Filter.Select = fields
	return g
}

func (g GrpcCriteria) Paginate(page int, perPage int) criteria.ICriteriaBuilder {
	g.Filter.Page = int32(page)
	g.Filter.PerPage = int32(perPage)
	return g
}

func (g GrpcCriteria) Order(field string, direction string) criteria.ICriteriaBuilder {
	g.Filter.Sort = direction
	g.Filter.SortBy = field
	return g
}

func (g GrpcCriteria) Where(condition ...criteria.ICondition) criteria.ICriteriaBuilder {
	for _, w := range condition {
		fltr := &rate_service.RateQuery{
			Field:    w.Field(),
			Operator: int32(w.Operator()),
		}
		var val []byte
		switch v := w.Value().(type) {
		case json.Marshaler:
			val, _ = v.MarshalJSON()
			break
		case string:
			val = []byte(v)
		case time.Time:
			val = []byte(strconv.Itoa(int(v.Unix())))
		case int64:
			val = []byte(strconv.Itoa(int(v)))
		case int32:
			val = []byte(strconv.Itoa(int(v)))
		case float64:
			val = []byte(strconv.Itoa(int(v)))
		case float32:
			val = []byte(strconv.Itoa(int(v)))
		case int:
			val = []byte(strconv.Itoa(v))
		default:
			continue
		}
		fltr.Value = &any.Any{Value: val}
		g.Filter.Query = append(g.Filter.Query, fltr)
	}
	return g
}

func (g GrpcCriteria) And(other ...criteria.ICriteriaBuilder) criteria.ICriteriaBuilder {
	panic("implement me")
}

func (g GrpcCriteria) Or(other ...criteria.ICriteriaBuilder) criteria.ICriteriaBuilder {
	panic("implement me")
}

func (g GrpcCriteria) ToString() string {
	panic("implement me")
}

