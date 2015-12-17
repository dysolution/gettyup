package main

import (
	"github.com/codegangsta/cli"
	models "github.com/dysolution/espapi"
)

var client models.Client

var batchTypes = models.BatchTypes()
var releaseTypes = models.ReleaseTypes()

var token models.Token

type Batch struct{ context *cli.Context }
type Release struct{ context *cli.Context }
type Contribution struct{ context *cli.Context }

func (b Batch) Index()       { get(Batches) }
func (b Batch) Get()         { get(b.path()) }
func (b Batch) Create()      { post(b.build(b.context), Batches) }
func (b Batch) Update()      { put(b.buildUpdate(), b.path()) }
func (b Batch) Delete()      { _delete(b.path()) }
func (b Batch) path() string { return Batches + "/" + getBatchID(b.context) }
func (b Batch) id() string   { return getRequiredValue(b.context, "submission-batch-id") }

func (b Batch) build(c *cli.Context) models.SubmissionBatch {
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

func (b Batch) buildUpdate() models.SubmissionBatchUpdate {
	return models.SubmissionBatchUpdate{
		models.SubmissionBatchChanges{
			SubmissionName: b.context.String("submission-name"),
			Note:           b.context.String("note"),
		},
	}
}

func (r Release) Index()     { get(childPath("releases", r.context, "")) }
func (r Release) Get()       { get(childPath("releases", r.context, r.id())) }
func (r Release) Create()    { post(r.build(r.context), batchPath(r.context)+"/releases") }
func (r Release) Delete()    { _delete(childPath("releases", r.context, r.id())) }
func (r Release) id() string { return getRequiredValue(r.context, "release-id") }

func (release Release) build(c *cli.Context) models.Release {
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

func (c Contribution) Index()     { get(childPath("contributions", c.context, "")) }
func (c Contribution) Get()       { get(childPath("contributions", c.context, c.id())) }
func (c Contribution) Create()    { post(c.build(c.context), batchPath(c.context)+"/contributions") }
func (c Contribution) Delete()    { _delete(childPath("contributions", c.context, c.id())) }
func (c Contribution) id() string { return getRequiredValue(c.context, "contribution-id") }

//func (c Contribution) Update() { put(buildContributionUpdate(c.context), c.path()) }

func (contribution Contribution) build(c *cli.Context) models.Contribution {
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

// Token is a memoizing wrapper for the API's token-providing function.
func Token() models.Token {
	if token != "" {
		return token
	} else {
		token = client.GetToken()
		return token
	}
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

type Serializable interface {
	Marshal() ([]byte, error)
}
