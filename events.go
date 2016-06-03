package main

import (
	"github.com/codegangsta/cli"
	"github.com/dysolution/espsdk"
)

// GetEvents uses optional parameters to refine a query for event metadata.
func GetEvents(context *cli.Context) (*espsdk.EventResponse, error) {
	return client.GetEvents(espsdk.EventQuery{
		DateTo:           context.String("date_to"),
		DateFrom:         context.String("date_from"),
		EventName:        context.String("event_name"),
		PhotographerName: context.String("photographer_name"),
		MEID:             context.String("meid"),
	})
}

func registerEventCmds() {
	app.Commands = append(app.Commands, cli.Command{
		Name:  "events",
		Usage: "media event metadata (Sports, News, Entertainment)",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "date_from",
				Usage: "include events on or after this ISO8601 date",
			},
			cli.StringFlag{
				Name:  "date_to",
				Usage: "include events on or before this ISO8601 date",
			},
			cli.StringFlag{
				Name:  "event_name",
				Usage: "include events matching this name",
			},
			cli.StringFlag{
				Name:  "photographer_name",
				Usage: "include events matching this photographer name",
			},
			cli.StringFlag{
				Name:  "meid",
				Usage: "include events matching this Media Event ID",
			},
		},
		Action: func(c *cli.Context) {
			events, _ := GetEvents(c)
			pp(events)
		},
	})
}
