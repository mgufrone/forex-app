package models

import (
	"fmt"
	"github.com/mgufrone/forex/internal/domains/rate"
	"time"
)

type Rate struct {
	ID         string    `json:"id" gorm:"primaryKey;autoIncrement:false"`
	Base       string    `json:"base" gorm:"index:,sort:asc"`
	Symbol     string    `json:"symbol" gorm:"index:,sort:asc"`
	Source     string    `json:"source" gorm:"index"`
	SourceType string    `json:"source_type" gorm:"index"`
	Sell       float64   `json:"sell" gorm:"type:decimal(19,10)"`
	Buy        float64   `json:"buy" gorm:"type:decimal(19,10)"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"index:,sort:desc"`
}

func (r *Rate) FromDomain(rate *rate.Rate) {
	if rate == nil {
		return
	}
	r.ID = rate.ID()
	r.Buy = rate.Buy()
	r.Sell = rate.Sell()
	r.SourceType = rate.SourceType()
	r.Source = rate.Source()
	r.Base = rate.Base()
	r.Symbol = rate.Symbol()
	r.UpdatedAt = rate.UpdatedAt()
}
func (r *Rate) ToDomain() (rt *rate.Rate, err error) {
	fmt.Println("values", r.Sell, r.Buy)
	rt, err = rate.NewRate(r.Base, r.Symbol, r.Source, r.SourceType, r.Sell, r.Buy, r.UpdatedAt)
	if err != nil {
		return nil, err
	}
	rt.SetID(r.ID)
	return
}
