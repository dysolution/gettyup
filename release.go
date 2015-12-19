package main

import (
	"github.com/codegangsta/cli"
	sdk "github.com/dysolution/espsdk"
)

type Release struct{ context *cli.Context }

func (r Release) Index()       { get(childPath("releases", r.context, "")) }
func (r Release) Get()         { get(r.path()) }
func (r Release) Create()      { post(r.build(r.context), batchPath(r.context)+"/releases") }
func (r Release) Delete()      { _delete(r.path()) }
func (r Release) id() string   { return getRequiredValue(r.context, "release-id") }
func (r Release) path() string { return childPath("releases", r.context, r.id()) }
func (release Release) build(c *cli.Context) sdk.Release {
	return sdk.Release{
		SubmissionBatchId:    c.String("submission-batch-id"),
		FileName:             c.String("file-name"),
		FilePath:             c.String("file-path"),
		ExternalFileLocation: c.String("external-file-location"),
		ReleaseType:          c.String("release-type"),
		ModelDateOfBirth:     c.String("model-date-of-birth"),
		ModelEthnicities:     c.StringSlice("model-ethnicities"),
		ModelGender:          c.String("model-gender"),
	}
}
