package rate

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
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
