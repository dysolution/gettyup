package main

import (
	"fmt"
	"os"
	"strings"

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
		return nil
	}
	app.Commands = []cli.Command{
		{
			Name:   "token",
			Usage:  "retrieve and print an OAuth2 authorization token",
			Action: func(c *cli.Context) { fmt.Println(Token(c, client)) },
		},
		{
			Name:  "batch",
			Usage: "work with Submission Batches",
			Subcommands: []cli.Command{
				{
					Name:   "create",
					Action: func(c *cli.Context) { CreateBatch(c, client) },
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
			},
		},
		{
			Name:  "contribution",
			Usage: "work with Contributions",
			Subcommands: []cli.Command{
				{
					Name:   "create",
					Action: func(c *cli.Context) { CreateContribution(c, client) },
					Flags: []cli.Flag{
						cli.StringFlag{Name: "file-name"},
						cli.StringFlag{Name: "file-path"},
						cli.StringFlag{Name: "submitted-to-review-at"},
						cli.StringFlag{Name: "upload-bucket"},
						cli.StringFlag{Name: "external-file-location"},
						cli.StringFlag{Name: "upload-id"},
						cli.StringFlag{Name: "mime-type"},
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
					Action: func(c *cli.Context) { CreateRelease(c, client) },
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
			},
		},
	}

	app.Run(os.Args)

}
