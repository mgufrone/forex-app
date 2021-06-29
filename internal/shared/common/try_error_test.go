package common

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mocked struct {
	mock.Mock
}

func (m *mocked) mockRunner() error {
	return m.Called().Error(0)
}


func TestTryOrError(t *testing.T) {
	t.Parallel()
	try := TryOrError()
	assert.Nil(t, try)
	var m mocked
	m.On("mockRunner").Once().
		Return(errors.New("test error"))
	m.On("mockRunner").
		Return(nil)

	try = TryOrError(m.mockRunner, m.mockRunner, m.mockRunner)
	m.AssertNumberOfCalls(t, "mockRunner", 1)
	try = TryOrError(m.mockRunner, m.mockRunner, m.mockRunner)
	m.AssertNumberOfCalls(t, "mockRunner", 4)
}
