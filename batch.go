package main

import (
	"runtime"
	"strings"

	"github.com/codegangsta/cli"
	sdk "github.com/dysolution/espsdk"
)

var batchTypes = string(strings.Join(sdk.Batch{}.ValidTypes(), " OR "))

// A Batch wraps the verbs provided by the ESP API for Submission Batches.
type Batch struct{ context *cli.Context }

// Index requests a list of all Submission Batches belonging to the user.
func (b Batch) Index() *sdk.DeserializedObject {
	data := b.build()
	return client.Index(data.Path())
}

// Get requests the metadata for a specific Submission Batch.
func (b Batch) Get() *sdk.Batch {
	desc := "Batch.Get"
	data := b.build()
	var batch *sdk.Batch

	result, err := client.VerboseGet(data)
	if err != nil {
		result.Log().Error(desc)
		return batch
	}
	if result.StatusCode == 404 {
		result.Log().Error(desc)
		return batch
	}
	result.Log().Info(desc)
	batch, err = sdk.Batch{}.Unmarshal(result.Payload)
	if err != nil {
		result.Log().Errorf("%s: %v", desc, err)
		return batch
	}
	return batch

}

// Create adds a new Submission Batch.
func (b Batch) Create() *sdk.Batch {
	myPC, _, _, _ := runtime.Caller(0)
	desc := runtime.FuncForPC(myPC).Name() + ": "
	data := b.build()
	var batch *sdk.Batch

	result, err := client.Create(data)
	if err != nil {
		result.Log().Errorf("%s: %v", desc, err)
		return batch
	}

	switch result.StatusCode {
	case 201:
		result.Log().Info(desc + "created")
		batch, err = sdk.Batch{}.Unmarshal(result.Payload)
		if err != nil {
			result.Log().Errorf("%s: %v", desc, err)
		}
	case 401:
		result.Log().Error(desc + "unauthorized")
	case 403:
		result.Log().Error(desc + "forbidden")
	case 422:
		result.Log().Error(desc + "invalid batch data provided")
	}
	return batch
}

// Update changes fields for an existing Submission Batch.
func (b Batch) Update() sdk.DeserializedObject {
	data := sdk.BatchUpdate{Batch: b.build()}
	return client.Update(data)
}

// Delete destroys a specific Submission Batch.
func (b Batch) Delete() *sdk.Batch {
	myPC, _, _, _ := runtime.Caller(0)
	desc := runtime.FuncForPC(myPC).Name() + ": "
	data := b.build()
	var batch *sdk.Batch

	result, err := client.VerboseDelete(data)
	if err != nil {
		result.Log().Errorf("%s: %v", desc, err)
		return batch
	}

	switch result.StatusCode {
	case 204:
		result.Log().Info(desc + "deleted")
	case 401:
		result.Log().Error(desc + "unauthorized")
	case 403:
		result.Log().Error(desc + "forbidden")
	case 404:
		result.Log().Error(desc + "submission batch not found")
	case 422:
		result.Log().Error(desc + "batch is protected from deletion")
	}
	// successful deletion usually returns a 204 without a payload/body
	if len(result.Payload) > 0 {
		batch, err = sdk.Batch{}.Unmarshal(result.Payload)
		if err != nil {
			result.Log().Errorf("%s: %v", desc, err)
		}
	}
	return batch
}

// Last returns the newest Submission Batch.
func (b Batch) Last() sdk.Batch { return client.Index(sdk.Batches).Last() }

// Unmarshal attempts to deserialize the provided JSON payload into a SubmissionBatch object.
func (b Batch) Unmarshal(payload []byte) sdk.DeserializedObject {
	return sdk.Unmarshal(payload)
}

func (b Batch) id() int      { return getBatchID(b.context) }
func (b Batch) path() string { return sdk.Batch{ID: b.id()}.Path() }

func (b Batch) build() sdk.Batch {
	return sdk.Batch{
		BriefID:               b.context.String("brief-id"),
		EventID:               b.context.String("event-id"),
		ID:                    b.context.Int("submission-batch-id"),
		Note:                  b.context.String("note"),
		AssignmentID:          b.context.String("assignment-id"),
		SaveExtractedMetadata: b.context.Bool("save-extracted-metadata"),
		SubmissionName:        b.context.String("submission-name"),
		SubmissionType:        b.context.String("submission-type"),
	}
}
