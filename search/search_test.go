package search

import (
	"fmt"

	"github.com/sh3rp/databox/db"
	"github.com/sh3rp/databox/msg"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type SearchEngineTestSuite struct {
	suite.Suite
	NewSearchEngine func() (SearchEngine, string)
	SearchEngine    SearchEngine
	TearDown        func(string)
	ID              string
}

func (suite *SearchEngineTestSuite) SetupTest() {
	suite.SearchEngine, suite.ID = suite.NewSearchEngine()
}

func (suite *SearchEngineTestSuite) TearDownTest() {
	if suite.TearDown != nil {
		suite.TearDown(suite.ID)
	}
}

func (suite *SearchEngineTestSuite) TestIndexLink() {
	s := suite.SearchEngine

	link := &msg.Link{
		Id:   db.NewLinkKey(mockBox()),
		Url:  "http://www.cnn.com",
		Name: "CNN",
		Tags: []string{
			"fake",
			"news",
		},
	}

	err := s.Index(*link.Id, link.Tags)

	assert.Nil(suite.T(), err)

	assert.Equal(suite.T(), len(s.Find("fake", 10, 0)), 1)
}

func (suite *SearchEngineTestSuite) TestIndexLinkUpdate() {
	s := suite.SearchEngine

	link := &msg.Link{
		Id:   db.NewLinkKey(mockBox()),
		Url:  "http://www.cnn.com",
		Name: "CNN",
		Tags: []string{
			"fake",
			"news",
		},
	}

	err := s.Index(*link.Id, link.Tags)

	assert.Nil(suite.T(), err)

	assert.Equal(suite.T(), len(s.Find("fake", 10, 0)), 1)

	link.Tags = []string{"real", "news"}

	err = s.Index(*link.Id, link.Tags)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 1, len(s.Find("real", 10, 0)))
	assert.Equal(suite.T(), 0, len(s.Find("fake", 10, 0)))
}

func (suite *SearchEngineTestSuite) TestIndexFindLink() {
	s := suite.SearchEngine

	for i := 0; i < 1000; i++ {
		link := &msg.Link{
			Id:   db.NewLinkKey(mockBox()),
			Url:  fmt.Sprintf("http://www.cnn%d.com", i),
			Name: fmt.Sprintf("Name%d", i),
		}
		if i%100 == 0 {
			link.Tags = []string{"search", "term", "pants"}
		} else {
			link.Tags = []string{"term", "shoes", "storm"}
		}
		s.Index(*link.Id, link.Tags)
	}

	links := s.Find("search", 10, 0)

	assert.Equal(suite.T(), len(links), 10)
}

func (suite *SearchEngineTestSuite) TestIndexFindLinkDups() {
	s := suite.SearchEngine
	tags1 := []string{"fake", "news"}
	tags2 := []string{"real", "fake", "news"}
	tags3 := []string{"real", "fake", "news"}
	key := db.NewLinkKey(mockBox())
	err := s.Index(*key, tags1)
	assert.Nil(suite.T(), err)
	err = s.Index(*key, tags2)
	assert.Nil(suite.T(), err)
	err = s.Index(*key, tags3)

	keys := s.Find("fake", 10, 0)
	assert.Equal(suite.T(), 1, len(keys))
	keys = s.Find("news", 10, 0)
	assert.Equal(suite.T(), 1, len(keys))
	keys = s.Find("real", 10, 0)
	assert.Equal(suite.T(), 1, len(keys))
}

func mockBox() *msg.Key {
	return &msg.Key{
		Id:   "mockboxid",
		Type: msg.Key_BOX,
	}
}
