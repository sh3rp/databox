package db

import (
	"fmt"

	"github.com/sh3rp/databox/msg"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type DBTestSuite struct {
	suite.Suite
	NewDB    func() (BoxDB, string)
	TearDown func(string)
	DB       BoxDB
	ID       string
}

func (suite *DBTestSuite) SetupTest() {
	db, id := suite.NewDB()
	suite.DB = db
	suite.ID = id
}

func (suite *DBTestSuite) TearDownTest() {
	if suite.TearDown != nil {
		suite.TearDown(suite.ID)
	}
}

func (suite *DBTestSuite) TestNewBox() {
	name := "test"
	descr := "test description"

	box, err := suite.DB.NewBox(name, descr)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), box.Name, name)
	assert.Equal(suite.T(), box.Description, descr)

	box, err = suite.DB.NewBox(name, descr)

	assert.Nil(suite.T(), err)

	box, err = suite.DB.NewBox("", descr)
	assert.NotNil(suite.T(), err)

	box, err = suite.DB.NewBox(name, "")
	assert.NotNil(suite.T(), err)
}

func (suite *DBTestSuite) TestSaveBox() {
	name := "test"
	descr := "test description"

	box, err := suite.DB.NewBox(name, descr)
	box.Name = "new name"
	err = suite.DB.SaveBox(box)

	assert.Nil(suite.T(), err)

	box, err = suite.DB.GetBoxById(*box.Id)
	assert.Equal(suite.T(), box.Name, "new name")

	badBox := &msg.Box{
		Name:        "blah",
		Description: "blah blah",
	}

	err = suite.DB.SaveBox(badBox)
	assert.NotNil(suite.T(), err)
}

func (suite *DBTestSuite) TestGetBoxById() {
	name := "test"
	descr := "test description"

	box, err := suite.DB.NewBox(name, descr)
	assert.Nil(suite.T(), err)

	newBox, err := suite.DB.GetBoxById(*box.Id)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), newBox.Name, box.Name)
	assert.Equal(suite.T(), newBox.Description, box.Description)
}

func (suite *DBTestSuite) TestGetBoxes() {
	name := "test"
	descr := "test description"

	boxes := make(map[msg.Key]*msg.Box)

	for i := 0; i < 10; i++ {
		b, err := suite.DB.NewBox(name+fmt.Sprintf("%d", i), descr+fmt.Sprintf("%d", i))
		assert.Nil(suite.T(), err)
		boxes[*b.Id] = b
	}

	newBoxes, err := suite.DB.GetBoxes()

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), len(newBoxes), len(boxes))
	for i := 0; i < 10; i++ {
		assert.Equal(suite.T(), boxes[*newBoxes[i].Id].Name, newBoxes[i].Name)
		assert.Equal(suite.T(), boxes[*newBoxes[i].Id].Description, newBoxes[i].Description)
	}
}

func (suite *DBTestSuite) TestDeleteBox() {
	name := "test"
	descr := "test description"

	box, _ := suite.DB.NewBox(name, descr)
	err := suite.DB.DeleteBox(*box.Id)

	assert.Nil(suite.T(), err)

	_, err = suite.DB.GetBoxById(*box.Id)

	assert.NotNil(suite.T(), err)
}

func (suite *DBTestSuite) TestNewLink() {
	box, _ := suite.DB.NewBox("test", "test description")

	link, err := suite.DB.NewLink("testlink", "http://www.cnn.com", *box.Id)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), link.Name, "testlink")
	assert.Equal(suite.T(), link.Url, "http://www.cnn.com")

	_, err = suite.DB.NewLink("badlink", "http://www.pants.com", msg.Key{Id: "", Type: msg.Key_BOX})
	assert.NotNil(suite.T(), err)

	_, err = suite.DB.NewLink("badlink", "http://www.pants.com", msg.Key{Id: "aasdf", Type: msg.Key_BOX})
	assert.NotNil(suite.T(), err)
}

func (suite *DBTestSuite) TestSaveLink() {
	box, _ := suite.DB.NewBox("test", "test description")
	link, _ := suite.DB.NewLink("testlink", "http://www.cnn.com", *box.Id)
	link.Url = "http://www.msnbc.com"
	err := suite.DB.SaveLink(link)
	assert.Nil(suite.T(), err)
}

func (suite *DBTestSuite) TestGetLinkById() {
	box, _ := suite.DB.NewBox("test", "test description")
	link, _ := suite.DB.NewLink("testlink", "http://www.cnn.com", *box.Id)

	newLink, err := suite.DB.GetLinkById(*link.Id)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), link.Name, newLink.Name)
	assert.Equal(suite.T(), link.Url, newLink.Url)
}

func (suite *DBTestSuite) TestGetLinksByBoxId() {
	box1, _ := suite.DB.NewBox("test", "test description")
	box2, _ := suite.DB.NewBox("test2", "blah blah")

	for i := 0; i < 10; i++ {
		suite.DB.NewLink(fmt.Sprintf("test%d", i), "http://www.cnn.com", *box1.Id)
	}

	for i := 0; i < 5; i++ {
		suite.DB.NewLink(fmt.Sprintf("test%d-2", i), "http://www.msnbc.com", *box2.Id)
	}

	links1, err := suite.DB.GetLinksByBoxId(*box1.Id)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 10, len(links1))

	links2, err := suite.DB.GetLinksByBoxId(*box2.Id)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 5, len(links2))

	_, err = suite.DB.GetLinksByBoxId(msg.Key{Id: "asdf", Type: msg.Key_BOX})
	assert.NotNil(suite.T(), err)
}
func (suite *DBTestSuite) TestDeleteLink() {
	assert.True(suite.T(), true)
}
