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

func getRequiredValue(context *cli.Context, param string) string {
	v := context.String(param)
	if len(v) < 1 {
		log.Fatalf("--%s must be set", param)
	}
	return v
}

func getBatchID(context *cli.Context) int {
	v := context.Int("submission-batch-id")
	if v == 0 {
		log.Fatal("--submission-batch-id must be set")
	}
	return v
}

func getReleaseID(context *cli.Context) int {
	v := context.Int("release-id")
	if v == 0 {
		log.Fatal("--release-id must be set")
	}
	return v
}

func getContributionID(context *cli.Context) string {
	return getRequiredValue(context, "contribution-id")
}

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

func _delete(path string) {
	request := sdk.NewRequest("DELETE", path, Token(), nil)
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
}

func post(object Serializable, path string) []byte {
	serializedObject, err := object.Marshal()
	if err != nil {
		log.Fatal(err)
	}

	request := sdk.NewRequest("POST", path, Token(), serializedObject)
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

func put(object Serializable, path string) []byte {
	serializedObject, err := object.Marshal()
	if err != nil {
		log.Fatal(err)
	}

	request := sdk.NewRequest("PUT", path, Token(), serializedObject)
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
