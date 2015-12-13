package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/dysolution/espapi"
)

var (
	apiKey      string
	apiSecret   string
	espUsername string
	espPassword string
)

func BatchCreate(batch espapi.SubmissionBatch) {
	if batch.TypeIsValid() != true {
		validTypes := strings.Join(espapi.BatchTypes(), ", ")
		log.Errorf("Invalid submission batch type. Must be one of: %v", validTypes)
	} else if batch.NameIsValid() != true {
		log.Errorf("invalid batch name")
	} else {
		out, err := json.MarshalIndent(batch, "", "  ")
		if err != nil {
			log.Errorf("error marshaling batch")
		}
		fmt.Printf("%s\n", out)
		body, err := espapi.Response(apiKey, apiSecret, espUsername, espPassword)
		if err != nil {
			log.Errorf("error contacting API")
		}
		responseJson, err := json.Marshal(body)
		log.Infof("%s", responseJson)
		if errMsg := body["Error"]; errMsg != "" {
			log.Errorf(errMsg)
		}
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "gettyUp"
	app.Version = "0.0.1"
	app.Usage = "interact with the Getty Images ESP API"
	app.Author = "Jordan Peterson"
	app.Email = "dysolution@gmail.com"
	app.Action = func(c *cli.Context) {
		println("Use `gettyup help` for usage info")
	}
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "key, k",
			Usage:       "your key for the ESP API",
			EnvVar:      "ESP_API_KEY",
			Destination: &apiKey,
		},
		cli.StringFlag{
			Name:        "secret",
			Usage:       "your secret for the ESP API",
			EnvVar:      "ESP_API_SECRET",
			Destination: &apiSecret,
		},
		cli.StringFlag{
			Name:        "username, u",
			Usage:       "your ESP username",
			EnvVar:      "ESP_USERNAME",
			Destination: &espUsername,
		},
		cli.StringFlag{
			Name:        "password, p",
			Usage:       "your ESP password",
			EnvVar:      "ESP_PASSWORD",
			Destination: &espPassword,
		},
	}
	app.Commands = []cli.Command{
		{
			Name:  "batch",
			Usage: "work with Submission Batches",
			Action: func(c *cli.Context) {
				batch := espapi.SubmissionBatch{
					SubmissionName:        c.String("submission-name"),
					SubmissionType:        c.String("submission-type"),
					Note:                  c.String("note"),
					AssignmentId:          c.String("assignment-id"),
					BriefId:               c.String("brief-id"),
					EventId:               c.String("event-id"),
					SaveExtractedMetadata: c.Bool("save-extracted-metadata"),
				}
				BatchCreate(batch)
			},
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
		{
			Name:  "contribution",
			Usage: "work with Contributions",
			Action: func(c *cli.Context) {
				log.Errorf("not implemented")
			},
			Flags: []cli.Flag{},
		},
	}

	app.Run(os.Args)

}
