package main

import (
	"fmt"
	"os"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/dysolution/espapi"
)

var client espapi.Client

func BuildBatch(c *cli.Context) espapi.SubmissionBatch {
	return espapi.SubmissionBatch{
		SubmissionName:        c.String("submission-name"),
		SubmissionType:        c.String("submission-type"),
		Note:                  c.String("note"),
		AssignmentId:          c.String("assignment-id"),
		BriefId:               c.String("brief-id"),
		EventId:               c.String("event-id"),
		SaveExtractedMetadata: c.Bool("save-extracted-metadata"),
	}
}

func BuildRelease(c *cli.Context) espapi.Release {
	return espapi.Release{
		FileName:             c.String("file-name"),
		FilePath:             c.String("file-path"),
		ExternalFileLocation: c.String("external-file-location"),
		ReleaseType:          c.String("release-type"),
		ModelDateOfBirth:     c.String("model-date-of-birth"),
		ModelEthnicities:     c.StringSlice("model-ethnicities"),
		ModelGender:          c.String("model-gender"),
	}
}

func BuildContribution(c *cli.Context) espapi.Contribution {
	return espapi.Contribution{
		FileName:             c.String("file-name"),
		FilePath:             c.String("file-path"),
		SubmittedToReviewAt:  c.String("submitted-to-review-at"),
		UploadBucket:         c.String("upload-bucket"),
		ExternalFileLocation: c.String("external-file-location"),
		UploadId:             c.String("upload-id"),
		MimeType:             c.String("mime-type"),
	}
}

func CreateBatch(context *cli.Context, client espapi.Client) {
	batch, err := BuildBatch(context).Marshal()
	if err != nil {
		log.Errorf("error creating batch")
	}
	client.PostBatch(batch)
}

func CreateRelease(context *cli.Context, client espapi.Client) {
	release, err := BuildRelease(context).Marshal()
	if err != nil {
		log.Errorf("error creating release")
	}
	client.PostRelease(release)
}

func CreateContribution(context *cli.Context, client espapi.Client) {
	release, err := BuildContribution(context).Marshal()
	if err != nil {
		log.Errorf("error creating contribution")
	}
	client.PostContribution(release)
}
func main() {
	app := cli.NewApp()
	app.Name = "gettyup"
	app.Version = "0.0.1"
	app.Usage = "interact with the Getty Images ESP API"
	app.Author = "Jordan Peterson"
	app.Email = "dysolution@gmail.com"
	app.Action = func(c *cli.Context) {
		println("Use `gettyup help` for usage info")
	}
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "key, k",
			Usage:  "your key for the ESP API",
			EnvVar: "ESP_API_KEY",
		},
		cli.StringFlag{
			Name:   "secret",
			Usage:  "your secret for the ESP API",
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
	}
	app.Before = func(c *cli.Context) error {
		client = espapi.Client{espapi.Credentials{
			ApiKey:      c.String("key"),
			ApiSecret:   c.String("secret"),
			EspUsername: c.String("username"),
			EspPassword: c.String("password"),
		},
		}
		return nil
	}
	app.Commands = []cli.Command{
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
							Usage: fmt.Sprintf("[%s]", strings.Join(espapi.BatchTypes(), "|")),
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
						cli.StringFlag{Name: "file-name"},
						cli.StringFlag{Name: "file-path"},
						cli.StringFlag{Name: "external-file-location"},
						cli.StringFlag{Name: "release-type"},
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
