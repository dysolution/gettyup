package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/codegangsta/cli"
	sdk "github.com/dysolution/espsdk"
)

type Contribution struct{ context *cli.Context }

func (c Contribution) Index()                { get(childPath("contributions", c.context, "")) }
func (c Contribution) Get() sdk.Contribution { return c.Unmarshal(get(c.path())) }
func (c Contribution) Create()               { post(c.build(c.context), batchPath(c.context)+"/contributions") }
func (c Contribution) Update()               { put(c.buildUpdate(), c.path()) }
func (c Contribution) Delete()               { _delete(c.path()) }
func (c Contribution) path() string          { return string(childPath("contributions", c.context, c.id())) }
func (c Contribution) id() string            { return getRequiredValue(c.context, "contribution-id") }

func (contribution Contribution) build(c *cli.Context) sdk.Contribution {
	return sdk.Contribution{
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

func (c Contribution) buildUpdate() sdk.ContributionUpdate {
	return sdk.ContributionUpdate{
		sdk.Contribution{
			Headline: c.context.String("headline"),
		},
	}
}

func (c Contribution) Unmarshal(payload []byte) sdk.Contribution {
	var contribution sdk.Contribution
	if err := json.Unmarshal(payload, &contribution); err != nil {
		log.Fatal(err)
	}
	return contribution
}

func (c Contribution) PrettyPrint() string {
	prettyOutput, err := c.Get().Marshal()
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%s\n", prettyOutput)
}
