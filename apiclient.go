package main

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	api "github.com/dysolution/espapi"
)

var client api.Client
var uploadBucket string

var batchTypes = api.BatchTypes()
var releaseTypes = api.ReleaseTypes()

func getClient(key, secret, username, password string) api.Client {
	return api.Client{
		api.Credentials{
			ApiKey:      key,
			ApiSecret:   secret,
			EspUsername: username,
			EspPassword: password,
		},
		uploadBucket,
	}
}

func BuildBatch(c *cli.Context) api.SubmissionBatch {
	return api.SubmissionBatch{
		SubmissionName:        c.String("submission-name"),
		SubmissionType:        c.String("submission-type"),
		Note:                  c.String("note"),
		AssignmentId:          c.String("assignment-id"),
		BriefId:               c.String("brief-id"),
		EventId:               c.String("event-id"),
		SaveExtractedMetadata: c.Bool("save-extracted-metadata"),
	}
}

func BuildRelease(c *cli.Context) api.Release {
	return api.Release{
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

func BuildContribution(c *cli.Context) api.Contribution {
	return api.Contribution{
		FileName:             c.String("file-name"),
		FilePath:             c.String("file-path"),
		SubmittedToReviewAt:  c.String("submitted-to-review-at"),
		UploadBucket:         c.String("upload-bucket"),
		ExternalFileLocation: c.String("external-file-location"),
		UploadId:             c.String("upload-id"),
		MimeType:             c.String("mime-type"),
	}
}

func Token(context *cli.Context, client api.Client) api.Token {
	return client.GetToken()
}

func CreateBatch(context *cli.Context, client api.Client) {
	path := "/submission/v1/submission_batches"
	batch, err := BuildBatch(context).Marshal()
	if err != nil {
		log.Errorf("error creating batch")
	}
	response, err := client.Post(batch, Token(context, client), path)
	if err != nil {
		log.Errorf("error POSTing batch")
	}
	log.Infof("%s\n", response)
}

func CreateRelease(context *cli.Context, client api.Client) {
	batch_id := context.String("submission-batch-id")
	if len(batch_id) < 1 {
		log.Fatalf("--submission-batch-id must be set")
	}
	path := fmt.Sprintf("/submission/v1/submission_batches/%s/releases", batch_id)
	release, err := BuildRelease(context).Marshal()
	if err != nil {
		log.Errorf("error creating release")
	}
	response, err := client.Post(release, Token(context, client), path)
	if err != nil {
		log.Errorf("error POSTing batch")
	}
	log.Infof("%s\n", response)
}

func CreateContribution(context *cli.Context, client api.Client) {
	release, err := BuildContribution(context).Marshal()
	if err != nil {
		log.Errorf("error creating contribution")
	}
	client.PostContribution(release)
}
