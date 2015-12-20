package main

import (
	"encoding/json"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	sdk "github.com/dysolution/espsdk"
)

var batchTypes = string(strings.Join(sdk.SubmissionBatch{}.ValidTypes(), " OR "))

// A Batch wraps the verbs provided by the ESP API for Submission Batches.
type Batch struct{ context *cli.Context }

// Index requests a list of all Submission Batches belonging to the user.
func (b Batch) Index() { get(Batches) }

// Get requests the metadata for a specific Submission Batch.
func (b Batch) Get() sdk.SubmissionBatch { return b.Unmarshal(b.get()) }

// Create adds a new Submission Batch.
func (b Batch) Create() sdk.SubmissionBatch { return b.Unmarshal(b.post()) }

// Update changes fields for an existing Submission Batch.
func (b Batch) Update() sdk.SubmissionBatch { return b.Unmarshal(b.put()) }

// Delete destroys a specific Submission Batch.
func (b Batch) Delete() { _delete(b.path()) }

// Unmarshal attempts to deserialize the provided JSON payload into a SubmissionBatch object.
func (b Batch) Unmarshal(payload []byte) sdk.SubmissionBatch {
	var batch sdk.SubmissionBatch
	if err := json.Unmarshal(payload, &batch); err != nil {
		log.Fatal(err)
	}
	return batch
}

func (b Batch) path() string { return Batches + "/" + b.id() }
func (b Batch) id() string   { return getRequiredValue(b.context, "submission-batch-id") }
func (b Batch) get() []byte  { return get(b.path()) }
func (b Batch) post() []byte { return post(b.build(), Batches) }
func (b Batch) put() []byte  { return put(b.buildUpdate(), b.path()) }

func (b Batch) build() sdk.SubmissionBatch {
	return sdk.SubmissionBatch{
		BriefID:               b.context.String("brief-id"),
		EventID:               b.context.String("event-id"),
		Note:                  b.context.String("note"),
		AssignmentID:          b.context.String("assignment-id"),
		SaveExtractedMetadata: b.context.Bool("save-extracted-metadata"),
		SubmissionName:        b.context.String("submission-name"),
		SubmissionType:        b.context.String("submission-type"),
	}
}

func (b Batch) buildUpdate() sdk.SubmissionBatchUpdate {
	return sdk.SubmissionBatchUpdate{b.build()}
}
