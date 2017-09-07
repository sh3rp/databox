package auth

import (
	"bytes"
	"encoding/gob"
	"errors"

	"golang.org/x/crypto/bcrypt"

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

func (db *BoltDBAuthenticator) Authenticate(username string, pass []byte) bool {
	user, err := db.getUser(username)
	if err != nil {
		log.Error().Msgf("Authenticate: error authenticating: %v", err)
		return false
	}
	return bcrypt.CompareHashAndPassword(user.Password, pass) == nil
}

func (db *BoltDBAuthenticator) AddUser(username string, pass []byte) error {
	user := &User{
		Username: username,
		Password: pass,
	}
	return db.saveUser(user)
}

func (db *BoltDBAuthenticator) DeleteUser(username string) error {
	return db.deleteUser(username)
}

func (db *BoltDBAuthenticator) saveUser(user *User) error {
	var buf bytes.Buffer
	err := gob.NewEncoder(&buf).Encode(user)

	if err != nil {
		return err
	}

	err = db.DB.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BOLT_USER_BUCKET))
		return bucket.Put([]byte(user.Username), buf.Bytes())
	})
	return err
}

func (db *BoltDBAuthenticator) deleteUser(username string) error {
	err := db.DB.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BOLT_USER_BUCKET))
		return bucket.Delete([]byte(username))
	})
	return err
}

func (db *BoltDBAuthenticator) getUser(username string) (*User, error) {
	var buf bytes.Buffer
	var data []byte

	err := db.DB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BOLT_USER_BUCKET))
		data = bucket.Get([]byte(username))
		if len(data) == 0 {
			return errors.New(ERR_AUTH_NO_USER)
		}
		_, err := buf.Write(data)
		return err
	})

	user := &User{}
	err = gob.NewDecoder(&buf).Decode(user)

	return user, err
}
