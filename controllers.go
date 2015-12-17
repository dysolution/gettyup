package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
)

var uploadBucket string

const BatchesPath string = "/submission/v1/submission_batches"

func CreateBatch(context *cli.Context) {
	post(buildBatch(context), BatchesPath)
}

func CreateRelease(context *cli.Context) {
	post(buildRelease(context), batchPath(context)+"/releases")
}

func CreateContribution(context *cli.Context) {
	post(buildContribution(context), batchPath(context)+"/contributions")
}

func GetBatch(context *cli.Context) {
	get(batchPath(context))
}

func GetRelease(context *cli.Context) {
	getFromBatch("releases", context, getReleaseID(context))
}

func GetContribution(context *cli.Context) {
	getFromBatch("contributions", context, getContributionID(context))
}

func GetBatches(context *cli.Context) {
	get(BatchesPath)
}

func GetReleases(context *cli.Context) {
	getFromBatch("releases", context, "")
}

func GetContributions(context *cli.Context) {
	getFromBatch("contributions", context, "")
}

// Private

func batchPath(context *cli.Context) string {
	return BatchesPath + "/" + getBatchID(context)
}

func getRequiredValue(context *cli.Context, param string) string {
	v := context.String(param)
	if len(v) < 1 {
		log.Fatalf("--%s must be set", param)
	}
	return v
}

func getBatchID(context *cli.Context) string {
	return getRequiredValue(context, "submission-batch-id")
}

func getReleaseID(context *cli.Context) string {
	return getRequiredValue(context, "release-id")
}

func getContributionID(context *cli.Context) string {
	return getRequiredValue(context, "contribution-id")
}

// GetFromBatch returns objects associated with a Submission Batch.
//
// If a childID is provided, a single child object, e.g., Release{Id:123},
// otherwise all child objects associated with that release
// will be returned.
func getFromBatch(children string, context *cli.Context, childID string) {
	var path string
	if childID == "" {
		path = batchPath(context) + "/" + children
	} else {
		path = batchPath(context) + "/" + children + "/" + childID
	}
	get(path)
}

func get(path string) {
	response, err := client.Get(path, Token())
	if err != nil {
		log.Fatal(err)
	}
	log.Infof("%s\n", response)
}

func post(object Serializable, path string) {
	serializedObject, err := object.Marshal()
	if err != nil {
		log.Fatal(err)
	}
	response, err := client.Post(serializedObject, Token(), path)
	if err != nil {
		log.Fatal(err)
	}
	log.Infof("%s\n", response)
}
