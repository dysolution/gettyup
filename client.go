package main

import (
	sdk "github.com/dysolution/espsdk"
)

var client sdk.Client
var token sdk.Token

type Serializable interface {
	Marshal() ([]byte, error)
}

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

// Token is a memoizing wrapper for the API's token-providing function.
func Token() sdk.Token {
	if token != "" {
		return token
	} else {
		token = client.GetToken()
		return token
	}
}
