package rate

import (
	"encoding/json"
	"time"
)

const (
	baseColumn       = "base"
	symbolColumn     = "symbol"
	sourceColumn     = "source"
	sourceTypeColumn = "source_type"
	buyColumn        = "buy"
	sellColumn       = "sell"
	updatedAtColumn  = "updated_at"
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

	if base, ok := maps[baseColumn].(string); ok {
		r.SetBase(base)
	}

	if symbol, ok := maps[symbolColumn].(string); ok {
		r.SetSymbol(symbol)
	}

	if source, ok := maps[sourceColumn].(string); ok {
		r.SetSource(source)
	}

	if sourceType, ok := maps[sourceTypeColumn].(string); ok {
		r.SetSourceType(sourceType)
	}

	if buy, ok := maps[buyColumn].(float64); ok {
		r.SetBuy(buy)
	}

	if sell, ok := maps[sellColumn].(float64); ok {
		r.SetBuy(sell)
	}

	if updatedAt, ok := maps[sellColumn].(int64); ok {
		r.SetUpdatedAt(time.Unix(updatedAt, 0))
	}

	return nil
}

func (r *Rate) MarshalJSON() ([]byte, error) {
	maps := map[string]interface{}{
		baseColumn:       r.Base(),
		symbolColumn:     r.Symbol(),
		sourceColumn:     r.Source(),
		sourceTypeColumn: r.SourceType(),
		buyColumn:        r.Buy(),
		sellColumn:       r.Sell(),
		updatedAtColumn:  r.UpdatedAt().Unix(),
	}

	return json.Marshal(maps)
}

func (r *Rate) Base() string {
	return r.base
}

func (r *Rate) SetBase(base string) {
	r.base = base
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

func (r *Rate) SetSource(source string) {
	r.source = source
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

func NewRate(base, symbol, source, sourceType string, sell float64, buy float64, date time.Time) *Rate {
	return &Rate{base: base, symbol: symbol, source: source, sourceType: sourceType, sell: sell, buy: buy, updatedAt: date}
}
