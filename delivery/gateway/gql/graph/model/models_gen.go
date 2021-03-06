// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type Pagination struct {
	Page    int `json:"page"`
	PerPage int `json:"perPage"`
}

type QueryFilter struct {
	Pagination *Pagination `json:"pagination"`
	Sort       *Sort       `json:"sort"`
	Group      string      `json:"group"`
	Filter     *RateFilter `json:"filter"`
}

type Rate struct {
	ID         string    `json:"id"`
	Base       string    `json:"base"`
	Symbol     string    `json:"symbol"`
	Source     string    `json:"source"`
	SourceType string    `json:"sourceType"`
	Sell       float64   `json:"sell"`
	Buy        float64   `json:"buy"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

type RateFilter struct {
	Base        string       `json:"base"`
	Symbol      string       `json:"symbol"`
	Source      string       `json:"source"`
	SourceType  string       `json:"sourceType"`
	DateBetween []*time.Time `json:"dateBetween"`
	Date        time.Time    `json:"date"`
}

type Sort struct {
	SortBy string        `json:"sortBy"`
	Sort   SortDirection `json:"sort"`
}

type SortDirection string

const (
	SortDirectionDesc SortDirection = "DESC"
	SortDirectionAsc  SortDirection = "ASC"
)

var AllSortDirection = []SortDirection{
	SortDirectionDesc,
	SortDirectionAsc,
}

func (e SortDirection) IsValid() bool {
	switch e {
	case SortDirectionDesc, SortDirectionAsc:
		return true
	}
	return false
}

func (e SortDirection) String() string {
	return string(e)
}

func (e *SortDirection) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SortDirection(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SortDirection", str)
	}
	return nil
}

func (e SortDirection) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
