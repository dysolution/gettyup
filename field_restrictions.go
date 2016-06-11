package main

import (
	"github.com/codegangsta/cli"
	"github.com/dysolution/espsdk"
)

// GetFieldRestrictions uses optional parameters to refine a query for event metadata.
func GetFieldRestrictions(context *cli.Context) (*espsdk.FieldRestrictionResponse, error) {
	return client.GetFieldRestrictions(espsdk.FieldRestrictionQuery{
		UserID:                context.String("user_id"),
		FieldRestrictionsType: context.String("field_restrictions_type"),
	})
}
func registerFieldRestrictionCmds() {
	app.Commands = append(app.Commands, cli.Command{
		Name:  "field_restrictions",
		Usage: "Submission Batch metadata defaults",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "user_id",
				Usage: "limit the query/operation to a specific user",
			},
			cli.StringFlag{
				Name:  "field_restrictions_type",
				Usage: batchTypes,
			},
		},
		Action: func(c *cli.Context) {
			fr, _ := GetFieldRestrictions(c)
			pp(fr)
		},
	})
}
