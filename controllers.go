package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
)

var uploadBucket string

const (
	Batches                      string = "/submission/v1/submission_batches"
	ControlledValues             string = "/submission/v1/controlled_values/index"
	Keywords                     string = "/submission/v1/keywords/getty"
	Personalities                string = "/submission/v1/personalities"
	VideoTranscoderMappingValues string = "/submission/v1/video_transcoder_mapping_values"
)
const (
	Compositions   string = "/submission/v1/people_metadata/compositions"
	Expressions    string = "/submission/v1/people_metadata/expressions"
	NumberOfPeople string = "/submission/v1/people_metadata/number_of_people"
)

func GetCompositions(context *cli.Context)                 { get(Compositions) }
func GetControlledValues(context *cli.Context)             { get(ControlledValues) }
func GetExpressions(context *cli.Context)                  { get(Expressions) }
func GetKeywords(context *cli.Context)                     { get(Keywords) }
func GetNumberOfPeople(context *cli.Context)               { get(NumberOfPeople) }
func GetPersonalities(context *cli.Context)                { get(Personalities) }
func GetVideoTranscoderMappingValues(context *cli.Context) { get(VideoTranscoderMappingValues) }

func CreateBatch(context *cli.Context) {
	post(buildBatch(context), Batches)
}

func UpdateBatch(context *cli.Context) {
	put(buildBatchUpdate(context), batchPath(context))
}

func DeleteBatch(context *cli.Context) {
	_delete(batchPath(context))
}

func DeleteContribution(context *cli.Context) {
	deleteFromBatch("contributions", context, getContributionID(context))
}

func DeleteRelease(context *cli.Context) {
	deleteFromBatch("releases", context, getReleaseID(context))
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
	get(Batches)
}

func GetReleases(context *cli.Context) {
	getFromBatch("releases", context, "")
}

func GetContributions(context *cli.Context) {
	getFromBatch("contributions", context, "")
}

// Private

func batchPath(context *cli.Context) string {
	return Batches + "/" + getBatchID(context)
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

func deleteFromBatch(children string, context *cli.Context, childID string) {
	var path string
	if childID == "" {
		path = batchPath(context) + "/" + children
	} else {
		path = batchPath(context) + "/" + children + "/" + childID
	}
	_delete(path)
}
func get(path string) {
	response, err := client.Request("GET", path, Token(), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Infof("%s\n", response)
}

func _delete(path string) {
	response, err := client.Request("DELETE", path, Token(), nil)
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
	response, err := client.Request("POST", path, Token(), serializedObject)
	if err != nil {
		log.Fatal(err)
	}
	log.Infof("%s\n", response)
}

func put(object Serializable, path string) {
	serializedObject, err := object.Marshal()
	if err != nil {
		log.Fatal(err)
	}
	response, err := client.Request("PUT", path, Token(), serializedObject)
	if err != nil {
		log.Fatal(err)
	}
	log.Infof("%s\n", response)
}
