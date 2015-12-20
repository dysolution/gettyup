package main

import (
	"encoding/json"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	sdk "github.com/dysolution/espsdk"
)

// A Release wraps the verbs provided by the ESP API for Releases,
// legal agreements for property owners or models to be associated
// with Submission Batches.
type Release struct{ context *cli.Context }

// Index requests a list of all Releases associated with the specified
// Submission Batch.
func (r Release) Index() { get(childPath("releases", r.context, "")) }

// Get requests the metadata for a specific Release.
func (r Release) Get() sdk.Release { return r.Unmarshal(r.get()) }

// Create associates a new Release with the specified Submission Batch.
func (r Release) Create() sdk.Release { return r.Unmarshal(r.post()) }

// Delete destroys a specific Release.
func (r Release) Delete() { _delete(r.path()) }

func (r Release) id() string   { return getRequiredValue(r.context, "release-id") }
func (r Release) path() string { return childPath("releases", r.context, r.id()) }
func (r Release) get() []byte  { return get(r.path()) }
func (r Release) post() []byte { return post(r.build(), batchPath(r.context)+"/releases") }

func (r Release) build() sdk.Release {
	return sdk.Release{
		ExternalFileLocation: r.context.String("external-file-location"),
		FileName:             r.context.String("file-name"),
		FilePath:             r.context.String("file-path"),
		ModelDateOfBirth:     r.context.String("model-date-of-birth"),
		ModelEthnicities:     r.context.StringSlice("model-ethnicities"),
		ModelGender:          r.context.String("model-gender"),
		ReleaseType:          r.context.String("release-type"),
		SubmissionBatchID:    r.context.Int("submission-batch-id"),
	}
}

// Unmarshal attempts to deserialize the provided JSON payload into a
// Release object as defined by the SDK.
func (r Release) Unmarshal(payload []byte) sdk.Release {
	var release sdk.Release
	if err := json.Unmarshal(payload, &release); err != nil {
		log.Fatal(err)
	}
	return release
}
