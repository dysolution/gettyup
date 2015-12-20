package main

import (
	"encoding/json"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	sdk "github.com/dysolution/espsdk"
)

// A Batch wraps the verbs provided by the ESP API for Submission Batches.
type Batch struct{ context *cli.Context }

// Index requests a list of all Submission Batches belonging to the user.
func (b Batch) Index() { get(Batches) }

// Get requests the metadata for a specific Submission Batch.
func (b Batch) Get() sdk.SubmissionBatch { return b.Unmarshal(get(b.path())) }

// Create adds a new Submission Batch.
func (b Batch) Create() sdk.SubmissionBatch { return b.Unmarshal(post(b.build(b.context), Batches)) }

// Update changes fields for an existing Submission Batch.
func (b Batch) Update() sdk.SubmissionBatch { return b.Unmarshal(put(b.buildUpdate(), b.path())) }

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

func (b Batch) path() string { return Batches + "/" + getBatchID(b.context) }
func (b Batch) id() string   { return getRequiredValue(b.context, "submission-batch-id") }

func (b Batch) build(c *cli.Context) sdk.SubmissionBatch {
	return sdk.SubmissionBatch{
		SubmissionName:        c.String("submission-name"),
		SubmissionType:        c.String("submission-type"),
		Note:                  c.String("note"),
		AssignmentID:          c.String("assignment-id"),
		BriefID:               c.String("brief-id"),
		EventID:               c.String("event-id"),
		SaveExtractedMetadata: c.Bool("save-extracted-metadata"),
	}
}

func (b Batch) buildUpdate() sdk.SubmissionBatchUpdate {
	return sdk.SubmissionBatchUpdate{
		sdk.SubmissionBatchChanges{
			SubmissionName: b.context.String("submission-name"),
			Note:           b.context.String("note"),
		},
	}
}
