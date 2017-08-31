package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestData struct {
	A                string
	B                string
	ExpectedDistance int
}

func TestMin(t *testing.T) {
	assert.Equal(t, 1, min(1, 2, 3))
	assert.Equal(t, 3, min(5, 3, 8))
	assert.Equal(t, 10, min(12, 23, 10))
}

func TestDistance(t *testing.T) {
	test := []*TestData{
		&TestData{A: "hell", B: "hello", ExpectedDistance: 1},
		&TestData{A: "hell", B: "help", ExpectedDistance: 1},
		&TestData{A: "hell", B: "he", ExpectedDistance: 2},
	}

	for _, data := range test {
		distance := FindDistance(data.A, data.B)
		assert.Equal(t, data.ExpectedDistance, distance)
	}
}
