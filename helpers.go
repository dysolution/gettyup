package main

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/dysolution/espsdk"
	"github.com/dysolution/sleepwalker"
)

// PrettyPrintable applies to all objects that should have an easy-to-read
// JSON representation of themselves availalbe for printing.
type PrettyPrintable interface {
	PrettyPrint() string
}

var uploadBucket string

// GetKeywords requests suggestions from the Getty controlled vocabulary
// for the keywords provided.
//TODO: use search input from context
func GetKeywords(context *cli.Context) espsdk.TermList {
	return espsdk.Client{}.GetTermList(espsdk.Keywords)
}

// GetPersonalities requests suggestions from the Getty controlled vocabulary
// for the famous personalities provided.
//TODO: use search input from context
func GetPersonalities(context *cli.Context) espsdk.TermList {
	return espsdk.Client{}.GetTermList(espsdk.Personalities)
}

// GetControlledValues returns complete lists of values and descriptions for
// fields with controlled vocabularies, grouped by submission type.
func GetControlledValues(context *cli.Context) espsdk.TermList {
	return espsdk.Client{}.GetTermList(espsdk.ControlledValues)
}

// GetTranscoderMappings lists acceptable transcoder mapping values
// for Getty and iStock video.
func GetTranscoderMappings(context *cli.Context) espsdk.TranscoderMappingList {
	return espsdk.Client{}.GetTranscoderMappings()
}

// Private

func prettyPrint(object interface{}) {
	if quiet != true {
		prettyOutput, err := sleepwalker.Marshal(object)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(prettyOutput))
	}
}

func contribution(id int) espsdk.Contribution { return espsdk.Contribution{ID: id} }
func release(id int) espsdk.Release           { return espsdk.Release{ID: id} }
func batch(id int) *espsdk.Batch              { return &espsdk.Batch{ID: id} }

func getRequiredID(context *cli.Context, param string) int {
	v := context.Int(param)
	if v == 0 {
		log.Fatalf("--%s must be set", param)
	}
	return v
}

func getBatchID(context *cli.Context) int { return getRequiredID(context, "submission-batch-id") }

// func get(path string) []byte {
// 	request := sdk.NewRequest("GET", path, Token(), nil)
// 	result := client.PerformRequest(request)
// 	if result.Err != nil {
// 		log.Fatal(result.Err)
// 	}
// 	stats, err := result.Marshal()
// 	if err != nil {
// 		log.Fatal(result.Err)
// 	}
// 	log.Info(string(stats))
// 	log.Debugf("%s\n", result.Payload)
// 	return result.Payload
// }
