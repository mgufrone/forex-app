package rate

import (
	"github.com/mgufrone/forex/internal/shared/criteria"
	"time"
)

func SavedAt(date time.Time) criteria.ICondition {
	return criteria.NewCondition(updatedAtColumn, criteria.Eq, date)
}

func WhereSource(source string) criteria.ICondition {
	return criteria.NewCondition(sourceColumn, criteria.Eq, source)
}

func WhereSourceType(sourceType string) criteria.ICondition {
	return criteria.NewCondition(sourceTypeColumn, criteria.Eq, sourceType)
}
func WhereSymbol(symbol string) criteria.ICondition {
	return criteria.NewCondition(symbolColumn, criteria.Eq, symbol)
}

func SavedBetween(start, end time.Time) criteria.ICondition {
	return criteria.NewCondition(updatedAtColumn, criteria.Between, []time.Time{start, end})
}
