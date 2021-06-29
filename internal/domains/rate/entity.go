package rate

import (
	"encoding/json"
	"errors"
	"time"
)

const (
	IDColumn         = "id"
	BaseColumn       = "base"
	SymbolColumn     = "symbol"
	SourceColumn     = "source"
	SourceTypeColumn = "source_type"
	BuyColumn        = "buy"
	SellColumn       = "sell"
	UpdatedAtColumn  = "updated_at"
)

type Rate struct {
	id         string
	base       string
	symbol     string
	source     string
	sourceType string
	sell       float64
	buy        float64
	updatedAt  time.Time
}

func (r *Rate) ID() string {
	return r.id
}

func (r *Rate) SetID(id string) {
	r.id = id
}

func (r *Rate) UnmarshalJSON(bytes []byte) error {
	var maps map[string]interface{}
	err := json.Unmarshal(bytes, &maps)
	if err != nil {
		return err
	}


	if base, ok := maps[BaseColumn].(string); ok {
		r.SetBase(base)
	}

	if symbol, ok := maps[SymbolColumn].(string); ok {
		r.SetSymbol(symbol)
	}

	if source, ok := maps[SourceColumn].(string); ok {
		r.SetSource(source)
	}

	if sourceType, ok := maps[SourceTypeColumn].(string); ok {
		r.SetSourceType(sourceType)
	}

	if buy, ok := maps[BuyColumn].(float64); ok {
		r.SetBuy(buy)
	}

	if sell, ok := maps[SellColumn].(float64); ok {
		r.SetBuy(sell)
	}

	if updatedAt, ok := maps[SellColumn].(int64); ok {
		r.SetUpdatedAt(time.Unix(updatedAt, 0))
	}

	return nil
}

func (r *Rate) MarshalJSON() ([]byte, error) {
	maps := map[string]interface{}{
		IDColumn:         r.ID(),
		BaseColumn:       r.Base(),
		SymbolColumn:     r.Symbol(),
		SourceColumn:     r.Source(),
		SourceTypeColumn: r.SourceType(),
		BuyColumn:        r.Buy(),
		SellColumn:       r.Sell(),
		UpdatedAtColumn:  r.UpdatedAt().Unix(),
	}

	return json.Marshal(maps)
}

func (r *Rate) Base() string {
	return r.base
}

func (r *Rate) SetBase(base string) (err error) {
	if base == "" || len(base) > 5 {
		return errors.New("invalid base value")
	}
	r.base = base
	return
}

func (r *Rate) Symbol() string {
	return r.symbol
}

func (r *Rate) SetSymbol(symbol string) {
	r.symbol = symbol
}

func (r *Rate) Source() string {
	return r.source
}

func (r *Rate) SetSource(source string) (err error) {
	if source == "" {
		return errors.New("invalid source value")
	}
	r.source = source
	return
}

func (r *Rate) SourceType() string {
	return r.sourceType
}

func (r *Rate) SetSourceType(sourceType string) {
	r.sourceType = sourceType
}

func (r *Rate) Sell() float64 {
	return r.sell
}

func (r *Rate) SetSell(sell float64) {
	r.sell = sell
}

func (r *Rate) Buy() float64 {
	return r.buy
}

func (r *Rate) SetBuy(buy float64) {
	r.buy = buy
}

func (r *Rate) UpdatedAt() time.Time {
	return r.updatedAt
}

func (r *Rate) SetUpdatedAt(date time.Time) {
	r.updatedAt = date
}

func (r *Rate) Copy() *Rate {
	rt := NewRate(
		r.Base(),
		r.Symbol(),
		r.Source(),
		r.SourceType(),
		r.Sell(),
		r.Buy(),
		r.UpdatedAt(),
	)
	rt.SetID(r.ID())
	return rt
}

func NewRate(base, symbol, source, sourceType string, sell float64, buy float64, date time.Time) *Rate {
	return &Rate{base: base, symbol: symbol, source: source, sourceType: sourceType, sell: sell, buy: buy, updatedAt: date}
}
