package auth

import (
	"errors"

	"github.com/boltdb/bolt"
	"github.com/rs/zerolog/log"
)

var BOLT_USER_BUCKET = "user"

type BoltDBAuthenticator struct {
	DB *bolt.DB
}

func NewBoltDBAuth(dbpath string) *BoltDBAuthenticator {
	db, err := bolt.Open(dbpath, 0600, nil)

	if err != nil {
		log.Error().Msgf("Error opening authenticator DB: %v\n", err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(BOLT_USER_BUCKET))
		return err
	})

	if err != nil {
		log.Error().Msgf("Error creating user auth bucket: %v\n", err)
	}

	return &BoltDBAuthenticator{
		DB: db,
	}
}

func (db *BoltDBAuthenticator) Authenticate(user, pass string) bool {
	persistedPass, err := db.getUser(user)
	if err != nil {
		log.Error().Msgf("Authenticate: error authenticating: %v", err)
		return false
	}
	return persistedPass == pass
}

func (db *BoltDBAuthenticator) AddUser(user, pass string) error {
	return db.saveUser(user, pass)
}

func (db *BoltDBAuthenticator) DeleteUser(user string) error {
	return db.saveUser(user, "")
}

func (db *BoltDBAuthenticator) saveUser(user, pass string) error {
	err := db.DB.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BOLT_USER_BUCKET))
		if pass == "" {
			return bucket.Delete([]byte(user))
		} else {
			return bucket.Put([]byte(user), []byte(pass))
		}
	})
	return err
}

func (db *BoltDBAuthenticator) getUser(user string) (string, error) {
	var password string

	err := db.DB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BOLT_USER_BUCKET))
		data := bucket.Get([]byte(user))
		if len(data) == 0 {
			return errors.New(ERR_AUTH_NO_USER)
		}
		password = string(data)
		return nil
	})

	return password, err
}
