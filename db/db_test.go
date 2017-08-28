package db

import (
	"fmt"

	"github.com/sh3rp/databox/msg"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type DBTestSuite struct {
	suite.Suite
	NewDB func() BoxDB
	DB    BoxDB
}

func (suite *DBTestSuite) SetupTest() {
	suite.DB = suite.NewDB()
}

func (suite *DBTestSuite) TestNewBox() {
	name := "test"
	descr := "test description"

	box, err := suite.DB.NewBox(name, descr, false)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), box.Name, name)
	assert.Equal(suite.T(), box.Description, descr)
	assert.False(suite.T(), box.IsDefault)

	box, err = suite.DB.NewBox(name, descr, true)

	assert.Nil(suite.T(), err)
	assert.True(suite.T(), box.IsDefault)

	box, err = suite.DB.NewBox("", descr, false)
	assert.NotNil(suite.T(), err)

	box, err = suite.DB.NewBox(name, "", false)
	assert.NotNil(suite.T(), err)
}

func (suite *DBTestSuite) TestSaveBox() {
	name := "test"
	descr := "test description"

	box, err := suite.DB.NewBox(name, descr, true)
	box.Name = "new name"
	err = suite.DB.SaveBox(box)

	assert.Nil(suite.T(), err)

	box, err = suite.DB.GetBoxById(box.Id)
	assert.Equal(suite.T(), box.Name, "new name")

	badBox := &msg.Box{
		Name:        "blah",
		Description: "blah blah",
		IsDefault:   false,
	}

	err = suite.DB.SaveBox(badBox)
	assert.NotNil(suite.T(), err)
}

func (suite *DBTestSuite) TestGetBoxById() {
	name := "test"
	descr := "test description"

	box, err := suite.DB.NewBox(name, descr, true)

	newBox, err := suite.DB.GetBoxById(box.Id)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), newBox.Name, box.Name)
	assert.Equal(suite.T(), newBox.Description, box.Description)
}

func (suite *DBTestSuite) TestGetBoxes() {
	name := "test"
	descr := "test description"

	boxes := make(map[string]*msg.Box)

	for i := 0; i < 10; i++ {
		b, _ := suite.DB.NewBox(name+fmt.Sprintf("%d", i), descr+fmt.Sprintf("%d", i), false)
		boxes[b.Id] = b
	}

	newBoxes, err := suite.DB.GetBoxes()

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), len(newBoxes), len(boxes))
	for i := 0; i < 10; i++ {
		assert.Equal(suite.T(), boxes[newBoxes[i].Id].Name, newBoxes[i].Name)
		assert.Equal(suite.T(), boxes[newBoxes[i].Id].Description, newBoxes[i].Description)
	}
}

func (suite *DBTestSuite) TestDeleteBox() {
	name := "test"
	descr := "test description"

	box, _ := suite.DB.NewBox(name, descr, false)
	err := suite.DB.DeleteBox(box.Id)

	assert.Nil(suite.T(), err)

	_, err = suite.DB.GetBoxById(box.Id)

	assert.NotNil(suite.T(), err)
}

func (suite *DBTestSuite) TestGetDefaultBox() {

	box, _ := suite.DB.NewBox("test", "test description", true)
	newBox, _ := suite.DB.NewBox("blah", "blah blah", true)

	assert.True(suite.T(), newBox.IsDefault)

	b, _ := suite.DB.GetDefaultBox()
	assert.True(suite.T(), b.IsDefault)
	newBox, _ = suite.DB.GetBoxById(newBox.Id)
	assert.True(suite.T(), newBox.IsDefault)
	box, _ = suite.DB.GetBoxById(box.Id)
	assert.False(suite.T(), box.IsDefault)

}

func (suite *DBTestSuite) TestNewLink() {
	box, _ := suite.DB.NewBox("test", "test description", true)

	link, err := suite.DB.NewLink("testlink", "http://www.cnn.com", box.Id)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), link.Name, "testlink")
	assert.Equal(suite.T(), link.Url, "http://www.cnn.com")

	_, err = suite.DB.NewLink("badlink", "http://www.pants.com", "")
	assert.NotNil(suite.T(), err)
}

func (suite *DBTestSuite) TestSaveLink() {
	box, _ := suite.DB.NewBox("test", "test description", true)
	link, _ := suite.DB.NewLink("testlink", "http://www.cnn.com", box.Id)
	link.Url = "http://www.msnbc.com"
	err := suite.DB.SaveLink(link)
	assert.Nil(suite.T(), err)
}

func (suite *DBTestSuite) TestGetLinkById() {
	box, _ := suite.DB.NewBox("test", "test description", true)
	link, _ := suite.DB.NewLink("testlink", "http://www.cnn.com", box.Id)

	newLink, err := suite.DB.GetLinkById(link.Id)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), link.Name, newLink.Name)
	assert.Equal(suite.T(), link.Url, newLink.Url)
}

func (suite *DBTestSuite) TestGetLinks() {
	box, _ := suite.DB.NewBox("test", "test description", true)
	for i := 0; i < 10; i++ {
		suite.DB.NewLink("testlink"+fmt.Sprintf("%d", i), "http://www.cnn"+fmt.Sprintf("%d", i)+".com", box.Id)
	}

	links, err := suite.DB.GetLinks()

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), len(links), 10)
}

func (suite *DBTestSuite) TestGetLinksByBoxId() {}
func (suite *DBTestSuite) TestDeleteLink()      {}
