package lab2

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"strings"
	"bytes"
	
)

func TestForEmptyInput(t *testing.T) {
	handler := ComputeHandler{
		Reader: strings.NewReader(""),
		Writer: bytes.NewBuffer(nil),
	}
	err := handler.Compute()
	assert.NotNil(t, err)
}

func TestForCorrectInput(t *testing.T) {
	handler := ComputeHandler{
		Reader: strings.NewReader("2 2 +"),
		Writer: bytes.NewBuffer(nil),
	}
	err := handler.Compute()
	assert.Nil(t, err)
}

func TestForIncorrectInput(t *testing.T) {
	handler := ComputeHandler{
		Reader: strings.NewReader("2 2 %"),
		Writer: bytes.NewBuffer(nil),
	}
	err := handler.Compute()
	assert.NotNil(t, err)
}