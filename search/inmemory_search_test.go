package search

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestInMemorySearchEngine(t *testing.T) {
	suite.Run(t, &SearchEngineTestSuite{NewSearchEngine: NewInMemorySearchEngine})
}
