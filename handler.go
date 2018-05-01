package main

import (
	"encoding/json"
	"log"

	"github.com/goboilerplates/core"
)

var (
	api GetSamplesAllAPI
)

// Handle a serverless request .
func Handle(req []byte) string {
	log.Println("Request with ", req)
	api = CreateAPI()
	result, err := api.All()
	if err != nil {
		return err.Error()
	}
	jsonBytes, err := json.Marshal(result)
	if err != nil {
		return err.Error()
	}
	return string(jsonBytes)
}

// CreateAPI creates API for GetSamplesAll.
func CreateAPI() GetSamplesAllAPI {
	if api == nil {
		return GetSamplesAllAPIImpl{
			Interactor: core.CreateDefaultGetSamples("Kaka", "Ronaldo"),
		}
	}
	return api
}
