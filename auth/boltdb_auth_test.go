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

	auth := &BoltDBAuthenticator{
		DB: db,
	}

	auth.DB.Update(func(tx *bolt.Tx) error {
		b, _ := tx.CreateBucketIfNotExists([]byte("user"))
		return b.Put([]byte(TEST_USER), []byte(TEST_PASSWORD))
	})
	tokenStore := NewInMemoryTokenStore()
	return auth, tokenStore, id
}
