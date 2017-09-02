package search

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestInMemorySearchEngine(t *testing.T) {
	suite.Run(t, &SearchEngineTestSuite{
		NewSearchEngine: getInMemorySearchEngine,
	})
}

func getInMemorySearchEngine() (SearchEngine, string) {
	return NewInMemorySearchEngine(), ""
}
