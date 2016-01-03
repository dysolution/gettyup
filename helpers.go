package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	sdk "github.com/dysolution/espsdk"
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
func GetKeywords(context *cli.Context) []byte { return sdk.GetKeywords(&client) }

// GetPersonalities requests suggestions from the Getty controlled vocabulary
// for the famous personalities provided.
//TODO: use search input from context
func GetPersonalities(context *cli.Context) []byte { return sdk.GetPersonalities(&client) }

// GetControlledValues returns complete lists of values and descriptions for
// fields with controlled vocabularies, grouped by submission type.
func GetControlledValues(context *cli.Context) []byte { return sdk.GetControlledValues(&client) }

// GetTranscoderMappings lists acceptable transcoder mapping values
// for Getty and iStock video.
func GetTranscoderMappings(context *cli.Context) []byte { return sdk.GetTranscoderMappings(&client) }

// Private

// prettyPrint allows the CLI to pretty-print JSON responses by default. It
// can be disabled with the -q (--quiet) global option.
func prettyPrint(o PrettyPrintable) {
	if quiet != true {
		fmt.Println(o.PrettyPrint())
	}
}

func contribution(id int) sdk.Contribution { return sdk.Contribution{ID: id} }
func release(id int) sdk.Release           { return sdk.Release{ID: id} }
func batch(id int) *sdk.Batch              { return &sdk.Batch{ID: id} }

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
