package model

import "github.com/mgufrone/forex/internal/domains/rate"

func (m *Rate) FromDomain(rate *rate.Rate) {
	if rate == nil {
		return
	}
	m.Source = rate.Symbol()
	m.Buy = rate.Buy()
	m.SourceType = rate.SourceType()
	m.Sell = rate.Sell()
	m.UpdatedAt = rate.UpdatedAt()
	m.ID = rate.ID()
	m.Base = rate.Base()
}

func (m *Rate) ToDomain() *rate.Rate {
	if m == nil {
		return nil
	}
	rt, _ := rate.NewRate(
		m.Base,
		m.Symbol,
		m.Source,
		m.SourceType,
		m.Sell,
		m.Buy,
		m.UpdatedAt,
	)
	rt.SetID(m.ID)
	return rt
}