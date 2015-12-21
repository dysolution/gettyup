package main

import (
	"github.com/codegangsta/cli"
	sdk "github.com/dysolution/espsdk"
)

// The TermService uses the SDK to deserialize responses from endpoints that
// provide TermLists.
type TermService struct{ context *cli.Context }

// Unmarshal attempts to deserialize the provided JSON payload into a SubmissionBatch object.
func (m TermService) Unmarshal(payload []byte) sdk.TermList {
	return sdk.TermList{}.Unmarshal(payload)
}

// GetNumberOfPeople lists all possible values for Number of People.
func (m TermService) GetNumberOfPeople(context *cli.Context) sdk.TermList {
	return sdk.TermList{}.GetNumberOfPeople(&client)
}

// GetExpressions lists all possible facial expression values.
func (m TermService) GetExpressions(context *cli.Context) sdk.TermList {
	return sdk.TermList{}.GetExpressions(&client)
}

// GetCompositions lists all possible composition values.
func (m TermService) GetCompositions(context *cli.Context) sdk.TermList {
	return sdk.TermList{}.GetCompositions(&client)
}
