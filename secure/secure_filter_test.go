package secure

import (
	"testing"

	"github.com/sh3rp/databox/auth"
	"github.com/sh3rp/databox/msg"
	"github.com/stretchr/testify/assert"
)

func TestLinkCrypto(t *testing.T) {
	authenticator := auth.NewInMemoryAuthenticator()
	authenticator.AddUser("test", "password")

	filter := &SecureFilter{
		Auth: authenticator,
	}

	link := &msg.Link{
		Url:  "http://www.cnn.com",
		Name: "cnn",
	}

	token := &msg.Token{
		Username: "test",
	}

	secureLink, err := filter.EncryptLink(token, link)

	assert.Nil(t, err)
	assert.NotNil(t, secureLink)
	assert.NotEqual(t, secureLink.Url, "http://www.cnn.com")

	plaintextLink, err := filter.DecryptLink(token, secureLink)

	assert.Nil(t, err)
	assert.NotNil(t, plaintextLink)
	assert.Equal(t, plaintextLink.Url, "http://www.cnn.com")
}
