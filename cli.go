/*
GettyUp is a minimal Command Line Interface (CLI)
for Getty Images' Enterprise Submission Portal (ESP).

You will need a username and password that allows you
to log in to https://sandbox.espaws.com/ as well as
an API Key and and API Secret, both of which are accessible
at https://developer.gettyimages.com/apps/mykeys/.

These values can be provided on the command line as global
options or set as environment variables (recommended).

Retrieve a token needed for authenticating requests to the ESP API:
  gettyup token

Get a list of all of your Submission Batches:
  gettyup batch index

Retrieve a specific Submission Batch:
  gettyup batch get -b 86102

Delete a specific Submission Batch:
  gettyup batch delete -b 86125

Enable debug (verbose output) mode for any command with global option -D:
  gettyup -D token
  gettyup -D batch index
  gettyup -D batch get -b 86102

List all of the Contributions in a specific Submission Batch:
  gettyup contribution index -b 86102
*/
package main

import (
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "gettyup"
	app.Version = "0.0.1"
	app.Usage = "interact with the Getty Images ESP API"
	app.Author = "Jordan Peterson"
	app.Email = "dysolution@gmail.com"
	app.Flags = []cli.Flag{
		cli.BoolFlag{Name: "debug, D", Usage: "enable debug output"},
		cli.BoolFlag{Name: "quiet, q", Usage: "suppress pretty-printed JSON from ESP's response"},
		cli.StringFlag{Name: "key, k", Usage: "your ESP API key", EnvVar: "ESP_API_KEY"},
		cli.StringFlag{Name: "secret", Usage: "your ESP API secret", EnvVar: "ESP_API_SECRET"},
		cli.StringFlag{Name: "username, u", Usage: "your ESP username", EnvVar: "ESP_USERNAME"},
		cli.StringFlag{Name: "password, p", Usage: "your ESP password", EnvVar: "ESP_PASSWORD"},
		cli.StringFlag{Name: "token, t", Usage: "use an existing OAuth2 token", EnvVar: "ESP_TOKEN"},
		cli.StringFlag{
			Name:        "s3-bucket, b",
			Value:       "oregon",
			Usage:       "nearest S3 bucket = [germany|ireland|oregon|singapore|tokyo|virginia]",
			Destination: &uploadBucket,
			EnvVar:      "S3_BUCKET",
		},
	}
	app.Before = func(c *cli.Context) error {
		client = getClient(c.String("key"), c.String("secret"), c.String("username"), c.String("password"))
		quiet = c.Bool("quiet") == true
		if c.Bool("debug") == true {
			log.SetLevel(log.DebugLevel)
		}
		token = stringToToken(c.String("token"))
		return nil
	}
	app.Commands = []cli.Command{
		{
			Name:   "token",
			Usage:  "retrieve and print an OAuth2 authorization token",
			Action: func(c *cli.Context) { fmt.Println(Token()) },
		},
		{
			Name:  "batch",
			Usage: "work with Submission Batches",
			Subcommands: []cli.Command{
				{
					Name:   "create",
					Usage:  "create a new Submission Batch",
					Action: func(c *cli.Context) { prettyPrint(Batch{c}.Create()) },
					Flags: []cli.Flag{
						cli.StringFlag{Name: "submission-name, n"},
						cli.StringFlag{Name: "submission-type, t", Usage: batchTypes},
						cli.StringFlag{Name: "note"},
						cli.StringFlag{Name: "assignment-id"},
						cli.StringFlag{Name: "brief-id"},
						cli.StringFlag{Name: "event-id"},
						cli.BoolTFlag{Name: "save-extracted-metadata"},
					},
				},
				{
					Name:   "get",
					Usage:  "get a specific Submission Batch",
					Action: func(c *cli.Context) { prettyPrint(Batch{c}.Get()) },
					Flags: []cli.Flag{
						cli.StringFlag{Name: "submission-batch-id, b"},
					},
				},
				{
					Name:   "index",
					Usage:  "get all Submission Batches",
					Action: func(c *cli.Context) { prettyPrint(Batch{c}.Index()) },
				},
				{
					Name:   "update",
					Usage:  "update an existing Submission Batch",
					Action: func(c *cli.Context) { prettyPrint(Batch{c}.Update()) },
					Flags: []cli.Flag{
						cli.StringFlag{Name: "submission-batch-id, b"},
						cli.StringFlag{Name: "submission-name, n"},
						cli.StringFlag{Name: "note"},
					},
				},
				{
					Name:   "delete",
					Usage:  "delete an existing Submission Batch",
					Action: func(c *cli.Context) { Batch{c}.Delete() },
					Flags: []cli.Flag{
						cli.StringFlag{Name: "submission-batch-id, b"},
					},
				},
			},
		},
		{
			Name:  "contribution",
			Usage: "work with Contributions",
			Subcommands: []cli.Command{
				{
					Name:   "create",
					Usage:  "create a new Contribution within a Submission Batch",
					Action: func(c *cli.Context) { prettyPrint(Contribution{c}.Create()) },
					Flags: []cli.Flag{
						cli.StringFlag{Name: "camera-shot-date"},
						cli.StringFlag{Name: "collection-code"},
						cli.StringFlag{Name: "content-provider-name"},
						cli.StringFlag{Name: "content-provider-title"},
						cli.StringFlag{Name: "country-of-shoot"},
						cli.StringFlag{Name: "credit-line"},
						cli.StringFlag{Name: "external-file-location"},
						cli.StringFlag{Name: "file-name"},
						cli.StringFlag{Name: "file-path"},
						cli.StringFlag{Name: "headline"},
						cli.StringFlag{Name: "iptc-category"},
						cli.StringFlag{Name: "mime-type"},
						cli.StringFlag{Name: "parent-source"},
						cli.StringFlag{Name: "recorded-date"},
						cli.StringFlag{Name: "risk-category"},
						cli.StringFlag{Name: "shot-speed"},
						cli.StringFlag{Name: "source"},
						cli.StringFlag{Name: "submission-batch-id, b"},
						cli.StringFlag{Name: "submitted-to-review-at"},
						cli.StringFlag{Name: "upload-id"},
						cli.StringSliceFlag{Name: "site-destination"},
					},
				},
				{
					Name:   "get",
					Usage:  "get a specific Contribution",
					Action: func(c *cli.Context) { prettyPrint(Contribution{c}.Get()) },
					Flags: []cli.Flag{
						cli.StringFlag{Name: "submission-batch-id, b"},
						cli.StringFlag{Name: "contribution-id, c"},
					},
				},
				{
					Name:   "index",
					Usage:  "get all Contributions for a Submission Batch",
					Action: func(c *cli.Context) { prettyPrint(Contribution{c}.Index()) },
					Flags: []cli.Flag{
						cli.StringFlag{Name: "submission-batch-id, b"},
					},
				},
				{
					Name:   "update",
					Usage:  "update an existing Contribution",
					Action: func(c *cli.Context) { prettyPrint(Contribution{c}.Update()) },
					Flags: []cli.Flag{
						cli.StringFlag{Name: "camera-shot-date"},
						cli.StringFlag{Name: "collection-code"},
						cli.StringFlag{Name: "content-provider-name"},
						cli.StringFlag{Name: "content-provider-title"},
						cli.StringFlag{Name: "contribution-id, c"},
						cli.StringFlag{Name: "country-of-shoot"},
						cli.StringFlag{Name: "credit-line"},
						cli.StringFlag{Name: "external-file-location"},
						cli.StringFlag{Name: "file-name"},
						cli.StringFlag{Name: "file-path"},
						cli.StringFlag{Name: "headline"},
						cli.StringFlag{Name: "iptc-category"},
						cli.StringFlag{Name: "mime-type"},
						cli.StringFlag{Name: "parent-source"},
						cli.StringFlag{Name: "recorded-date"},
						cli.StringFlag{Name: "risk-category"},
						cli.StringFlag{Name: "shot-speed"},
						cli.StringFlag{Name: "source"},
						cli.StringFlag{Name: "submission-batch-id, b"},
						cli.StringFlag{Name: "submitted-to-review-at"},
						cli.StringFlag{Name: "upload-id"},
						cli.StringSliceFlag{Name: "site-destination"},
					},
				},
				{
					Name:   "delete",
					Usage:  "delete an existing Contribution",
					Action: func(c *cli.Context) { Contribution{c}.Delete() },
					Flags: []cli.Flag{
						cli.StringFlag{Name: "submission-batch-id, b"},
						cli.StringFlag{Name: "contribution-id, c"},
					},
				},
			},
		},
		{
			Name:  "release",
			Usage: "work with Releases",
			Subcommands: []cli.Command{
				{
					Name:   "create",
					Usage:  "create a new Release within a Submission Batch",
					Action: func(c *cli.Context) { prettyPrint(Release{c}.Create()) },
					Flags: []cli.Flag{
						cli.StringFlag{Name: "external-file-location"},
						cli.StringFlag{Name: "file-name"},
						cli.StringFlag{Name: "file-path"},
						cli.StringFlag{Name: "mime-type"},
						cli.StringFlag{Name: "model-date-of-birth"},
						cli.StringFlag{Name: "model-gender"},
						cli.StringFlag{Name: "release-type", Usage: releaseTypes},
						cli.StringFlag{Name: "submission-batch-id, b"},
						cli.StringSliceFlag{Name: "model-ethnicities"},
					},
				},
				{
					Name:   "get",
					Usage:  "get a specific Release",
					Action: func(c *cli.Context) { prettyPrint(Release{c}.Get()) },
					Flags: []cli.Flag{
						cli.StringFlag{Name: "submission-batch-id, b"},
						cli.StringFlag{Name: "release-id, r"},
					},
				},
				{
					Name:   "index",
					Usage:  "get all Releases for a Submission Batch",
					Action: func(c *cli.Context) { prettyPrint(Release{c}.Index()) },
					Flags: []cli.Flag{
						cli.StringFlag{Name: "submission-batch-id, b"},
					},
				},
				{
					Name:   "delete",
					Usage:  "delete an existing Release",
					Action: func(c *cli.Context) { Release{c}.Delete() },
					Flags: []cli.Flag{
						cli.StringFlag{Name: "submission-batch-id, b"},
						cli.StringFlag{Name: "release-id, r"},
					},
				},
			},
		},
		{
			Name:   "controlled_values",
			Usage:  "lists of values for fields with controlled vocabularies",
			Action: func(c *cli.Context) { GetControlledValues(c) },
		},
		{
			Name:   "transcoder",
			Usage:  "video transcoder mapping values",
			Action: func(c *cli.Context) { GetTranscoderMappings(c) },
		},
		{
			Name:   "keywords",
			Usage:  "controlled vocabularies for describing Contributions",
			Action: func(c *cli.Context) { GetKeywords(c) },
		},
		{
			Name:  "people",
			Usage: "controlled vocabularies and values for metadata about people",
			Subcommands: []cli.Command{
				{
					Name:   "compositions",
					Usage:  "all known values for Compositions",
					Action: func(c *cli.Context) { prettyPrint(TermService{}.GetCompositions(c)) },
				},
				{
					Name:   "expressions",
					Usage:  "all known values for Expressions",
					Action: func(c *cli.Context) { prettyPrint(TermService{}.GetExpressions(c)) },
				},
				{
					Name:   "number_of_people",
					Usage:  "all known values for Number Of People",
					Action: func(c *cli.Context) { prettyPrint(TermService{}.GetNumberOfPeople(c)) },
				},
				{
					Name:   "personalities",
					Usage:  "controlled vocabularies for describing famous personalities",
					Action: func(c *cli.Context) { GetPersonalities(c) },
				},
			},
		},
	}
	app.Run(os.Args)

}
