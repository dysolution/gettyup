package main

import (
	"runtime"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/dysolution/espsdk"
)

var batchTypes = string(strings.Join(espsdk.Batch{}.ValidTypes(), " OR "))

// A Batch wraps the verbs provided by the ESP API for Submission Batches.
type Batch struct{ context *cli.Context }

// Index requests a list of all Submission Batches belonging to the user.
func (b Batch) Index() espsdk.BatchList {
	return espsdk.Batch{}.Index(client)
}

// Get requests the metadata for a specific Submission Batch.
func (b Batch) Get() *espsdk.Batch {
	desc := "Batch.Get"
	data := b.build()
	var batch *espsdk.Batch

	result, err := client.Get(data)
	if err != nil {
		result.Log().Error(desc)
		return batch
	}
	if result.StatusCode == 404 {
		result.Log().Error(desc)
		return batch
	}
	result.Log().Info(desc)
	batch, err = espsdk.Batch{}.Unmarshal(result.Payload)
	if err != nil {
		result.Log().Errorf("%s: %v", desc, err)
		return batch
	}
	return batch

}

// Create adds a new Submission Batch.
func (b Batch) Create() *espsdk.Batch {
	myPC, _, _, _ := runtime.Caller(0)
	desc := runtime.FuncForPC(myPC).Name() + ": "
	data := b.build()
	var batch *espsdk.Batch

	result, err := client.Create(data)
	if err != nil {
		result.Log().Errorf("%s: %v", desc, err)
		return batch
	}

	switch result.StatusCode {
	case 201:
		result.Log().Info(desc + "created")
		batch, err = espsdk.Batch{}.Unmarshal(result.Payload)
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
func (b Batch) Update() *espsdk.Batch {
	myPC, _, _, _ := runtime.Caller(0)
	desc := runtime.FuncForPC(myPC).Name() + ": "
	data := b.build()
	var batch *espsdk.Batch

	result, err := client.Update(data)
	if err != nil {
		result.Log().Errorf("%s: %v", desc, err)
		return batch
	}

	switch result.StatusCode {
	case 200:
		result.Log().Info(desc + "updated")
		batch, err = espsdk.Batch{}.Unmarshal(result.Payload)
		if err != nil {
			result.Log().Errorf("%s: %v", desc, err)
		}
	case 401:
		result.Log().Error(desc + "unauthorized")
	case 403:
		result.Log().Error(desc + "forbidden")
	case 404:
		result.Log().Error(desc + "submission batch not found")
	case 422:
		result.Log().Error(desc + "unprocessable: bad params or closed batch")
	}
	return batch
}

// Delete destroys a specific Submission Batch.
func (b Batch) Delete() *espsdk.Batch {
	myPC, _, _, _ := runtime.Caller(0)
	desc := runtime.FuncForPC(myPC).Name() + ": "
	data := b.build()
	var batch *espsdk.Batch

	result, err := client.Delete(data)
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
		batch, err = espsdk.Batch{}.Unmarshal(result.Payload)
		if err != nil {
			result.Log().Errorf("%s: %v", desc, err)
		}
	}
	return batch
}

// Last returns the newest Submission Batch.
func (b Batch) Last() espsdk.Batch {
	return espsdk.Batch{}.Index(client).Last()
}

// Unmarshal attempts to deserialize the provided JSON payload into a SubmissionBatch object.
func (b Batch) Unmarshal(payload []byte) *espsdk.Batch {
	batch, err := espsdk.Batch{}.Unmarshal(payload)
	if err != nil {
		return &espsdk.Batch{}
	}
	return batch
}

func (b Batch) id() string   { return getBatchID(b.context) }
func (b Batch) path() string { return espsdk.Batch{ID: b.id()}.Path() }

func (b Batch) build() espsdk.Batch {
	return espsdk.Batch{
		BriefID:               b.context.String("brief-id"),
		EventID:               b.context.String("event-id"),
		ID:                    b.context.String("submission-batch-id"),
		Note:                  b.context.String("note"),
		AssignmentID:          b.context.String("assignment-id"),
		SaveExtractedMetadata: b.context.Bool("save-extracted-metadata"),
		SubmissionName:        b.context.String("submission-name"),
		SubmissionType:        b.context.String("submission-type"),
	}
}
