package rate

import (
	"encoding/json"
	"errors"
	"github.com/mgufrone/forex/internal/shared/common"
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

func (r *Rate) SetSymbol(symbol string) (err error) {
	if symbol == "" || len(symbol) > 5 {
		return errors.New("invalid symbol value")
	}
	r.symbol = symbol
	return
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

func (r *Rate) SetSourceType(sourceType string) (err error) {
	if sourceType == "" {
		return errors.New("invalid sourceType value")
	}
	r.sourceType = sourceType
	return
}

func (r *Rate) Sell() float64 {
	return r.sell
}

func (r *Rate) SetSell(sell float64) (err error) {
	if !(sell > 0.0) {
		return errors.New("invalid sell value")
	}
	r.sell = sell
	return
}

func (r *Rate) Buy() float64 {
	return r.buy
}

func (r *Rate) SetBuy(buy float64) (err error){
	if !(buy > 0.0) {
		return errors.New("invalid buy value")
	}
	r.buy = buy
	return
}

func (r *Rate) UpdatedAt() time.Time {
	return r.updatedAt
}

func (r *Rate) SetUpdatedAt(date time.Time) (err error) {
	if date.IsZero() {
		return errors.New("invalid updated_at value")
	}
	r.updatedAt = date
	return
}

func (r *Rate) Copy() (*Rate, error) {
	rt, err := NewRate(
		r.Base(),
		r.Symbol(),
		r.Source(),
		r.SourceType(),
		r.Sell(),
		r.Buy(),
		r.UpdatedAt(),
	)
	if err != nil {
		return nil, err
	}
	rt.SetID(r.ID())
	return rt, nil
}

func MustNew(base, symbol, source, sourceType string, sell, buy float64, date time.Time) (res *Rate) {
	var err error
	if res, err = NewRate(base, symbol, source, sourceType, sell, buy, date); err != nil {
		panic(err)
	}
	return
}
func NewRate(base, symbol, source, sourceType string, sell, buy float64, date time.Time) (*Rate, error) {
	var rt Rate
	if err := common.TryOrError(func() error {
		return rt.SetBase(base)
	}, func() error {
		return rt.SetSymbol(symbol)
	}, func() error {
		return rt.SetSource(source)
	}, func() error {
		return rt.SetSourceType(sourceType)
	}, func() error {
		return rt.SetSell(sell)
	}, func() error {
		return rt.SetBuy(buy)
	}, func() error {
		return rt.SetUpdatedAt(date)
	}); err != nil {
		return nil, err
	}
	return &rt, nil
}
