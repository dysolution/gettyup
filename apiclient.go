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

// BuildBatch takes a Context of CLI-provided values
// and returns a SubmissionBatch as defined by the api package.
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
		CameraShotDate:       c.String("camera-shot-date"),
		CollectionCode:       c.String("collection-code"),
		ContentProviderName:  c.String("content-provider-name"),
		ContentProviderTitle: c.String("content-provider-title"),
		CountryOfShoot:       c.String("country-of-shoot"),
		CreditLine:           c.String("credit-line"),
		ExternalFileLocation: c.String("external-file-location"),
		FileName:             c.String("file-name"),
		FilePath:             c.String("file-path"),
		Headline:             c.String("headline"),
		IptcCategory:         c.String("iptc-category"),
		MimeType:             c.String("mime-type"),
		ParentSource:         c.String("parent-source"),
		RecordedDate:         c.String("recorded-date"),
		RiskCategory:         c.String("risk-category"),
		ShotSpeed:            c.String("shot-speed"),
		SiteDestination:      c.StringSlice("site-destination"),
		Source:               c.String("source"),
		SubmittedToReviewAt:  c.String("submitted-to-review-at"),
		UploadBucket:         uploadBucket,
		UploadId:             c.String("upload-id"),
	}
}

func Token(context *cli.Context, client api.Client) api.Token {
	return client.GetToken()
}

func CreateBatch(context *cli.Context, client api.Client) {
	path := "/submission/v1/submission_batches"
	batch, err := BuildBatch(context).Marshal()
	if err != nil {
		log.Error(err)
	}

	response, err := client.Post(batch, Token(context, client), path)
	if err != nil {
		log.Error(err)
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
		log.Fatal(err)
	}
	response, err := client.Post(release, Token(context, client), path)
	if err != nil {
		log.Fatal(err)
	}
	log.Infof("%s\n", response)
}

func CreateContribution(context *cli.Context, client api.Client) {
	batch_id := context.String("submission-batch-id")
	if len(batch_id) < 1 {
		log.Fatalf("--submission-batch-id must be set")
	}
	path := fmt.Sprintf("/submission/v1/submission_batches/%s/contributions", batch_id)

	contribution, err := BuildContribution(context).Marshal()
	if err != nil {
		log.Error(err)
	}

	response, err := client.Post(contribution, Token(context, client), path)
	if err != nil {
		log.Fatal(err)
	}
	log.Infof("%s\n", response)
}
