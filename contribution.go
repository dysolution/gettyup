package main

import (
	"github.com/codegangsta/cli"
	sdk "github.com/dysolution/espsdk"
)

// A ContributionList contains zero or more Contribtions.
type ContributionList []Contribution

// Unmarshal attempts to deserialize the provided JSON payload into a slice of Contribution objects.
func (cl ContributionList) Unmarshal(payload []byte) sdk.ContributionList {
	return sdk.ContributionList{}.Unmarshal(payload)
}

func contribution(id int) sdk.Contribution { return sdk.Contribution{ID: id} }

// A Contribution wraps the verbs provided by the ESP API for Contributions,
// media assets that are associated with a Submission Batch.
type Contribution struct{ context *cli.Context }

// Index requests a list of all Contributions associated with the specified
// Submission Batch.
func (c Contribution) Index() sdk.ContributionList {
	return ContributionList{}.Unmarshal(get(childPath("contributions", c.context, "")))
}

// Get requests the metadata for a specific Contribution.
func (c Contribution) Get() sdk.Contribution {
	return contribution(c.id()).Get(&client, getBatchID(c.context))
}

// Create associates a new Contribution with the specified Submission Batch.
func (c Contribution) Create() sdk.Contribution { return c.Unmarshal(c.post()) }

// Update changes metadata for an existing Contribution.
func (c Contribution) Update() sdk.Contribution { return c.Unmarshal(c.put()) }

// Delete destroys a specific Contribution.
func (c Contribution) Delete() { contribution(c.id()).Delete(&client, getBatchID(c.context)) }

func (c Contribution) id() int { return getContributionID(c.context) }
func (c Contribution) path() string {
	return string(childPath("contributions", c.context, string(c.id())))
}
func (c Contribution) get() []byte { return get(c.path()) }
func (c Contribution) post() []byte {
	return post(c.build(), batchPath(c.context)+"/contributions")
}
func (c Contribution) put() []byte { return put(c.buildUpdate(), c.path()) }

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

func (c Contribution) buildUpdate() sdk.ContributionUpdate {
	return sdk.ContributionUpdate{c.build()}
}

// Unmarshal attempts to deserialize the provided JSON payload into a Contribution object.
func (c Contribution) Unmarshal(payload []byte) sdk.Contribution {
	return sdk.Contribution{}.Unmarshal(payload)
}
