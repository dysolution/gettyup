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
func (r Release) Get() sdk.Release { return r.Unmarshal(get(r.path())) }

// Create associates a new Release with the specified Submission Batch.
func (r Release) Create() sdk.Release {
	return r.Unmarshal(post(r.build(r.context), batchPath(r.context)+"/releases"))
}

// Delete destroys a specific Release.
func (r Release) Delete()      { _delete(r.path()) }
func (r Release) id() string   { return getRequiredValue(r.context, "release-id") }
func (r Release) path() string { return childPath("releases", r.context, r.id()) }

func (r Release) build(c *cli.Context) sdk.Release {
	return sdk.Release{
		SubmissionBatchID:    c.Int("submission-batch-id"),
		FileName:             c.String("file-name"),
		FilePath:             c.String("file-path"),
		ExternalFileLocation: c.String("external-file-location"),
		ReleaseType:          c.String("release-type"),
		ModelDateOfBirth:     c.String("model-date-of-birth"),
		ModelEthnicities:     c.StringSlice("model-ethnicities"),
		ModelGender:          c.String("model-gender"),
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
