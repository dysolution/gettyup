package main

import (
	sdk "github.com/dysolution/espsdk"
)

// These constants represent the relative paths for various ESP API endpoints.
const (
	Batches            string = sdk.Batches
	ControlledValues   string = sdk.ControlledValues
	Keywords           string = sdk.Keywords
	Personalities      string = sdk.Personalities
	TranscoderMappings string = sdk.TranscoderMappings
	Compositions       string = sdk.Compositions
	Expressions        string = sdk.Expressions
	NumberOfPeople     string = sdk.NumberOfPeople
)
