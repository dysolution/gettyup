package main

import (
	"fmt"
	"os"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/dysolution/espsdk"
	"github.com/dysolution/sleepwalker"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

// Log is the package-level logger.
var Log = logrus.New()

var app = cli.NewApp()

// patchVersion can be passed in at build time to be included in the
// usage output.
var patchVersion string

var client espsdk.Client

var quiet = false

var oAuthToken sleepwalker.Token

func init() {
	Log.Formatter = &prefixed.TextFormatter{TimestampFormat: time.RFC3339}
	app.Name = "gettyup"
	app.Version = "0.1.0" + patchVersion
	app.Usage = "interact with the Getty Images ESP API"
	app.Author = "Jordan Peterson"
	app.Email = "dysolution@gmail.com"
}

func registerCommands() {
	app.Commands = append(app.Commands, cli.Command{
		Name:  "token",
		Usage: "retrieve and print an OAuth2 authorization token",
		Action: func(c *cli.Context) {
			fmt.Println(client.Token)
		},
	})

	Batch{}.registerCmds()
	Contribution{}.registerCmds()
	Release{}.registerCmds()

	registerCVCmds()
}

func registerFlags() {
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "debug, D",
			Usage: "enable debug output",
		},
		cli.BoolFlag{
			Name:        "quiet, q",
			Usage:       "suppress pretty-printed output of JSON response",
			Destination: &quiet,
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
			Name:   "token, t",
			Usage:  "use an existing OAuth2 token",
			EnvVar: "ESP_TOKEN",
		},
		cli.StringFlag{
			Name:   "api-url, a",
			Value:  espsdk.SandboxAPI,
			Usage:  "API URL",
			EnvVar: "ESP_API_URL",
		},
	}
	buckets := "[germany|ireland|oregon|singapore|tokyo|virginia]"
	app.Flags = append(app.Flags, cli.StringFlag{
		Name:   "s3-bucket, b",
		Value:  "oregon",
		Usage:  "nearest S3 bucket = " + buckets,
		EnvVar: "S3_BUCKET",
	})
}

func main() {
	registerFlags()
	app.Before = func(c *cli.Context) error {
		client = espsdk.GetClient(
			c.String("key"),
			c.String("secret"),
			c.String("username"),
			c.String("password"),
			c.String("api-url"),
			Log,
		)
		if c.Bool("debug") == true {
			Log.Level = logrus.DebugLevel
		}
		oAuthToken = sleepwalker.Token(c.String("token"))
		return nil
	}
	registerCommands()
	app.Run(os.Args)
}
