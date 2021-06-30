package rate_service

import (
	"errors"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/mgufrone/forex/internal/domains/rate"
)

func (x *SpanFilter) Validate() error {
	if x == nil {
		return errors.New("invalid value")
	}
	ins := []int32{
		int32(rate.Daily),
		int32(rate.Minute),
		int32(rate.Hourly),
		int32(rate.FiveMinute),
		int32(rate.TenMinute),
	}
	return validation.ValidateStruct(
		x,
		validation.Field(&x.End, validation.Required),
		validation.Field(&x.Start, validation.Required),
		validation.Field(&x.Start, validation.Max(x.End)),
		validation.Field(&x.End, validation.Min(x.End)),
		validation.Field(&x.Span, validation.In(ins)),
	)
}
