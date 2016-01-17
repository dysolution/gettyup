package main

import (
	"github.com/dysolution/espsdk"
)

var client *espsdk.Client
var quiet bool
var token espsdk.Token

// Serializable objects can be Marshaled into JSON.
type Serializable interface {
	Marshal() ([]byte, error)
}

func getClient(key, secret, username, password string) *espsdk.Client {
	return espsdk.GetClient(key, secret, username, password, uploadBucket)
}

func stringToToken(tokenString string) espsdk.Token {
	return espsdk.Token(tokenString)
}
