package main

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/dysolution/espsdk"
	"github.com/dysolution/sleepwalker"
)

var batchTypes = string(strings.Join(espsdk.Batch{}.ValidTypes(), " OR "))

// A Batch wraps the verbs provided by the ESP API for Submission Batches.
type Batch struct {
	context *cli.Context
}

// Index requests a list of all Submission Batches belonging to the user.
func (b Batch) Index() espsdk.BatchList {
	return espsdk.Batch{}.Index(client)
}

// Get requests the metadata for a specific Submission Batch.
func (b Batch) Get() *espsdk.Batch {
	data := b.build()
	return b.do(client.Get, data)
}

// Create adds a new Submission Batch.
func (b Batch) Create() *espsdk.Batch {
	data := b.build()
	return b.do(client.Create, data)
}

// Update changes fields for an existing Submission Batch.
func (b Batch) Update() *espsdk.Batch {
	data := espsdk.BatchUpdate{Batch: b.build()}
	return b.do(client.Update, data)
}

// Delete destroys a specific Submission Batch.
func (b Batch) Delete() *espsdk.Batch {
	data := b.build()
	return b.do(client.Delete, data)
}

// Last returns the newest Submission Batch.
func (b Batch) Last() espsdk.Batch {
	return espsdk.Batch{}.Index(client).Last()
}

func (b Batch) do(fn func(sleepwalker.Findable) (sleepwalker.Result, error), data sleepwalker.Findable) *espsdk.Batch {
	myPC, _, _, _ := runtime.Caller(1)
	desc := runtime.FuncForPC(myPC).Name()

	result, err := fn(data)
	if err != nil {
		return &espsdk.Batch{}
	}
	result.Report(desc)
	if result.StatusCode >= 200 && result.StatusCode <= 300 {
		result.Log().Info(desc)
	}
	if len(result.Payload) == 0 {
		return &espsdk.Batch{}
	}

	var batch *espsdk.Batch
	batch, err = espsdk.Batch{}.Unmarshal(result.Payload)
	if err != nil {
		return &espsdk.Batch{}
	}
	return batch
}

// Unmarshal attempts to deserialize the provided JSON payload into a SubmissionBatch object.
func (b Batch) Unmarshal(payload []byte) *espsdk.Batch {
	batch, err := espsdk.Batch{}.Unmarshal(payload)
	if err != nil {
		return &espsdk.Batch{}
	}
	return batch
}

func (b Batch) registerCmds() {
	app.Commands = append(app.Commands, cli.Command{
		Name:  "batch",
		Usage: "work with Submission Batches",
		Subcommands: []cli.Command{
			{
				Name:   "create",
				Usage:  "create a new Submission Batch",
				Action: func(c *cli.Context) { pp(Batch{c}.Create()) },
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
				Action: func(c *cli.Context) { pp(Batch{c}.Get()) },
				Flags: []cli.Flag{
					cli.StringFlag{Name: "submission-batch-id, b"},
				},
			},
			{
				Name:   "index",
				Usage:  "get all Submission Batches",
				Action: func(c *cli.Context) { pp(Batch{c}.Index()) },
			},
			{
				Name:   "update",
				Usage:  "update an existing Submission Batch",
				Action: func(c *cli.Context) { pp(Batch{c}.Update()) },
				Flags: []cli.Flag{
					cli.StringFlag{Name: "submission-batch-id, b"},
					cli.StringFlag{Name: "submission-name, n"},
					cli.StringFlag{Name: "save-extracted-metadata"},
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
			{
				Name:  "last",
				Usage: "get the most recent Submission Batch ID",
				Action: func(c *cli.Context) {
					fmt.Print(Batch{c}.Last().ID)
				},
			},
		},
	})
}

func (b Batch) id() string {
	return getBatchID(b.context)
}

func (b Batch) path() string {
	return espsdk.Batch{ID: b.id()}.Path()
}

func (b Batch) build() espsdk.Batch {
	return espsdk.Batch{
		BriefID:               b.context.String("brief-id"),
		EventID:               b.context.String("event-id"),
		ID:                    b.context.String("submission-batch-id"),
		Note:                  b.context.String("note"),
		AssignmentID:          b.context.String("assignment-id"),
		SaveExtractedMetadata: b.context.Bool("save-extracted-metadata"),
		SubmissionName:        b.context.String("submission-name"),
		SubmissionType:        b.context.String("submission-type"),
	}
}
