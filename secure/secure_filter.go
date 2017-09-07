package secure

import (
	"github.com/sh3rp/crypt"
	"github.com/sh3rp/databox/auth"
	"github.com/sh3rp/databox/msg"
)

type SecureFilter struct {
	Auth auth.Authenticator
}

func (f *SecureFilter) EncryptLink(token *msg.Token, link *msg.Link) (*msg.Link, error) {
	key, err := f.Auth.GetEncryptionKey(token.Username)
	if err != nil {
		return nil, err
	}

	encryptedUrl, err := crypt.GCMEncrypt(key, []byte(link.Url))
	encryptedName, err := crypt.GCMEncrypt(key, []byte(link.Name))

	encryptedObj := link
	encryptedObj.Url = string(encryptedUrl)
	encryptedObj.Name = string(encryptedName)

	return encryptedObj, nil
}

func (f *SecureFilter) DecryptLink(token *msg.Token, link *msg.Link) (*msg.Link, error) {
	key, err := f.Auth.GetEncryptionKey(token.Username)
	if err != nil {
		return nil, err
	}

	decryptedUrl, err := crypt.GCMDecrypt(key, []byte(link.Url))
	decryptedName, err := crypt.GCMDecrypt(key, []byte(link.Name))

	decryptedObj := link
	decryptedObj.Url = string(decryptedUrl)
	decryptedObj.Name = string(decryptedName)

	return decryptedObj, nil
}
