package db

import (
	"fmt"

	"github.com/sh3rp/databox/msg"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type SearchEngineTestSuite struct {
	suite.Suite
	NewSearchEngine func() SearchEngine
	SearchEngine    SearchEngine
}

func (suite *SearchEngineTestSuite) SetupTest() {
	suite.SearchEngine = suite.NewSearchEngine()
}

func (suite *SearchEngineTestSuite) TestIndexLink() {
	s := suite.SearchEngine

	link := &msg.Link{
		Id:   GenerateID(),
		Url:  "http://www.cnn.com",
		Name: "CNN",
		Tags: []string{
			"fake",
			"news",
		},
	}

	err := s.IndexLink(link)

	assert.Nil(suite.T(), err)
}

func (suite *SearchEngineTestSuite) TestIndexFindLink() {
	s := suite.SearchEngine

	for i := 0; i < 1000; i++ {
		link := &msg.Link{
			Id:   GenerateID(),
			Url:  fmt.Sprintf("http://www.cnn%d.com", i),
			Name: fmt.Sprintf("Name%d", i),
		}
		if i%10 == 0 {
			link.Tags = []string{"search", "term", "pants"}
		}
		s.IndexLink(link)
	}

	links := s.FindLinks("search")

	assert.Equal(suite.T(), len(links), 100)
}
