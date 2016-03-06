package main

import (
	"path"
	"runtime"

	"github.com/codegangsta/cli"
	"github.com/dysolution/espsdk"
	"github.com/dysolution/sleepwalker"
)

// A ContributionList contains zero or more Contributions.
type ContributionList []Contribution

// Unmarshal attempts to deserialize the provided JSON payload into a slice of Contribution objects.
func (cl ContributionList) Unmarshal(payload []byte) (espsdk.ContributionList, error) {
	return espsdk.ContributionList{}.Unmarshal(payload)
}

// A Contribution wraps the verbs provided by the ESP API for Contributions,
// media assets that are associated with a Submission Batch.
type Contribution struct{ context *cli.Context }

// Index requests a list of all Contributions associated with the specified
// Submission Batch.
func (c Contribution) Index() espsdk.ContributionList {
	return espsdk.Contribution{}.Index(client, getBatchID(c.context))
}

// Create associates a new Contribution with the specified Submission Batch.
func (c Contribution) Create() *espsdk.Contribution {
	data := c.build()
	return c.do(client.Create, data)
}

// Update changes metadata for an existing Contribution.
func (c Contribution) Update() *espsdk.Contribution {
	data := c.build()
	return c.do(client.Update, data)
}

// Delete destroys a specific Contribution.
func (c Contribution) Delete() *espsdk.Contribution {
	data := c.build()
	return c.do(client.Delete, data)
}

// Get requests the metadata for a specific Contribution.
func (c Contribution) Get() *espsdk.Contribution {
	data := c.build()
	return c.do(client.Get, data)
}

// Submit submits an existing Contribution for review and publication.
func (c Contribution) Submit() *espsdk.Contribution {
	myPC, _, _, _ := runtime.Caller(0)
	desc := runtime.FuncForPC(myPC).Name() + ": "
	data := c.build()
	var contribution *espsdk.Contribution

	result, err := client.Put(data, data.Path()+"/submit")
	if err != nil {
		result.Log().Errorf("%s: %v", desc, err)
		return contribution
	}

	switch result.StatusCode {
	case 200:
		result.Log().Info(desc + "submitted")
		contribution, err = espsdk.Contribution{}.Unmarshal(result.Payload)
		if err != nil {
			result.Log().Errorf("%s: %v", desc, err)
		}
	case 401:
		result.Log().Error(desc + "unauthorized")
	case 403:
		result.Log().Error(desc + "forbidden")
	case 404:
		result.Log().Error(desc + "submission batch or contribution not found")
	case 422:
		result.Log().Error(desc + "unprocessable: already-submitted contribution or closed batch")
	}
	return contribution
}

func (c Contribution) do(fn func(sleepwalker.Findable) (sleepwalker.Result, error), data sleepwalker.Findable) *espsdk.Contribution {
	myPC, _, _, _ := runtime.Caller(1)
	desc := runtime.FuncForPC(myPC).Name()

	result, err := fn(data)
	if err != nil {
		return &espsdk.Contribution{}
	}
	result.Report(desc)
	if result.StatusCode >= 200 && result.StatusCode <= 300 {
		result.Log().Info(desc)
	}
	if len(result.Payload) == 0 {
		return &espsdk.Contribution{}
	}

	var contribution *espsdk.Contribution
	contribution, err = espsdk.Contribution{}.Unmarshal(result.Payload)
	if err != nil {
		return &espsdk.Contribution{}
	}
	return contribution
}

func (c Contribution) id() string { return getRequiredID(c.context, "contribution-id") }

func (c Contribution) path() string {
	obj := espsdk.Contribution{
		ID:                c.id(),
		SubmissionBatchID: getBatchID(c.context),
	}
	return obj.Path()
}

func (c Contribution) validateKeywords(input []string) []espsdk.Keyword {
	// TODO: iStock keywords
	mediaType := "image"
	if path.Ext(c.context.String("file-name")) == ".mov" {
		mediaType = "film"
	}
	return client.ValidateKeywords(input, mediaType)
}

func (c Contribution) build() espsdk.Contribution {

	keywords := []espsdk.Keyword{}
	inKeywords := c.context.StringSlice("keyword")
	if len(inKeywords) > 0 {
		keywords = c.validateKeywords(inKeywords)
	}

	personalities := []espsdk.Keyword{}
	inPersonalities := c.context.StringSlice("personality")
	if len(inPersonalities) > 0 {
		personalities = c.validateKeywords(inPersonalities)
	}

	inExpressions := c.context.StringSlice("facial-expression")
	var facialExpressions []espsdk.TermItem
	if len(inExpressions) > 0 {
		corpus := *client.GetTermList(Expressions)
		facialExpressions = espsdk.TermItem{}.ValidateList(inExpressions, corpus)
	}

	inPersonCompositions := c.context.StringSlice("person-composition")
	var personCompositions []espsdk.TermItem
	if len(inPersonCompositions) > 0 {
		corpus := *client.GetTermList(Compositions)
		personCompositions = espsdk.TermItem{}.ValidateList(inPersonCompositions, corpus)
	}

	inNumberOfPeople := c.context.String("number-of-people")
	var numberOfPeople espsdk.TermItem
	if inNumberOfPeople != "" {
		corpus := *client.GetTermList(NumberOfPeople)
		numberOfPeople = espsdk.TermItem{}.Validate(inNumberOfPeople, corpus)
	}

	return espsdk.Contribution{
		CameraShotDate:       c.context.String("camera-shot-date"),
		CollectionCode:       c.context.String("collection-code"),
		ContentProviderName:  c.context.String("content-provider-name"),
		ContentProviderTitle: c.context.String("content-provider-title"),
		CountryOfShoot:       c.context.String("country-of-shoot"),
		CreditLine:           c.context.String("credit-line"),
		ExternalFileLocation: c.context.String("external-file-location"),
		FacialExpressions:    facialExpressions,
		FileName:             c.context.String("file-name"),
		FilePath:             c.context.String("file-path"),
		Headline:             c.context.String("headline"),
		ID:                   c.context.String("contribution-id"),
		IPTCCategory:         c.context.String("iptc-category"),
		Keywords:             keywords,
		MimeType:             c.context.String("mime-type"),
		NumberOfPeople:       numberOfPeople,
		ParentSource:         c.context.String("parent-source"),
		Personalities:        personalities,
		PersonCompositions:   personCompositions,
		RecordedDate:         c.context.String("recorded-date"),
		RiskCategory:         c.context.String("risk-category"),
		ShotSpeed:            c.context.String("shot-speed"),
		SiteDestination:      c.context.StringSlice("site-destination"),
		Source:               c.context.String("source"),
		SubmissionBatchID:    c.context.String("submission-batch-id"),
		SubmittedToReviewAt:  c.context.String("submitted-to-review-at"),
		UploadBucket:         c.context.String("upload-bucket"),
		UploadID:             c.context.String("upload-id"),
	}
}

// Unmarshal attempts to deserialize the provided JSON payload into a Contribution object.
func (c Contribution) Unmarshal(payload []byte) (*espsdk.Contribution, error) {
	return espsdk.Contribution{}.Unmarshal(payload)
}

func (c Contribution) registerCmds() {
	inputFlags := []cli.Flag{
		cli.BoolFlag{Name: "exclusive-coverage"},
		cli.StringFlag{Name: "number-of-people"},
		cli.StringFlag{Name: "alternate-id"},
		cli.StringFlag{Name: "call-for-image"},
		cli.StringFlag{Name: "camera-shot-date"},
		cli.StringFlag{Name: "caption"},
		cli.StringFlag{Name: "city"},
		cli.StringFlag{Name: "collection-code"},
		cli.StringFlag{Name: "content-provider-name"},
		cli.StringFlag{Name: "content-provider-title"},
		cli.StringFlag{Name: "content-warnings"},
		cli.StringFlag{Name: "contribution-id, c"},
		cli.StringFlag{Name: "copyright"},
		cli.StringFlag{Name: "country-of-shoot"},
		cli.StringFlag{Name: "credit-line"},
		cli.StringFlag{Name: "event-id"},
		cli.StringFlag{Name: "exclusion-routes"},
		cli.StringFlag{Name: "external-file-location"},
		cli.StringFlag{Name: "facial-expressions"},
		cli.StringFlag{Name: "file-name"},
		cli.StringFlag{Name: "file-path"},
		cli.StringFlag{Name: "headline"},
		cli.StringFlag{Name: "inactive-date"},
		cli.StringFlag{Name: "iptc-caption-writer"},
		cli.StringFlag{Name: "iptc-category"},
		cli.StringFlag{Name: "mime-type"},
		cli.StringFlag{Name: "parent-source"},
		cli.StringFlag{Name: "recorded-date"},
		cli.StringFlag{Name: "risk-category"},
		cli.StringFlag{Name: "shot-speed"},
		cli.StringFlag{Name: "source"},
		cli.StringFlag{Name: "submission-batch-id, b"},
		cli.StringFlag{Name: "submitted-to-review-at"},
		cli.StringFlag{Name: "upload-id"},
		cli.StringSliceFlag{Name: "iptc-subject"},
		cli.StringSliceFlag{Name: "keyword"},
		cli.StringSliceFlag{Name: "facial-expression"},
		cli.StringSliceFlag{Name: "personality"},
		cli.StringSliceFlag{Name: "person-composition"},
		cli.StringSliceFlag{Name: "site-destination"},
	}
	app.Commands = append(app.Commands, cli.Command{
		Name:  "contribution",
		Usage: "work with Contributions",
		Subcommands: []cli.Command{
			{
				Name:  "create",
				Usage: "create a new Contribution within a Submission Batch",
				Action: func(c *cli.Context) {
					pp(Contribution{c}.Create())
				},
				Flags: inputFlags,
			},
			{
				Name:  "get",
				Usage: "get a specific Contribution",
				Action: func(c *cli.Context) {
					pp(Contribution{c}.Get())
				},
				Flags: []cli.Flag{
					cli.StringFlag{Name: "submission-batch-id, b"},
					cli.StringFlag{Name: "contribution-id, c"},
				},
			},
			{
				Name:  "index",
				Usage: "get all Contributions for a Submission Batch",
				Action: func(c *cli.Context) {
					pp(Contribution{c}.Index())
				},
				Flags: []cli.Flag{
					cli.StringFlag{Name: "submission-batch-id, b"},
				},
			},
			{
				Name:  "update",
				Usage: "update an existing Contribution",
				Action: func(c *cli.Context) {
					pp(Contribution{c}.Update())
				},
				Flags: inputFlags,
			},
			{
				Name:   "submit",
				Usage:  "submit an existing Contribution for review and publication",
				Action: func(c *cli.Context) { pp(Contribution{c}.Submit()) },
				Flags: []cli.Flag{
					cli.StringFlag{Name: "contribution-id, c"},
					cli.StringFlag{Name: "submission-batch-id, b"},
				},
			},
			{
				Name:   "delete",
				Usage:  "delete an existing Contribution",
				Action: func(c *cli.Context) { Contribution{c}.Delete() },
				Flags: []cli.Flag{
					cli.StringFlag{Name: "submission-batch-id, b"},
					cli.StringFlag{Name: "contribution-id, c"},
				},
			},
		},
	})
}
