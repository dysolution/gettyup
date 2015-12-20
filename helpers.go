package main

import (
	sdk "github.com/dysolution/espsdk"
)

func contribution(id int) sdk.Contribution { return sdk.Contribution{ID: id} }
func release(id int) sdk.Release           { return sdk.Release{ID: id} }
func batch(id int) *sdk.Batch              { return &sdk.Batch{ID: id} }
