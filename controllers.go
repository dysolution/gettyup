package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	sdk "github.com/dysolution/espsdk"
)

var uploadBucket string

func GetKeywords(context *cli.Context) {
	//TODO: use search input from context
	get(Keywords)
}
func GetPersonalities(context *cli.Context) {
	//TODO: use search input from context
	get(Personalities)
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
	params := sdk.RequestParams{"GET", path, Token(), nil}
	response, err := client.Request(&params)
	if err != nil {
		log.Fatal(err)
	}
	log.Debugf("%s\n", response)
	return response
}

func _delete(path string) {
	params := sdk.RequestParams{"DELETE", path, Token(), nil}
	response, err := client.Request(&params)
	if err != nil {
		log.Fatal(err)
	}
	log.Debugf("%s\n", response)
}

func post(object Serializable, path string) {
	serializedObject, err := object.Marshal()
	if err != nil {
		log.Fatal(err)
	}
	params := sdk.RequestParams{"POST", path, Token(), serializedObject}
	response, err := client.Request(&params)
	if err != nil {
		log.Fatal(err)
	}
	log.Debugf("%s\n", response)
}

func put(object Serializable, path string) {
	serializedObject, err := object.Marshal()
	if err != nil {
		log.Fatal(err)
	}
	params := sdk.RequestParams{"PUT", path, Token(), serializedObject}
	response, err := client.Request(&params)
	if err != nil {
		log.Fatal(err)
	}
	log.Debugf("%s\n", response)
}
