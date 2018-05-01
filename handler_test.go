package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestHandle .
func TestHandle(test *testing.T) {
	result := Handle(nil)

	assert.NotNil(test, result)
	assert.True(test, len(result) > 0)
}

// TestCreateAPI .
func TestCreateAPI(test *testing.T) {
	result := CreateAPI()
	assert.NotNil(test, result)
}
