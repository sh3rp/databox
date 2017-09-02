package db

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

// func TestNewBoltDB(t *testing.T) {
// 	db := NewBoltDB("/tmp/test.db")

// 	assert.NotNil(t, db)
// 	err := db.DeleteBox("none")
// 	assert.NotNil(t, err)
// }

func TestBoltDB(t *testing.T) {
	suite.Run(t, &DBTestSuite{
		NewDB:    getTestDB,
		TearDown: tearDown,
	})
}

func tearDown(id string) {
	os.RemoveAll("/tmp/test-" + id + ".db")
}

func getTestDB() (BoxDB, string) {
	id := GenerateID()
	return NewBoltDB("/tmp/test-" + id + ".db"), id
}
