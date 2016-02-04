package main

import (
	"runtime"

	"github.com/codegangsta/cli"
	"github.com/dysolution/espsdk"
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
	desc := "Contribution.Create: "
	data := c.build()
	var contribution *espsdk.Contribution

	result, err := client.Create(data)
	if err != nil {
		result.Log().Errorf("%s: %v", desc, err)
		return contribution
	}

	switch result.StatusCode {
	case 201:
		result.Log().Info(desc + "created")
		contribution, err = espsdk.Contribution{}.Unmarshal(result.Payload)
		if err != nil {
			result.Log().Errorf("%s: %v", desc, err)
		}
	case 401:
		result.Log().Error(desc + "unauthorized")
	case 403:
		result.Log().Error(desc + "forbidden")
	case 404:
		result.Log().Error(desc + "submission batch not found")
	case 422:
		result.Log().Error(desc + "submission batch is closed")
	}
	return contribution
}

// Update changes metadata for an existing Contribution.
func (c Contribution) Update() *espsdk.Contribution {
	myPC, _, _, _ := runtime.Caller(0)
	desc := runtime.FuncForPC(myPC).Name() + ": "
	data := c.build()
	var contribution *espsdk.Contribution

	result, err := client.Update(data)
	if err != nil {
		result.Log().Errorf("%s: %v", desc, err)
		return contribution
	}

	switch result.StatusCode {
	case 200:
		result.Log().Info(desc + "updated")
		contribution, err = espsdk.Contribution{}.Unmarshal(result.Payload)
		if err != nil {
			result.Log().Errorf("%s: %v", desc, err)
		}
	case 401:
		result.Log().Error(desc + "unauthorized")
	case 403:
		result.Log().Error(desc + "forbidden")
	case 404:
		result.Log().Error(desc + "submission batch not found")
	case 422:
		result.Log().Error(desc + "unprocessable: already-submitted contribution or closed batch")
	}
	return contribution
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

// Delete destroys a specific Contribution.
func (c Contribution) Delete() *espsdk.Contribution {
	myPC, _, _, _ := runtime.Caller(0)
	desc := runtime.FuncForPC(myPC).Name() + ": "
	data := c.build()
	var contribution *espsdk.Contribution

	result, err := client.Delete(data)
	if err != nil {
		result.Log().Errorf("%s: %v", desc, err)
		return contribution
	}

	switch result.StatusCode {
	case 204:
		result.Log().Info(desc + "deleted")
	case 401:
		result.Log().Error(desc + "unauthorized")
	case 403:
		result.Log().Error(desc + "forbidden")
	case 404:
		result.Log().Error(desc + "submission batch or contribution not found")
	case 422:
		result.Log().Error(desc + "unprocessable: already-submitted contribution or closed batch")
	}
	// successful deletion usually returns a 204 without a payload/body
	if len(result.Payload) > 0 {
		contribution, err = espsdk.Contribution{}.Unmarshal(result.Payload)
		if err != nil {
			result.Log().Errorf("%s: %v", desc, err)
		}
	}
	return contribution
}

// Get requests the metadata for a specific Contribution.
func (c Contribution) Get() *espsdk.Contribution {
	desc := "Contribution.Get"
	data := c.build()
	var contribution *espsdk.Contribution

	result, err := client.Get(data)
	if err != nil {
		result.Log().Error(desc)
		return contribution
	}
	if result.StatusCode == 404 {
		result.Log().Error(desc)
		return contribution
	}
	result.Log().Info(desc)
	contribution, err = espsdk.Contribution{}.Unmarshal(result.Payload)
	if err != nil {
		result.Log().Errorf("%s: %v", desc, err)
		return contribution
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

func (c Contribution) build() espsdk.Contribution {
	return espsdk.Contribution{
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
		ID:                   c.context.String("contribution-id"),
		IPTCCategory:         c.context.String("iptc-category"),
		MimeType:             c.context.String("mime-type"),
		ParentSource:         c.context.String("parent-source"),
		RecordedDate:         c.context.String("recorded-date"),
		RiskCategory:         c.context.String("risk-category"),
		ShotSpeed:            c.context.String("shot-speed"),
		SiteDestination:      c.context.StringSlice("site-destination"),
		Source:               c.context.String("source"),
		SubmissionBatchID:    c.context.String("submission-batch-id"),
		SubmittedToReviewAt:  c.context.String("submitted-to-review-at"),
		UploadBucket:         uploadBucket,
		UploadID:             c.context.String("upload-id"),
	}
}

// Unmarshal attempts to deserialize the provided JSON payload into a Contribution object.
func (c Contribution) Unmarshal(payload []byte) (*espsdk.Contribution, error) {
	return espsdk.Contribution{}.Unmarshal(payload)
}
