/*
GettyUp is a minimal Command Line Interface (CLI)
for Getty Images' Enterprise Submission Portal (ESP).

You will need a username and password that allows you
to log in to https://sandbox.espaws.com/ as well as
an API Key and and API Secret, both of which are accessible
at https://developer.gettyimages.com/apps/mykeys/.

These values can be provided on the command line as global
options or set as environment variables (recommended).
*/
package main

import (
	"fmt"
	"os"
	"strings"

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
		cli.BoolFlag{
			Name:  "debug, D",
			Usage: "enable debug output",
		},
		cli.StringFlag{
			Name:   "key, k",
			Usage:  "your ESP API key",
			EnvVar: "ESP_API_KEY",
		},
		cli.StringFlag{
			Name:   "secret",
			Usage:  "your ESP API secret",
			EnvVar: "ESP_API_SECRET",
		},
		cli.StringFlag{
			Name:   "username, u",
			Usage:  "your ESP username",
			EnvVar: "ESP_USERNAME",
		},
		cli.StringFlag{
			Name:   "password, p",
			Usage:  "your ESP password",
			EnvVar: "ESP_PASSWORD",
		},
		cli.StringFlag{
			Name:        "s3-bucket, b",
			Value:       "oregon",
			Usage:       "nearest S3 bucket = [germany|ireland|oregon|singapore|tokyo|virginia]",
			Destination: &uploadBucket,
			EnvVar:      "S3_BUCKET",
		},
	}
	app.Before = func(c *cli.Context) error {
		client = getClient(
			c.String("key"),
			c.String("secret"),
			c.String("username"),
			c.String("password"),
		)
		if c.Bool("debug") == true {
			log.SetLevel(log.DebugLevel)
		}
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
					Action: func(c *cli.Context) { CreateBatch(c) },
					Flags: []cli.Flag{
						cli.StringFlag{Name: "submission-name, n"},
						cli.StringFlag{
							Name:  "submission-type, t",
							Usage: fmt.Sprintf("[%s]", strings.Join(batchTypes, "|")),
						},
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
					Action: func(c *cli.Context) { GetBatch(c) },
					Flags: []cli.Flag{
						cli.StringFlag{Name: "submission-batch-id, b"},
					},
				},
				{
					Name:   "index",
					Usage:  "get all Submission Batches",
					Action: func(c *cli.Context) { GetBatches(c) },
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
					Action: func(c *cli.Context) { CreateContribution(c) },
					Flags: []cli.Flag{
						cli.StringFlag{Name: "submission-batch-id, b"},
						cli.StringFlag{Name: "file-name"},
						cli.StringFlag{Name: "file-path"},
						cli.StringFlag{Name: "submitted-to-review-at"},
						cli.StringFlag{Name: "external-file-location"},
						cli.StringFlag{Name: "upload-id"},
						cli.StringFlag{Name: "mime-type"},
						cli.StringFlag{Name: "camera-shot-date"},
						cli.StringFlag{Name: "collection-code"},
						cli.StringFlag{Name: "content-provider-name"},
						cli.StringFlag{Name: "content-provider-title"},
						cli.StringFlag{Name: "country-of-shoot"},
						cli.StringFlag{Name: "credit-line"},
						cli.StringFlag{Name: "headline"},
						cli.StringFlag{Name: "iptc-category"},
						cli.StringFlag{Name: "parent-source"},
						cli.StringFlag{Name: "recorded-date"},
						cli.StringFlag{Name: "risk-category"},
						cli.StringFlag{Name: "shot-speed"},
						cli.StringSliceFlag{Name: "site-destination"},
						cli.StringFlag{Name: "source"},
					},
				},
				{
					Name:   "get",
					Usage:  "get a specific Contribution",
					Action: func(c *cli.Context) { GetContribution(c) },
					Flags: []cli.Flag{
						cli.StringFlag{Name: "submission-batch-id, b"},
						cli.StringFlag{Name: "contribution-id, c"},
					},
				},
				{
					Name:   "index",
					Usage:  "get all Contributions for a Submission Batch",
					Action: func(c *cli.Context) { GetContributions(c) },
					Flags: []cli.Flag{
						cli.StringFlag{Name: "submission-batch-id, b"},
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
					Action: func(c *cli.Context) { CreateRelease(c) },
					Flags: []cli.Flag{
						cli.StringFlag{Name: "submission-batch-id, b"},
						cli.StringFlag{Name: "file-name"},
						cli.StringFlag{Name: "file-path"},
						cli.StringFlag{Name: "external-file-location"},
						cli.StringFlag{
							Name:  "release-type",
							Usage: fmt.Sprintf("[%s]", strings.Join(releaseTypes, "|")),
						},
						cli.StringFlag{Name: "model-date-of-birth"},
						cli.StringSliceFlag{Name: "model-ethnicities"},
						cli.StringFlag{Name: "model-gender"},
					},
				},
				{
					Name:   "get",
					Usage:  "get a specific Release",
					Action: func(c *cli.Context) { GetRelease(c) },
					Flags: []cli.Flag{
						cli.StringFlag{Name: "submission-batch-id, b"},
						cli.StringFlag{Name: "release-id, r"},
					},
				},
				{
					Name:   "index",
					Usage:  "get all Releases for a Submission Batch",
					Action: func(c *cli.Context) { GetReleases(c) },
					Flags: []cli.Flag{
						cli.StringFlag{Name: "submission-batch-id, b"},
					},
				},
			},
		},
	}
	app.Run(os.Args)

}
