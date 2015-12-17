package main

import (
	api "github.com/dysolution/espapi"
)

var client api.Client

func getClient(key, secret, username, password string) api.Client {
	return api.Client{
		api.Credentials{
			APIKey:      key,
			APISecret:   secret,
			ESPUsername: username,
			ESPPassword: password,
		},
		uploadBucket,
	}
}

var batchTypes = api.BatchTypes()
var releaseTypes = api.ReleaseTypes()

var token api.Token

// Token is a memoizing wrapper for the API's token-providing function.
func Token() api.Token {
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
