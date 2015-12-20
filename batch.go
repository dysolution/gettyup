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
func (b Batch) Index() sdk.BatchListContainer { return sdk.BatchListContainer{}.Unmarshal(get(Batches)) }

// Get requests the metadata for a specific Submission Batch.
func (b Batch) Get() sdk.Batch { return sdk.Batch{ID: b.id()}.Get(&client) }

// Create adds a new Submission Batch.
func (b Batch) Create() sdk.Batch { return b.Unmarshal(b.post()) }

// Update changes fields for an existing Submission Batch.
func (b Batch) Update() sdk.Batch { return b.Unmarshal(b.put()) }

// Delete destroys a specific Submission Batch.
func (b Batch) Delete() { _delete(b.path()) }

// Unmarshal attempts to deserialize the provided JSON payload into a SubmissionBatch object.
func (b Batch) Unmarshal(payload []byte) sdk.Batch {
	return sdk.Batch{}.Unmarshal(payload)
}

func (b Batch) id() int      { return getBatchID(b.context) }
func (b Batch) path() string { return sdk.BatchPath(&sdk.Batch{ID: b.id()}) }
func (b Batch) get() []byte  { return get(b.path()) }
func (b Batch) post() []byte { return post(b.build(), Batches) }
func (b Batch) put() []byte  { return put(b.buildUpdate(), b.path()) }

func (b Batch) build() sdk.Batch {
	return sdk.Batch{
		BriefID:               b.context.String("brief-id"),
		EventID:               b.context.String("event-id"),
		Note:                  b.context.String("note"),
		AssignmentID:          b.context.String("assignment-id"),
		SaveExtractedMetadata: b.context.Bool("save-extracted-metadata"),
		SubmissionName:        b.context.String("submission-name"),
		SubmissionType:        b.context.String("submission-type"),
	}
}

func (b Batch) buildUpdate() sdk.BatchUpdate {
	return sdk.BatchUpdate{b.build()}
}
