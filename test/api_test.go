package function_test

import (
	"testing"

	"github.com/goboilerplates/core"
	"github.com/goboilerplates/openfaas-getsamples-all/function"
	"github.com/stretchr/testify/assert"
)

// TestGetSamplesAllAPIImplAll .
func TestGetSamplesAllAPIImplAll(test *testing.T) {
	api := function.GetSamplesAllAPIImpl{
		Interactor: core.CreateDefaultGetSamples("Hello"),
	}

	assert.NotNil(test, api)
	assert.NotNil(test, api.Interactor)

	result, err := api.All()
	assert.Nil(test, err)

	assert.NotNil(test, result)
	assert.True(test, len(result) > 0)
}
