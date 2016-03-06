package main

import (
	"runtime"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/dysolution/espsdk"
)

var releaseTypes = string(strings.Join(espsdk.Release{}.ValidTypes(), " OR "))

// A Release wraps the verbs provided by the ESP API for Releases,
// legal agreements for property owners or models to be associated
// with Submission Batches.
type Release struct{ context *cli.Context }

// Index requests a list of all Releases associated with the specified
// Submission Batch.
func (r Release) Index() espsdk.ReleaseList {
	return espsdk.Release{}.Index(client, getBatchID(r.context))
}

// Create associates a new Release with the specified Submission Batch.
func (r Release) Create() *espsdk.Release {
	desc := "Release.Create: "
	data := r.build()
	var release *espsdk.Release

	result, err := client.Create(data)
	if err != nil {
		result.Log().Errorf("%s: %v", desc, err)
		return release
	}

	switch result.StatusCode {
	case 201:
		result.Log().Info(desc + "created")
		release, err = espsdk.Release{}.Unmarshal(result.Payload)
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
func (r Release) Delete() *espsdk.Release {
	myPC, _, _, _ := runtime.Caller(0)
	desc := runtime.FuncForPC(myPC).Name() + ": "
	data := r.build()
	var release *espsdk.Release

	result, err := client.Delete(data)
	if err != nil {
		result.Log().Errorf("%s: %v", desc, err)
		return release
	}

	switch result.StatusCode {
	case 204:
		result.Log().Info(desc + "deleted")
	case 401:
		result.Log().Error(desc + "unauthorized")
	case 403:
		result.Log().Error(desc + "forbidden")
	case 404:
		result.Log().Error(desc + "submission batch or release not found")
	case 422:
		result.Log().Error(desc + "unprocessable: release associated with already-submitted contribution")
	}
	// successful deletion usually returns a 204 without a payload/body
	if len(result.Payload) > 0 {
		release, err = espsdk.Release{}.Unmarshal(result.Payload)
		if err != nil {
			result.Log().Errorf("%s: %v", desc, err)
		}
	}
	return release
}

// Get requests the metadata for a specific Release.
func (r Release) Get() *espsdk.Release {
	desc := "Release.Get"
	data := r.build()
	var release *espsdk.Release

	result, err := client.Get(data)
	if err != nil {
		result.Log().Error(desc)
		return release
	}
	if result.StatusCode == 404 {
		result.Log().Error(desc)
		return release
	}
	result.Log().Info(desc)
	release, err = espsdk.Release{}.Unmarshal(result.Payload)
	if err != nil {
		result.Log().Errorf("%s: %v", desc, err)
		return release
	}
	return release

}

func (r Release) id() string { return getRequiredID(r.context, "release-id") }

func (r Release) path() string {
	obj := espsdk.Release{
		ID:                r.id(),
		SubmissionBatchID: getBatchID(r.context),
	}
	return obj.Path()
}

func (r Release) build() espsdk.Release {
	return espsdk.Release{
		ExternalFileLocation: r.context.String("external-file-location"),
		FileName:             r.context.String("file-name"),
		FilePath:             r.context.String("file-path"),
		ID:                   r.context.String("release-id"),
		MimeType:             r.context.String("mime-type"),
		ModelDateOfBirth:     r.context.String("model-date-of-birth"),
		ModelEthnicities:     r.context.StringSlice("model-ethnicities"),
		ModelGender:          r.context.String("model-gender"),
		ReleaseType:          r.context.String("release-type"),
		StorageURL:           r.context.String("storage-url"),
		SubmissionBatchID:    r.context.String("submission-batch-id"),
	}
}

// Unmarshal attempts to deserialize the provided JSON payload into a SubmissionBatch object.
func (r Release) Unmarshal(payload []byte) (*espsdk.Release, error) {
	return espsdk.Release{}.Unmarshal(payload)
}

func (r Release) registerCmds() {
	app.Commands = append(app.Commands, cli.Command{
		Name:  "release",
		Usage: "work with Releases",
		Subcommands: []cli.Command{
			{
				Name:   "create",
				Usage:  "create a new Release within a Submission Batch",
				Action: func(c *cli.Context) { pp(Release{c}.Create()) },
				Flags: []cli.Flag{
					cli.StringFlag{Name: "external-file-location"},
					cli.StringFlag{Name: "file-name"},
					cli.StringFlag{Name: "file-path"},
					cli.StringFlag{Name: "mime-type"},
					cli.StringFlag{Name: "model-date-of-birth"},
					cli.StringFlag{Name: "model-gender"},
					cli.StringFlag{Name: "release-type", Usage: releaseTypes},
					cli.StringFlag{Name: "submission-batch-id, b"},
					cli.StringSliceFlag{Name: "model-ethnicities"},
				},
			},
			{
				Name:  "get",
				Usage: "get a specific Release",
				Action: func(c *cli.Context) {
					pp(Release{c}.Get())
				},
				Flags: []cli.Flag{
					cli.StringFlag{Name: "submission-batch-id, b"},
					cli.StringFlag{Name: "release-id, r"},
				},
			},
			{
				Name:  "index",
				Usage: "get all Releases for a Submission Batch",
				Action: func(c *cli.Context) {
					pp(Release{c}.Index())
				},
				Flags: []cli.Flag{
					cli.StringFlag{Name: "submission-batch-id, b"},
				},
			},
			{
				Name:   "delete",
				Usage:  "delete an existing Release",
				Action: func(c *cli.Context) { Release{c}.Delete() },
				Flags: []cli.Flag{
					cli.StringFlag{Name: "submission-batch-id, b"},
					cli.StringFlag{Name: "release-id, r"},
				},
			},
		},
	})
}
