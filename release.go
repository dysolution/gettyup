package main

import (
	"encoding/json"
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	sdk "github.com/dysolution/espsdk"
)

type Release struct{ context *cli.Context }

func (r Release) Index()           { get(childPath("releases", r.context, "")) }
func (r Release) Get() sdk.Release { return r.Unmarshal(get(r.path())) }
func (r Release) Create() sdk.Release {
	return r.Unmarshal(post(r.build(r.context), batchPath(r.context)+"/releases"))
}
func (r Release) Delete()      { _delete(r.path()) }
func (r Release) id() string   { return getRequiredValue(r.context, "release-id") }
func (r Release) path() string { return childPath("releases", r.context, r.id()) }

func (release Release) build(c *cli.Context) sdk.Release {
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

func (r Release) Unmarshal(payload []byte) sdk.Release {
	var release sdk.Release
	if err := json.Unmarshal(payload, &release); err != nil {
		log.Fatal(err)
	}
	return release
}

func (r Release) PrettyPrint() string {
	prettyOutput, err := r.Get().Marshal()
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%s\n", prettyOutput)
}
