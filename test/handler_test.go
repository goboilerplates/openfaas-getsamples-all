package function_test

import (
	"testing"

	"github.com/goboilerplates/openfaas-getsamples-all/function"
	"github.com/stretchr/testify/assert"
)

// TestHandle .
func TestHandle(test *testing.T) {
	result := function.Handle(nil)

	assert.NotNil(test, result)
	assert.True(test, len(result) > 0)
}

// TestCreateAPI .
func TestCreateAPI(test *testing.T) {
	result := function.CreateAPI()
	assert.NotNil(test, result)
}
