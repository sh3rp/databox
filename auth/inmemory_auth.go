package auth

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"math/rand"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/sh3rp/databox/msg"
	"golang.org/x/crypto/bcrypt"
)

// InMemoryAuthenticator - base authenticator struct for Authenticator
type InMemoryAuthenticator struct {
	users map[string]*User
}

// NewInMemoryAuthenticator - ctor
func NewInMemoryAuthenticator() Authenticator {
	return &InMemoryAuthenticator{
		users: make(map[string]*User),
	}
}

// Authenticate - implementation for Authenticate contract
func (a *InMemoryAuthenticator) Authenticate(username string, pass []byte) bool {
	if user, ok := a.users[username]; ok {
		if bcrypt.CompareHashAndPassword(user.Password, pass) == nil {
			return true
		}
	}
	return false
}

// AddUser - adds a user with a password
func (a *InMemoryAuthenticator) AddUser(user string, pass []byte) error {
	hashedPassword, err := bcrypt.GenerateFromPassword(pass, 10)

	if err != nil {
		return err
	}

	a.users[user] = &User{
		Username: user,
		Password: hashedPassword,
	}
	return nil
}

// DeleteUser - deletes a user
func (a *InMemoryAuthenticator) DeleteUser(username string) error {
	if _, ok := a.users[username]; !ok {
		return errors.New(ERR_AUTH_NO_USER)
	}
	delete(a.users, username)
	return nil
}

type InMemoryTokenStore struct {
	tokens                  map[string]*msg.Token
	expirationTimeInSeconds int64
}

func NewInMemoryTokenStore(expTimeInSeconds int64) TokenStore {
	return &InMemoryTokenStore{
		tokens:                  make(map[string]*msg.Token),
		expirationTimeInSeconds: expTimeInSeconds,
	}
}

func (ts *InMemoryTokenStore) GenerateToken(user string) *msg.Token {
	hasher := sha256.New()
	hasher.Write([]byte(user))
	randBytes := make([]byte, 32)
	for i := 0; i < len(randBytes); i++ {
		randBytes[i] = byte(rand.Int())
	}
	hasher.Write(randBytes)
	hash := hasher.Sum(nil)

	token := &msg.Token{
		Username:       user,
		TokenHash:      hex.EncodeToString(hash),
		ExpirationTime: ts.getCurrentExpiration(),
	}

	ts.tokens[user] = token
	return token
}

func (ts *InMemoryTokenStore) ValidateToken(token *msg.Token) error {
	if token == nil {
		log.Error().Msgf("ValidateToken: no token passed")
		return errors.New("No token passed")
	}
	if _, ok := ts.tokens[token.Username]; ok {
		if ts.tokens[token.Username].ExpirationTime < time.Now().UnixNano() {
			delete(ts.tokens, token.Username)
			return errors.New(ERR_VALIDATION_EXPIRE)
		}
		if ts.tokens[token.Username].TokenHash == token.TokenHash {
			token.ExpirationTime = ts.getCurrentExpiration()
			ts.tokens[token.Username] = token
			return nil
		} else {
			return errors.New(ERR_VALIDATION_TOKEN)
		}
	}
	return errors.New(ERR_VALIDATION_USER)
}

func (ts *InMemoryTokenStore) getCurrentExpiration() int64 {
	return time.Now().Add(time.Duration(ts.expirationTimeInSeconds) * time.Second).UnixNano()
}
