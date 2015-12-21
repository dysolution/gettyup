package main

import (
	"strings"

	"github.com/codegangsta/cli"
	sdk "github.com/dysolution/espsdk"
)

// A ReleaseList contains zero or more Releases.
type ReleaseList []Release

// Unmarshal attempts to deserialize the provided JSON payload into a slice of Release objects.
func (rl ReleaseList) Unmarshal(payload []byte) sdk.ReleaseList {
	return sdk.ReleaseList{}.Unmarshal(payload)
}

var releaseTypes = string(strings.Join(sdk.Release{}.ValidTypes(), " OR "))

// A Release wraps the verbs provided by the ESP API for Releases,
// legal agreements for property owners or models to be associated
// with Submission Batches.
type Release struct{ context *cli.Context }

// Index requests a list of all Releases associated with the specified
// Submission Batch.
func (r Release) Index() sdk.ReleaseList {
	return ReleaseList{}.Unmarshal(get(childPath("releases", r.context, "")))
}

// Get requests the metadata for a specific Release.
//func (r Release) Get() sdk.Release { return r.Unmarshal(r.get()) }
func (r Release) Get() sdk.Release {
	return release(r.id()).Get(&client, getBatchID(r.context))
}

// Create associates a new Release with the specified Submission Batch.
func (r Release) Create() sdk.Release {
	return sdk.Release{}.Create(&client, getBatchID(r.context), r.build())
}

// Delete destroys a specific Release.
func (r Release) Delete() { release(r.id()).Delete(&client, getBatchID(r.context)) }

//func (r Release) id() string   { return getRequiredValue(r.context, "release-id") }
func (r Release) id() int      { return getReleaseID(r.context) }
func (r Release) path() string { return sdk.ReleasePath(getBatchID(r.context), r.id()) }
func (r Release) get() []byte  { return get(r.path()) }
func (r Release) post() []byte { return post(r.build(), batchPath(r.context)+"/releases") }

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
