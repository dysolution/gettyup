package main

import (
	sdk "github.com/dysolution/espsdk"
)

var client sdk.Client

func getClient(key, secret, username, password string) sdk.Client {
	return sdk.Client{
		sdk.Credentials{
			APIKey:      key,
			APISecret:   secret,
			ESPUsername: username,
			ESPPassword: password,
		},
		uploadBucket,
	}
}

var batchTypes = sdk.BatchTypes()
var releaseTypes = sdk.ReleaseTypes()

var token sdk.Token

// Token is a memoizing wrapper for the API's token-providing function.
func Token() sdk.Token {
	if token != "" {
		return token
	} else {
		token = client.GetToken()
		return token
	}
}

type Serializable interface {
	Marshal() ([]byte, error)
}
