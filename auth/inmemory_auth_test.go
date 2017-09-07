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
		users: make(map[string]*User),
	}
	authenticator.users[TEST_USER] = &User{
		Username: TEST_USER,
		Password: TEST_PASSWORD,
	}
	tokenStore := NewInMemoryTokenStore(1)
	return authenticator, tokenStore, ""
}
