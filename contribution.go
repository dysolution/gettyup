package main

import (
	"github.com/codegangsta/cli"
	models "github.com/dysolution/espapi"
)

type Contribution struct{ context *cli.Context }

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
