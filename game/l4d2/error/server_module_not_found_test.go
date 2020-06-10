package error

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestServerModuleNotFoundError_Error(t *testing.T) {
	err := NewServerModuleNotFoundError()
	assert.NotNil(t, err)
	assert.Equal(t, "server module not found", err.Error())
}
