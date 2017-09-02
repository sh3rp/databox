package db

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestInMemoryDB(t *testing.T) {
	suite.Run(t, &DBTestSuite{
		NewDB: newDB,
	})
}

func newDB() (BoxDB, string) {
	return NewInMemoryDB(), ""
}
