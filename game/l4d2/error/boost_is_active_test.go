package error

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBoostIsActiveError_Error(t *testing.T) {
	e := NewBoostIsActiveError()
	assert.NotNil(t, e)
	assert.Equal(t, "boost is already active", e.Error())
}
