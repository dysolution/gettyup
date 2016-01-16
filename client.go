package main

import (
	sdk "github.com/dysolution/espsdk"
)

var client *sdk.Client
var quiet bool
var token sdk.Token

// Serializable objects can be Marshaled into JSON.
type Serializable interface {
	Marshal() ([]byte, error)
}

func getClient(key, secret, username, password string) *sdk.Client {
	return sdk.GetClient(key, secret, username, password, uploadBucket)
}

func stringToToken(tokenString string) sdk.Token {
	return sdk.Token(tokenString)
}
