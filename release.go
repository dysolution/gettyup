package main

import (
	"strings"

	"github.com/codegangsta/cli"
	sdk "github.com/dysolution/espsdk"
)

var releaseTypes = string(strings.Join(sdk.Release{}.ValidTypes(), " OR "))

// A Release wraps the verbs provided by the ESP API for Releases,
// legal agreements for property owners or models to be associated
// with Submission Batches.
type Release struct{ context *cli.Context }

// Index requests a list of all Releases associated with the specified
// Submission Batch.
func (r Release) Index() sdk.ReleaseList {
	return sdk.Release{}.Index(&client, getBatchID(r.context))
}

// Create associates a new Release with the specified Submission Batch.
func (r Release) Create() sdk.Createable {
	data := r.build()
	return client.Create(data.Path(), data)
}

// Delete destroys a specific Release.
func (r Release) Delete() { client.Delete(r.path()) }

// Get requests the metadata for a specific Release.
func (r Release) Get() sdk.DeserializedObject { return client.Get(r.path()) }

func (r Release) id() int { return getRequiredID(r.context, "release-id") }

func (r Release) path() string {
	obj := sdk.Release{
		ID:                r.id(),
		SubmissionBatchID: getBatchID(r.context),
	}
	return obj.Path()
}

func (r Release) build() sdk.Release {
	return sdk.Release{
		ExternalFileLocation: r.context.String("external-file-location"),
		FileName:             r.context.String("file-name"),
		FilePath:             r.context.String("file-path"),
		MimeType:             r.context.String("mime-type"),
		ModelDateOfBirth:     r.context.String("model-date-of-birth"),
		StorageURL:           r.context.String("storage-url"),
		ModelEthnicities:     r.context.StringSlice("model-ethnicities"),
		ModelGender:          r.context.String("model-gender"),
		ReleaseType:          r.context.String("release-type"),
		SubmissionBatchID:    r.context.Int("submission-batch-id"),
	}
}

// Unmarshal attempts to deserialize the provided JSON payload into a SubmissionBatch object.
func (r Release) Unmarshal(payload []byte) sdk.Release {
	return sdk.Release{}.Unmarshal(payload)
}
