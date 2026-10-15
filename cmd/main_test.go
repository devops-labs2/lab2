package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	got, err := hello()
	assert.NoError(t, err)
	assert.Equal(t, response{Message: "Hello, world!"}, got)
}
