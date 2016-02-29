package main

import (
	"github.com/dysolution/espsdk"
	"github.com/dysolution/sleepwalker"
)

var client espsdk.Client
var quiet bool
var token sleepwalker.Token

// Serializable objects can be Marshaled into JSON.
type Serializable interface {
	Marshal() ([]byte, error)
}

func getClient(key, secret, username, password string) espsdk.Client {
	return espsdk.GetClient(key, secret, username, password, Log)
}

func stringToToken(tokenString string) sleepwalker.Token {
	return sleepwalker.Token(tokenString)
}
