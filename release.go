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
func (r Release) Create() *sdk.Release {
	desc := "Release.Create: "
	data := r.build()
	var release *sdk.Release

	result, err := client.VerboseCreate(data)
	if err != nil {
		result.Log().Errorf("%s: %v", desc, err)
		return release
	}

	switch result.StatusCode {
	case 201:
		result.Log().Info(desc + "created")
		release, err = sdk.Release{}.Unmarshal(result.Payload)
		if err != nil {
			result.Log().Errorf("%s: %v", desc, err)
		}
	case 401:
		result.Log().Error(desc + "unauthorized")
	case 403:
		result.Log().Error(desc + "forbidden")
	case 404:
		result.Log().Error(desc + "submission batch not found")
	case 422:
		msg := desc + "the release failed to upload: "
		msg += data.ExternalFileLocation
		result.Log().Error(msg)
	}
	return release
}

// Delete destroys a specific Release.
func (r Release) Delete() sdk.DeserializedObject {
	return client.Delete(r.path())
}

// Get requests the metadata for a specific Release.
func (r Release) Get() *sdk.Release {
	desc := "Release.Get"
	data := r.build()
	var release *sdk.Release

	result, err := client.VerboseGet(data)
	if err != nil {
		result.Log().Error(desc)
		return release
	}
	if result.StatusCode == 404 {
		result.Log().Error(desc)
		return release
	}
	result.Log().Info(desc)
	release, err = sdk.Release{}.Unmarshal(result.Payload)
	if err != nil {
		result.Log().Errorf("%s: %v", desc, err)
		return release
	}
	return release

}

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
		ID:                   r.context.Int("release-id"),
		MimeType:             r.context.String("mime-type"),
		ModelDateOfBirth:     r.context.String("model-date-of-birth"),
		ModelEthnicities:     r.context.StringSlice("model-ethnicities"),
		ModelGender:          r.context.String("model-gender"),
		ReleaseType:          r.context.String("release-type"),
		StorageURL:           r.context.String("storage-url"),
		SubmissionBatchID:    r.context.Int("submission-batch-id"),
	}
}

// Unmarshal attempts to deserialize the provided JSON payload into a SubmissionBatch object.
func (r Release) Unmarshal(payload []byte) (*sdk.Release, error) {
	return sdk.Release{}.Unmarshal(payload)
}
