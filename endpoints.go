package main

import (
	"github.com/dysolution/espsdk"
)

// The relative paths for various ESP API endpoints.
var (
	Batches            = espsdk.Endpoints.Batches
	ControlledValues   = espsdk.Endpoints.ControlledValues
	Keywords           = espsdk.Endpoints.Keywords
	Personalities      = espsdk.Endpoints.Personalities
	TranscoderMappings = espsdk.Endpoints.TranscoderMappings
	Compositions       = espsdk.Endpoints.Compositions
	Expressions        = espsdk.Endpoints.Expressions
	NumberOfPeople     = espsdk.Endpoints.NumberOfPeople
)
