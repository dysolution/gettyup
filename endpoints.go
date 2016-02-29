package main

import (
	"github.com/dysolution/espsdk"
)

// These constants represent the relative paths for various ESP API endpoints.
const (
	Batches            string = espsdk.BatchesEndpoint
	ControlledValues   string = espsdk.ControlledValuesEndpoint
	Keywords           string = espsdk.KeywordsEndpoint
	Personalities      string = espsdk.PersonalitiesEndpoint
	TranscoderMappings string = espsdk.TranscoderMappingsEndpoint
	Compositions       string = espsdk.CompositionsEndpoint
	Expressions        string = espsdk.ExpressionsEndpoint
	NumberOfPeople     string = espsdk.NumberOfPeopleEndpoint
)
