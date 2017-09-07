package auth

import (
	"time"

	"github.com/sh3rp/databox/msg"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

var TEST_USER = "test"
var TEST_PASSWORD = "password"

type AuthTestSuite struct {
	suite.Suite
	New        func() (Authenticator, TokenStore, string)
	Auth       Authenticator
	TokenStore TokenStore
	TearDown   func(string)
	ID         string
}

func (suite *AuthTestSuite) SetupTest() {
	suite.Auth, suite.TokenStore, suite.ID = suite.New()
}

func (suite *AuthTestSuite) TearDownTest() {
	if suite.TearDown != nil {
		suite.TearDown(suite.ID)
	}
}

func (suite *AuthTestSuite) TestAuthenticateSuccess() {
	authed := suite.Auth.Authenticate(TEST_USER, TEST_PASSWORD)
	assert.True(suite.T(), authed)
}

func (suite *AuthTestSuite) TestAuthenticateFailure() {
	authed := suite.Auth.Authenticate(TEST_USER, "badpassword")
	assert.False(suite.T(), authed)
}

func (suite *AuthTestSuite) TestAuthenticateEmptyFields() {
	authed := suite.Auth.Authenticate("", TEST_PASSWORD)
	assert.False(suite.T(), authed)
	authed = suite.Auth.Authenticate(TEST_USER, "")
	assert.False(suite.T(), authed)
}

func (suite *AuthTestSuite) TestAddUserSuccess() {
	err := suite.Auth.AddUser(TEST_USER, TEST_PASSWORD)
	assert.Nil(suite.T(), err)
	authed := suite.Auth.Authenticate(TEST_USER, TEST_PASSWORD)
	assert.True(suite.T(), authed)
}

func (suite *AuthTestSuite) TestDeleteUserSuccess() {
	err := suite.Auth.AddUser(TEST_USER, TEST_PASSWORD)
	assert.Nil(suite.T(), err)
	err = suite.Auth.DeleteUser(TEST_USER)
	assert.Nil(suite.T(), err)
}

func (suite *AuthTestSuite) TestGenerateToken() {
	token := suite.TokenStore.GenerateToken(TEST_USER, int64(time.Now().Add(1*time.Minute).UnixNano()))

	assert.Equal(suite.T(), TEST_USER, token.Username)
	assert.True(suite.T(), token.ExpirationTime > (time.Now().UnixNano()))
	assert.NotNil(suite.T(), token.TokenHash)
}

func (suite *AuthTestSuite) TestValidateTokenSuccess() {
	token := suite.TokenStore.GenerateToken(TEST_USER, int64(time.Now().Add(1*time.Minute).UnixNano()))

	validationError := suite.TokenStore.ValidateToken(token)
	assert.Nil(suite.T(), validationError)
}

func (suite *AuthTestSuite) TestValidateTokenFailureExpiration() {
	token := suite.TokenStore.GenerateToken(TEST_USER, time.Now().Add(1*time.Second).UnixNano())
	time.Sleep(1 * time.Second)
	validationError := suite.TokenStore.ValidateToken(token)
	assert.NotNil(suite.T(), validationError)
	assert.Equal(suite.T(), ERR_VALIDATION_EXPIRE, validationError.Error())
}

func (suite *AuthTestSuite) TestValidateTokenFailureUser() {
	token := &msg.Token{
		Username:       "baduser",
		TokenHash:      "asdfas",
		ExpirationTime: time.Now().UnixNano(),
	}
	validationError := suite.TokenStore.ValidateToken(token)
	assert.NotNil(suite.T(), validationError)
	assert.Equal(suite.T(), ERR_VALIDATION_USER, validationError.Error())
}

func (suite *AuthTestSuite) TestValidateTokenFailureToken() {
	suite.TokenStore.GenerateToken(TEST_USER, time.Now().Add(1*time.Second).UnixNano())
	badToken := &msg.Token{
		Username:       TEST_USER,
		TokenHash:      "asdfas",
		ExpirationTime: time.Now().UnixNano(),
	}
	validationError := suite.TokenStore.ValidateToken(badToken)
	assert.NotNil(suite.T(), validationError)
	assert.Equal(suite.T(), ERR_VALIDATION_TOKEN, validationError.Error())
}
