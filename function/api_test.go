package function

import (
	"testing"

	"github.com/goboilerplates/core"
	"github.com/stretchr/testify/assert"
)

// TestGetSamplesAllAPIImplAll .
func TestGetSamplesAllAPIImplAll(test *testing.T) {
	api := GetSamplesAllAPIImpl{
		Interactor: core.CreateDefaultGetSamples("Hello"),
	}

	assert.NotNil(test, api)
	assert.NotNil(test, api.Interactor)

	result, err := api.All()
	assert.Nil(test, err)

	assert.NotNil(test, result)
	assert.True(test, len(result) > 0)
}
