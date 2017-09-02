package search

import (
	"os"
	"testing"

	"github.com/sh3rp/databox/db"
	"github.com/stretchr/testify/suite"
)

func TestBoltSearchEngine(t *testing.T) {
	suite.Run(t, &SearchEngineTestSuite{
		NewSearchEngine: getBoltSearchEngine,
		TearDown:        tearDown,
	})
}
func tearDown(id string) {
	os.RemoveAll("/tmp/test-" + id + ".db")
}

func getBoltSearchEngine() (SearchEngine, string) {
	id := db.GenerateID()
	return NewBoltSearchEngine("/tmp/test-" + id + ".db"), id
}
