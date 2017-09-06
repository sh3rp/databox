package auth

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestInMemoryAuth(t *testing.T) {
	suite.Run(t, &AuthTestSuite{
		New: getInMemoryAuth,
	})
}

func getInMemoryAuth() (Authenticator, TokenStore, string) {
	authenticator := &InMemoryAuthenticator{
		users: make(map[string]string),
	}
	authenticator.users[TEST_USER] = TEST_PASSWORD
	tokenStore := NewInMemoryTokenStore()
	return authenticator, tokenStore, ""
}
