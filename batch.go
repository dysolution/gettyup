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
func (b Batch) Index() sdk.BatchListContainer { return sdk.Batch{}.Index(&client) }

// Get requests the metadata for a specific Submission Batch.
func (b Batch) Get() sdk.Batch { return batch(b.id()).Get(&client) }

// Create adds a new Submission Batch.
func (b Batch) Create() sdk.Batch { return sdk.Batch{}.Create(&client, b.build()) }

// Update changes fields for an existing Submission Batch.
func (b Batch) Update() sdk.Batch { return batch(b.id()).Update(&client, b.buildUpdate()) }

// Delete destroys a specific Submission Batch.
func (b Batch) Delete() { batch(b.id()).Delete(&client) }

// Unmarshal attempts to deserialize the provided JSON payload into a SubmissionBatch object.
func (b Batch) Unmarshal(payload []byte) sdk.Batch {
	return sdk.Batch{}.Unmarshal(payload)
}

func (b Batch) id() int { return getBatchID(b.context) }

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
