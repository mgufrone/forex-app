package rate_service

import (
	"github.com/mgufrone/forex/internal/domains/rate"
	"time"
)

func (x *Rate) FromDomain(in *rate.Rate) {
	if in == nil {
		return
	}

	x.Base = in.Base()
	x.Source = in.Source()
	x.SourceType = in.SourceType()
	x.Buy = in.Buy()
	x.Sell = in.Sell()
	x.Symbol = in.Symbol()
	x.UpdatedAt = in.UpdatedAt().Unix()
	x.Id = in.ID()
}
func (x *Rate) ToDomain() *rate.Rate {
	t := time.Unix(x.GetUpdatedAt(), 0)
	rt := rate.NewRate(
		x.GetBase(),
		x.GetSymbol(),
		x.GetSource(),
		x.GetSourceType(),
		x.GetBuy(),
		x.GetSell(),
		t,
	)
	if x.GetId() != "" {
		rt.SetID(x.GetId())
	}
	return rt
}