package main

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/dysolution/sleepwalker"
)

// pp pretty-prints a JSON representation of the object.
func pp(object interface{}) {
	if quiet != true {
		prettyOutput, err := sleepwalker.Marshal(object)
		if err != nil {
			Log.WithFields(map[string]interface{}{
				"object": object.(string),
			}).Error(err)
		}
		fmt.Println(string(prettyOutput))
	}
}

func getRequiredID(context *cli.Context, param string) string {
	v := context.String(param)
	if v == "" {
		Log.Fatalf("--%s must be set", param)
	}
	return v
}

func getBatchID(context *cli.Context) string { return getRequiredID(context, "submission-batch-id") }
