package game

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewResponse(t *testing.T) {
	msg := "hello world"

	r := NewResponse(msg)
	assert.NotNil(t, r)
	assert.Equal(t, msg, r.GetMessage())
}
