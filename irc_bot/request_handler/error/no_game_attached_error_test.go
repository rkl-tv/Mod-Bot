package error

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNoGameAttachedError_Error(t *testing.T) {
	e := NewNoGameAttachedError()
	assert.NotNil(t, e)
	assert.Equal(t, "no game attached", e.Error())
}
