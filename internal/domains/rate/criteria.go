package rate

import (
	"github.com/mgufrone/forex/internal/shared/criteria"
	"time"
)

func SavedAt(date time.Time) criteria.ICondition {
	if date.IsZero() {
		return nil
	}
	return criteria.NewCondition(UpdatedAtColumn, criteria.Eq, date)
}

func WhereSource(source string) criteria.ICondition {
	if source == "" {
		return nil
	}
	return criteria.NewCondition(SourceColumn, criteria.Eq, source)
}

func WhereSourceType(sourceType string) criteria.ICondition {
	if sourceType == "" {
		return nil
	}
	return criteria.NewCondition(SourceTypeColumn, criteria.Eq, sourceType)
}
func WhereSymbol(symbol string) criteria.ICondition {
	if symbol == "" {
		return nil
	}
	return criteria.NewCondition(SymbolColumn, criteria.Eq, symbol)
}
func WhereBase(base string) criteria.ICondition {
	if base == "" {
		return nil
	}
	return criteria.NewCondition(BaseColumn, criteria.Eq, base)
}

func SavedBetween(start, end time.Time) criteria.ICondition {
	if start.IsZero() || end.IsZero() {
		return nil
	}
	return criteria.NewCondition(UpdatedAtColumn, criteria.Between, []time.Time{start, end})
}
