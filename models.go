package main

import (
	"github.com/codegangsta/cli"
	models "github.com/dysolution/espapi"
)

var client models.Client

var batchTypes = models.BatchTypes()
var releaseTypes = models.ReleaseTypes()

// Token is a wrapper for the API's token-providing function.
func Token() models.Token {
	return client.GetToken()
}

func getClient(key, secret, username, password string) models.Client {
	return models.Client{
		models.Credentials{
			APIKey:      key,
			APISecret:   secret,
			ESPUsername: username,
			ESPPassword: password,
		},
		uploadBucket,
	}
}

// buildBatch takes a Context of CLI-provided values
// and returns a SubmissionBatch as defined by the models.
func buildBatch(c *cli.Context) models.SubmissionBatch {
	return models.SubmissionBatch{
		SubmissionName:        c.String("submission-name"),
		SubmissionType:        c.String("submission-type"),
		Note:                  c.String("note"),
		AssignmentId:          c.String("assignment-id"),
		BriefId:               c.String("brief-id"),
		EventId:               c.String("event-id"),
		SaveExtractedMetadata: c.Bool("save-extracted-metadata"),
	}
}

type Serializable interface {
	Marshal() ([]byte, error)
}

func buildRelease(c *cli.Context) models.Release {
	return models.Release{
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

func buildContribution(c *cli.Context) models.Contribution {
	return models.Contribution{
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
