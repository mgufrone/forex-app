package rate

import "time"

const (
	DateLayout = "2006-01-02"
)

type TimeSpan int

const (
	Daily TimeSpan = iota
	Hourly
	Minute
	FiveMinute
	TenMinute
)

func (t TimeSpan) ToDuration() time.Duration {
	switch t {
	case Hourly:
		return time.Hour
	case Minute:
		return time.Minute
	case FiveMinute:
		return 5 * time.Minute
	case TenMinute:
		return 10 * time.Minute
	}
	return 24 * time.Hour
}
func (t TimeSpan) Format() string {
	switch t {
	case Hourly:
		return "2006-01-02 03"
	case TenMinute, Minute, FiveMinute:
		return "2006-01-02 03:04"
	}
	return "2006-01-02"
}