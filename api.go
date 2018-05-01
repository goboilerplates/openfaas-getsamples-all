package main

import (
	"github.com/goboilerplates/core"
)

// GetSamplesAllAPI is the GetSamplesAllAPI interface.
type GetSamplesAllAPI interface {
	All() ([]*core.SampleEntity, error)
}

// GetSamplesAllAPIImpl is the implementation of GetSamplesAllAPI interface.
type GetSamplesAllAPIImpl struct {
	Interactor core.GetSamplesInteractor
}

// All gets all samples.
func (api GetSamplesAllAPIImpl) All() ([]*core.SampleEntity, error) {
	return api.Interactor.All()
}
