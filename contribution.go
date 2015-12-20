package main

import (
	"encoding/json"
	"log"

	"github.com/codegangsta/cli"
	sdk "github.com/dysolution/espsdk"
)

// A Contribution wraps the verbs provided by the ESP API for Contributions,
// media assets that are associated with a Submission Batch.
type Contribution struct{ context *cli.Context }

// Index requests a list of all Contributions associated with the specified
// Submission Batch.
func (c Contribution) Index() { get(childPath("contributions", c.context, "")) }

// Get requests the metadata for a specific Contribution.
func (c Contribution) Get() sdk.Contribution { return c.Unmarshal(c.get()) }

// Create associates a new Contribution with the specified Submission Batch.
func (c Contribution) Create() sdk.Contribution { return c.Unmarshal(c.post()) }

// Update changes metadata for an existing Contribution.
func (c Contribution) Update() sdk.Contribution { return c.Unmarshal(c.put()) }

// Delete destroys a specific Contribution.
func (c Contribution) Delete()      { _delete(c.path()) }
func (c Contribution) path() string { return string(childPath("contributions", c.context, c.id())) }
func (c Contribution) id() string   { return getRequiredValue(c.context, "contribution-id") }
func (c Contribution) get() []byte  { return get(c.path()) }
func (c Contribution) post() []byte {
	return post(c.build(c.context), batchPath(c.context)+"/contributions")
}
func (c Contribution) put() []byte { return put(c.buildUpdate(), c.path()) }

func (c Contribution) build(context *cli.Context) sdk.Contribution {
	return sdk.Contribution{
		CameraShotDate:       context.String("camera-shot-date"),
		CollectionCode:       context.String("collection-code"),
		ContentProviderName:  context.String("content-provider-name"),
		ContentProviderTitle: context.String("content-provider-title"),
		CountryOfShoot:       context.String("country-of-shoot"),
		CreditLine:           context.String("credit-line"),
		ExternalFileLocation: context.String("external-file-location"),
		FileName:             context.String("file-name"),
		FilePath:             context.String("file-path"),
		Headline:             context.String("headline"),
		IptcCategory:         context.String("iptc-category"),
		MimeType:             context.String("mime-type"),
		ParentSource:         context.String("parent-source"),
		RecordedDate:         context.String("recorded-date"),
		RiskCategory:         context.String("risk-category"),
		ShotSpeed:            context.String("shot-speed"),
		SiteDestination:      context.StringSlice("site-destination"),
		Source:               context.String("source"),
		SubmittedToReviewAt:  context.String("submitted-to-review-at"),
		UploadBucket:         uploadBucket,
		UploadId:             context.String("upload-id"),
	}
}

func (c Contribution) buildUpdate() sdk.ContributionUpdate {
	return sdk.ContributionUpdate{
		sdk.Contribution{
			Headline: c.context.String("headline"),
		},
	}
}

// Unmarshal attempts to deserialize the provided JSON payload into a
// Contribution object as defined by the SDK.
func (c Contribution) Unmarshal(payload []byte) sdk.Contribution {
	var contribution sdk.Contribution
	if err := json.Unmarshal(payload, &contribution); err != nil {
		log.Fatal(err)
	}
	return contribution
}
