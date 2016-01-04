package main

import (
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
func (b Batch) Get() sdk.DeserializedObject { return client.Get(b.path()) }

// Create adds a new Submission Batch.
func (b Batch) Create() sdk.DeserializedObject {
	data := b.build()
	return client.Create(data.Path(), data)
}

// Update changes fields for an existing Submission Batch.
func (b Batch) Update() sdk.DeserializedObject {
	data := sdk.BatchUpdate{b.build()}
	return client.Update(data.Path(), data)
}

// Delete destroys a specific Submission Batch.
func (b Batch) Delete() { client.Delete(b.path()) }

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
