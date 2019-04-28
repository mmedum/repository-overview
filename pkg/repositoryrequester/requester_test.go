package repositoryrequester

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIdentificationOfTestBCRegexp(t *testing.T) {
	input := "TestBC.Service"
	expected := "TestBC"
	actual := identifyBoundedContext(input)

	assert.Equal(t, expected, actual, "they should be equal")
}

func TestIdentificationOfUndefinedBoundedContext(t *testing.T) {
	input := "Service"
	expected := "Undefined"
	actual := identifyBoundedContext(input)

	assert.Equal(t, expected, actual, "they should be equal")
}

func TestIdentificationOfEmptyStringBoundedContext(t *testing.T) {
	input := ""
	expected := "Undefined"
	actual := identifyBoundedContext(input)

	assert.Equal(t, expected, actual, "they should be equal")
}

func TestIdentificationOfUpperCaseBoundedContext(t *testing.T) {
	input := "ABEBC.Service"
	expected := "ABEBC"
	actual := identifyBoundedContext(input)

	assert.Equal(t, expected, actual, "they should be equal")
}

func TestIdentificationOfLowerCaseBoundedContext(t *testing.T) {
	input := "abeBC.Service"
	expected := "abeBC"
	actual := identifyBoundedContext(input)

	assert.Equal(t, expected, actual, "they should be equal")
}

func TestIdentificationOfBoundedContext(t *testing.T) {
	input := "BCBC.Service"
	expected := "BCBC"
	actual := identifyBoundedContext(input)

	assert.Equal(t, expected, actual, "they should be equal")
}
