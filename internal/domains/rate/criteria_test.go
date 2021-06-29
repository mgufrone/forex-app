package rate

import (
	"github.com/mgufrone/forex/internal/shared/criteria"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestWhereBase(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		in       string
		expected criteria.ICondition
	}{
		{"", nil},
		{"a", criteria.NewCondition(BaseColumn, criteria.Eq, "a")},
	}
	for _, c := range testCases {
		w := WhereBase(c.in)
		assert.Equal(t, w, c.expected)
	}
}

func TestWhereSymbol(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		in       string
		expected criteria.ICondition
	}{
		{"", nil},
		{"a", criteria.NewCondition(SymbolColumn, criteria.Eq, "a")},
	}
	for _, c := range testCases {
		w := WhereSymbol(c.in)
		assert.Equal(t, w, c.expected)
	}
}

func TestWhereSource(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		in       string
		expected criteria.ICondition
	}{
		{"", nil},
		{"a", criteria.NewCondition(SourceColumn, criteria.Eq, "a")},
	}
	for _, c := range testCases {
		w := WhereSource(c.in)
		assert.Equal(t, w, c.expected)
	}
}

func TestWhereSourceType(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		in       string
		expected criteria.ICondition
	}{
		{"", nil},
		{"a", criteria.NewCondition(SourceTypeColumn, criteria.Eq, "a")},
	}
	for _, c := range testCases {
		w := WhereSourceType(c.in)
		assert.Equal(t, w, c.expected)
	}
}

func TestSavedAt(t *testing.T) {
	t.Parallel()
	now := time.Now()
	testCases := []struct {
		in       time.Time
		expected criteria.ICondition
	}{
		{time.Time{}, nil},
		{now, criteria.NewCondition(UpdatedAtColumn, criteria.Eq, now)},
	}
	for _, c := range testCases {
		w := SavedAt(c.in)
		assert.Equal(t, w, c.expected)
	}
}

func TestSavedBetween(t *testing.T) {
	t.Parallel()
	now := time.Now()
	before := now.Add(-time.Hour)
	testCases := []struct {
		in       []time.Time
		expected criteria.ICondition
	}{
		{[]time.Time{{}, time.Now()}, nil},
		{[]time.Time{before, {}}, nil},
		{[]time.Time{before, now}, criteria.NewCondition(UpdatedAtColumn, criteria.Between, []time.Time{before, now})},
	}
	for _, c := range testCases {
		w := SavedBetween(c.in[0], c.in[1])
		assert.Equal(t, w, c.expected)
	}
}
