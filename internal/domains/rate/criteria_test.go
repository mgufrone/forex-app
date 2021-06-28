package rate

import (
	"github.com/mgufrone/forex/internal/shared/criteria"
	"github.com/stretchr/testify/assert"
	"testing"
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
