package mock

import (
	"github.com/mgufrone/forex/internal/shared/criteria"
	"github.com/stretchr/testify/mock"
)

type CriteriaMock struct {
	mock.Mock
}

func (c *CriteriaMock) Group(field string) criteria.ICriteriaBuilder {
	panic("implement me")
}

func (c *CriteriaMock) Copy() criteria.ICriteriaBuilder {
	panic("implement me")
}

func (c *CriteriaMock) Select(fields ...string) criteria.ICriteriaBuilder {
	args := make([]interface{}, len(fields))
	for _, c1 := range fields {
		args = append(args, c1)
	}
	c.Called(args...)
	return c
}

func (c *CriteriaMock) Paginate(page int, perPage int) criteria.ICriteriaBuilder {
	c.Called(page, perPage)
	return c
}

func (c *CriteriaMock) Order(field string, direction string) criteria.ICriteriaBuilder {
	c.Called(field, direction)
	return c
}

func (c *CriteriaMock) Where(condition ...criteria.ICondition) criteria.ICriteriaBuilder {
	args := make([]interface{}, len(condition))
	for idx, c1 := range condition {
		args[idx] = c1
	}
	c.Called(args...)
	return c
}

func (c *CriteriaMock) And(other ...criteria.ICriteriaBuilder) criteria.ICriteriaBuilder {
	args := make([]interface{}, len(other))
	for idx, c1 := range other {
		args[idx] = c1
	}
	c.Called(args...)
	return c
}

func (c *CriteriaMock) Or(other ...criteria.ICriteriaBuilder) criteria.ICriteriaBuilder {
	args := make([]interface{}, len(other))
	for idx, c1 := range other {
		args[idx] = c1
	}
	c.Called(args...)
	return c
}

func (c *CriteriaMock) ToString() string {
	return c.Called().Get(0).(string)
}

