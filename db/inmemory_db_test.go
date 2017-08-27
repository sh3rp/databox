package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBox(t *testing.T) {
	db := NewInMemoryDB()
	name := "test"
	descr := "test description"

	box, err := db.NewBox(name, descr, false)

	assert.Nil(t, err)
	assert.Equal(t, box.Name, name)
	assert.Equal(t, box.Description, descr)
	assert.False(t, box.IsDefault)
}
