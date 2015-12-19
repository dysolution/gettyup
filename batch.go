package main

import (
	"encoding/json"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	sdk "github.com/dysolution/espsdk"
)

type Batch struct{ context *cli.Context }

func (b Batch) Index()                   { get(Batches) }
func (b Batch) Get() sdk.SubmissionBatch { return b.Unmarshal(get(b.path())) }
func (b Batch) Create()                  { post(b.build(b.context), Batches) }
func (b Batch) Update()                  { put(b.buildUpdate(), b.path()) }
func (b Batch) Delete()                  { _delete(b.path()) }
func (b Batch) path() string             { return Batches + "/" + getBatchID(b.context) }
func (b Batch) id() string               { return getRequiredValue(b.context, "submission-batch-id") }

func (b Batch) Unmarshal(payload []byte) sdk.SubmissionBatch {
	var batch sdk.SubmissionBatch
	if err := json.Unmarshal(payload, &batch); err != nil {
		log.Fatal(err)
	}
	return batch
}

func (b Batch) build(c *cli.Context) sdk.SubmissionBatch {
	return sdk.SubmissionBatch{
		SubmissionName:        c.String("submission-name"),
		SubmissionType:        c.String("submission-type"),
		Note:                  c.String("note"),
		AssignmentId:          c.String("assignment-id"),
		BriefId:               c.String("brief-id"),
		EventId:               c.String("event-id"),
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
