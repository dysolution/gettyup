package main

import (
	"github.com/codegangsta/cli"
	"github.com/dysolution/espsdk"
)

// GetKeywords requests suggestions from the Getty controlled vocabulary
// for the keywords provided.
//TODO: use search input from context
func GetKeywords(context *cli.Context) *espsdk.TermList {
	return client.GetTermList(Keywords)
}

// GetPersonalities requests suggestions from the Getty controlled vocabulary
// for the famous personalities provided.
//TODO: use search input from context
func GetPersonalities(context *cli.Context) *espsdk.TermList {
	return client.GetTermList(Personalities)
}

// GetControlledValues returns complete lists of values and descriptions for
// fields with controlled vocabularies, grouped by submission type.
func GetControlledValues(context *cli.Context) espsdk.ControlledValues {
	return client.GetControlledValues()
}

// GetTranscoderMappings lists acceptable transcoder mapping values
// for Getty and iStock video.
func GetTranscoderMappings(context *cli.Context) *espsdk.TranscoderMappingList {
	return client.GetTranscoderMappings()
}

func registerCVCmds() {
	app.Commands = append(app.Commands, cli.Command{
		Name:  "controlled_values",
		Usage: "lists of values for fields with controlled vocabularies",
		Action: func(c *cli.Context) {
			pp(GetControlledValues(c))
		},
	})
	app.Commands = append(app.Commands, cli.Command{
		Name:  "transcoder",
		Usage: "video transcoder mapping values",
		Action: func(c *cli.Context) {
			pp(GetTranscoderMappings(c))
		},
	})
	app.Commands = append(app.Commands, cli.Command{
		Name:  "keywords",
		Usage: "controlled vocabularies for describing Contributions",
		Action: func(c *cli.Context) {
			GetKeywords(c)
		},
	})
	app.Commands = append(app.Commands, cli.Command{
		Name:  "people",
		Usage: "controlled vocabularies and values for metadata about people",
		Subcommands: []cli.Command{
			{
				Name:  "compositions",
				Usage: "all known values for Compositions",
				Action: func(c *cli.Context) {
					pp(client.GetTermList(Compositions))
				},
			},
			{
				Name:  "expressions",
				Usage: "all known values for Expressions",
				Action: func(c *cli.Context) {
					pp(client.GetTermList(Expressions))
				},
			},
			{
				Name:  "number_of_people",
				Usage: "all known values for Number Of People",
				Action: func(c *cli.Context) {
					pp(client.GetTermList(NumberOfPeople))
				},
			},
			{
				Name:  "personalities",
				Usage: "controlled vocabularies for describing famous personalities",
				Action: func(c *cli.Context) {
					GetPersonalities(c)
				},
			},
		},
	})
}
