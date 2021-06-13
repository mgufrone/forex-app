package criteria

import (
	"fmt"
)

type Operator int

const (
	Eq Operator = iota
	Not
	Like
	NotLike
	In
	NotIn
	Gt
	Gte
	Lt
	Lte
	Between
)

type ICondition interface {
	Field() string
	Operator() Operator
	Value() interface{}
	ToString() string
}

type ICriteriaBuilder interface {
	Copy() ICriteriaBuilder
	Select(fields ...string) ICriteriaBuilder
	Paginate(page int, perPage int) ICriteriaBuilder
	Order(field string, direction string) ICriteriaBuilder
	// by default, it will run ands
	Where(condition ...ICondition) ICriteriaBuilder
	And(other ...ICriteriaBuilder) ICriteriaBuilder
	Or(other ...ICriteriaBuilder) ICriteriaBuilder
	ToString() string
}

type baseCondition struct {
	field    string
	value    interface{}
	operator Operator
}

func NewCondition(field string, operator Operator, value interface{}) ICondition {
	return baseCondition{field, value, operator}
}

func (b baseCondition) Field() string {
	return b.field
}

func (b baseCondition) Operator() Operator {
	return b.operator
}

func (b baseCondition) Value() interface{} {
	return b.value
}

func (b baseCondition) ToString() string {
	return fmt.Sprintf("%s;%d=%v", b.Field(), b.Operator(), b.Value())
}
