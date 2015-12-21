package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	sdk "github.com/dysolution/espsdk"
)

var uploadBucket string

// GetKeywords requests suggestions from the Getty controlled vocabulary
// for the keywords provided.
func GetKeywords(context *cli.Context) {
	//TODO: use search input from context
	get(Keywords)
}

// GetPersonalities requests suggestions from the Getty controlled vocabulary
// for the famous personalities provided.
func GetPersonalities(context *cli.Context) {
	//TODO: use search input from context
	get(Personalities)
}

// Private

func batchPath(context *cli.Context) string { return sdk.BatchPath(&sdk.Batch{ID: getBatchID(context)}) }

func getRequiredID(context *cli.Context, param string) int {
	v := context.Int(param)
	if v == 0 {
		log.Fatalf("--%s must be set", param)
	}
	return v
}

func getBatchID(context *cli.Context) int        { return getRequiredID(context, "submission-batch-id") }
func getReleaseID(context *cli.Context) int      { return getRequiredID(context, "release-id") }
func getContributionID(context *cli.Context) int { return getRequiredID(context, "contribution-id") }

func childPath(children string, context *cli.Context, childID string) string {
	var path string
	if childID == "" {
		path = batchPath(context) + "/" + children
	} else {
		path = batchPath(context) + "/" + children + "/" + childID
	}
	return path
}

func get(path string) []byte {
	request := sdk.NewRequest("GET", path, Token(), nil)
	result := client.PerformRequest(request)
	if result.Err != nil {
		log.Fatal(result.Err)
	}
	stats, err := result.Marshal()
	if err != nil {
		log.Fatal(result.Err)
	}
	log.Info(string(stats))
	log.Debugf("%s\n", result.Payload)
	return result.Payload
}
