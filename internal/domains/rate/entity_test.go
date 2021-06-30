package rate

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"time"
)

func TestRate_SetBase(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		in          string
		expected    string
		shouldError bool
	}{
		{"", "", true},
		{strings.Repeat("abc", 5), "", false},
		{"abc", "abc", false},
	}
	for _, c := range testCases {
		var r Rate
		err := r.SetBase(c.in)
		if c.shouldError {
			assert.NotNil(t, err)
			continue
		}
		assert.Equal(t, c.expected, r.Base())
	}
}
func TestRate_SetSymbol(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		in          string
		expected    string
		shouldError bool
	}{
		{"", "", true},
		{strings.Repeat("abc", 5), "", false},
		{"abc", "abc", false},
	}
	for _, c := range testCases {
		var r Rate
		err := r.SetSymbol(c.in)
		if c.shouldError {
			assert.NotNil(t, err)
			continue
		}
		assert.Equal(t, c.expected, r.Symbol())
	}
}

func TestRate_SetSource(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		in          string
		expected    string
		shouldError bool
	}{
		{"", "", true},
		{"abc", "abc", false},
	}
	for _, c := range testCases {
		var r Rate
		err := r.SetSource(c.in)
		if c.shouldError {
			assert.NotNil(t, err)
			continue
		}
		assert.Equal(t, c.expected, r.Source())
	}
}
func TestRate_SetSourceType(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		in          string
		expected    string
		shouldError bool
	}{
		{"", "", true},
		{"abc", "abc", false},
	}
	for _, c := range testCases {
		var r Rate
		err := r.SetSourceType(c.in)
		if c.shouldError {
			assert.NotNil(t, err)
			continue
		}
		assert.Equal(t, c.expected, r.SourceType())
	}
}
func TestRate_SetSell(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		in          float64
		expected    float64
		shouldError bool
	}{
		{0, 0, true},
		{-10, 0, true},
		{1.121212, 1.121212, false},
		{0.1212, 0.1212, false},
	}
	for _, c := range testCases {
		var r Rate
		err := r.SetSell(c.in)
		if c.shouldError {
			assert.NotNil(t, err)
			continue
		}
		assert.Equal(t, c.expected, r.Sell())
	}
}
func TestRate_SetBuy(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		in          float64
		expected    float64
		shouldError bool
	}{
		{0, 0, true},
		{-10, 0, true},
		{1.121212, 1.121212, false},
		{0.1212, 0.1212, false},
	}
	for _, c := range testCases {
		var r Rate
		err := r.SetBuy(c.in)
		if c.shouldError {
			assert.NotNil(t, err)
			continue
		}
		assert.Equal(t, c.expected, r.Buy())
	}
}
func TestRate_SetUpdatedAt(t *testing.T) {
	t.Parallel()
	now := time.Now()
	testCases := []struct {
		in          time.Time
		expected    time.Time
		shouldError bool
	}{
		{time.Time{}, time.Time{}, true},
		{now, now, false},
		{now.Add(-time.Hour), now.Add(-time.Hour), false},
	}
	for _, c := range testCases {
		var r Rate
		err := r.SetUpdatedAt(c.in)
		if c.shouldError {
			assert.NotNil(t, err)
			continue
		}
		assert.Equal(t, c.expected, r.UpdatedAt(), c.in)
	}
}
