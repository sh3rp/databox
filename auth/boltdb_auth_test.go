package auth

import (
	"os"
	"testing"

	"github.com/boltdb/bolt"
	"github.com/sh3rp/databox/db"

	"github.com/stretchr/testify/suite"
)

func TestBoltDBAuth(t *testing.T) {
	suite.Run(t, &AuthTestSuite{
		New:      getBoltDBAuth,
		TearDown: tearDown,
	})
}

func tearDown(id string) {
	os.RemoveAll("/tmp/bolt" + id)
}

func getBoltDBAuth() (Authenticator, TokenStore, string) {
	id := db.GenerateID()
	db, _ := bolt.Open("/tmp/bolt"+id, 0600, nil)

	a := &BoltDBAuthenticator{
		DB: db,
	}
	a.AddUser(TEST_USER, TEST_PASSWORD)
	tokenStore := NewInMemoryTokenStore()
	return a, tokenStore, id
}
