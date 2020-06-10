package error

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProcessNotFoundError_Error(t *testing.T) {
	e := NewProcessNotFoundError("deine-mudda.exe")
	assert.NotNil(t, e)
	assert.Equal(t, "process \"deine-mudda.exe\" not found", e.Error())
}
