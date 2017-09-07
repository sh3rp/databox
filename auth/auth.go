package auth

import "github.com/sh3rp/databox/msg"

var ERR_VALIDATION_EXPIRE = "Token expired"
var ERR_VALIDATION_TOKEN = "Token hash is incorrect"
var ERR_VALIDATION_USER = "User has not authenticated"
var ERR_AUTH_NO_USER = "No such user"

type Authenticator interface {
	Authenticate(string, []byte) bool
	AddUser(string, []byte) error
	DeleteUser(string) error
}

type TokenStore interface {
	GenerateToken(string) *msg.Token
	ValidateToken(*msg.Token) error
}

type User struct {
	Username string
	Password []byte
}
