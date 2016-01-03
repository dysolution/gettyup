package main

import (
	"github.com/codegangsta/cli"
	sdk "github.com/dysolution/espsdk"
)

// A ContributionList contains zero or more Contributions.
type ContributionList []Contribution

// Unmarshal attempts to deserialize the provided JSON payload into a slice of Contribution objects.
func (cl ContributionList) Unmarshal(payload []byte) sdk.ContributionList {
	return sdk.ContributionList{}.Unmarshal(payload)
}

// A Contribution wraps the verbs provided by the ESP API for Contributions,
// media assets that are associated with a Submission Batch.
type Contribution struct{ context *cli.Context }

// Index requests a list of all Contributions associated with the specified
// Submission Batch.
func (c Contribution) Index() sdk.ContributionList {
	return sdk.Contribution{}.Index(&client, getBatchID(c.context))
}

// Create associates a new Contribution with the specified Submission Batch.
func (c Contribution) Create() sdk.Createable {
	data := c.build()
	return client.Create(data.Path(), data)
}

// Update changes metadata for an existing Contribution.
func (c Contribution) Update() sdk.Createable {
	data := sdk.ContributionUpdate{c.build()}
	return client.Update(data.Path(), data)
}

// Delete destroys a specific Contribution.
func (c Contribution) Delete() { client.Delete(c.path()) }

// Get requests the metadata for a specific Contribution.
func (c Contribution) Get() sdk.DeserializedObject { return client.Get(c.path()) }

func (c Contribution) id() int { return getRequiredID(c.context, "contribution-id") }

func (c Contribution) path() string {
	obj := sdk.Contribution{
		ID:                c.id(),
		SubmissionBatchID: getBatchID(c.context),
	}
	return obj.Path()
}

func (c Contribution) build() sdk.Contribution {
	return sdk.Contribution{
		CameraShotDate:       c.context.String("camera-shot-date"),
		CollectionCode:       c.context.String("collection-code"),
		ContentProviderName:  c.context.String("content-provider-name"),
		ContentProviderTitle: c.context.String("content-provider-title"),
		CountryOfShoot:       c.context.String("country-of-shoot"),
		CreditLine:           c.context.String("credit-line"),
		ExternalFileLocation: c.context.String("external-file-location"),
		FileName:             c.context.String("file-name"),
		FilePath:             c.context.String("file-path"),
		Headline:             c.context.String("headline"),
		ID:                   c.context.Int("contribution-id"),
		IptcCategory:         c.context.String("iptc-category"),
		MimeType:             c.context.String("mime-type"),
		ParentSource:         c.context.String("parent-source"),
		RecordedDate:         c.context.String("recorded-date"),
		RiskCategory:         c.context.String("risk-category"),
		ShotSpeed:            c.context.String("shot-speed"),
		SiteDestination:      c.context.StringSlice("site-destination"),
		Source:               c.context.String("source"),
		SubmissionBatchID:    c.context.Int("submission-batch-id"),
		SubmittedToReviewAt:  c.context.String("submitted-to-review-at"),
		UploadBucket:         uploadBucket,
		UploadID:             c.context.String("upload-id"),
	}
}

// Unmarshal attempts to deserialize the provided JSON payload into a Contribution object.
func (c Contribution) Unmarshal(payload []byte) sdk.DeserializedObject {
	//return sdk.Contribution{}.Unmarshal(payload)
	return sdk.Unmarshal(payload)
}
