package secure

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/sh3rp/crypt"
	"github.com/sh3rp/databox/auth"
	"github.com/sh3rp/databox/logger"
	"github.com/sh3rp/databox/msg"
)

type SecureFilter struct {
	Auth        auth.Authenticator
	TokenStore  auth.TokenStore
	credentials map[userBox][]byte
}

func NewSecureFilter(a auth.Authenticator, ts auth.TokenStore) *SecureFilter {
	return &SecureFilter{
		Auth:        a,
		TokenStore:  ts,
		credentials: make(map[userBox][]byte),
	}
}

type userBox struct {
	UserId string
	BoxId  string
}

func (f *SecureFilter) UnlockBox(token *msg.Token, box *msg.Box, password []byte) bool {
	if f.TokenStore.ValidateToken(token) != nil {
		delete(f.credentials, userBox{token.Username, box.Id.Id})
		return false
	}
	if !bytes.Equal(box.EncryptedSignature, GetSignature(password, box)) {
		return false
	}
	ub := userBox{token.Username, box.Id.Id}
	f.credentials[ub] = password
	return true
}

func (f *SecureFilter) GetKey(token *msg.Token, box *msg.Box) ([]byte, error) {
	if _, ok := f.credentials[userBox{token.Username, box.Id.Id}]; !ok {
		return nil, errors.New(fmt.Sprintf("Box %s has not been unlocked", box.Id.Id))
	}
	return f.credentials[userBox{token.Username, box.Id.Id}], nil
}

func (f *SecureFilter) IsUnlocked(token *msg.Token, box *msg.Box) bool {
	_, ok := f.credentials[userBox{token.Username, box.Id.Id}]
	return ok
}

func (f *SecureFilter) EncryptLink(token *msg.Token, link *msg.Link) *msg.Link {
	key := f.credentials[userBox{token.Username, link.Id.BoxId}]

	encryptedUrl, err := crypt.GCMEncrypt(key, []byte(link.Url))

	if err != nil {
		logger.E(err)
		return nil
	}

	encryptedName, err := crypt.GCMEncrypt(key, []byte(link.Name))

	if err != nil {
		logger.E(err)
		return nil
	}

	encryptedObj := link
	encryptedObj.Url = string(encryptedUrl)
	encryptedObj.Name = string(encryptedName)

	return encryptedObj
}

func (f *SecureFilter) DecryptLink(token *msg.Token, link *msg.Link) *msg.Link {
	key := f.credentials[userBox{token.Username, link.Id.BoxId}]

	decryptedUrl, err := crypt.GCMDecrypt(key, []byte(link.Url))

	if err != nil {
		logger.E(err)
		return nil
	}

	decryptedName, err := crypt.GCMDecrypt(key, []byte(link.Name))

	if err != nil {
		logger.E(err)
		return nil
	}

	decryptedObj := link
	decryptedObj.Url = string(decryptedUrl)
	decryptedObj.Name = string(decryptedName)

	return decryptedObj
}
