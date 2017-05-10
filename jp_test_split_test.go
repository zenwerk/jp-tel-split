package jptelsplit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJpTelSplit(t *testing.T) {
	assert := assert.New(t)

	// valid telephone number
	result, err := JpTelSplit("0312345678")
	assert.Nil(err)
	assert.Equal("03", result[0])
	assert.Equal("1234", result[1])
	assert.Equal("5678", result[2])

	result, err = JpTelSplit("0222522222")
	assert.Nil(err)
	assert.Equal("022", result[0])
	assert.Equal("252", result[1])
	assert.Equal("2222", result[2])

	result, err = JpTelSplit("0997123456")
	assert.Nil(err)
	assert.Equal("0997", result[0])
	assert.Equal("12", result[1])
	assert.Equal("3456", result[2])

	result, err = JpTelSplit("0996912345")
	assert.Nil(err)
	assert.Equal("09969", result[0])
	assert.Equal("1", result[1])
	assert.Equal("2345", result[2])

	result, err = JpTelSplit("09012345678")
	assert.Nil(err)
	assert.Equal("090", result[0])
	assert.Equal("1234", result[1])
	assert.Equal("5678", result[2])

	result, err = JpTelSplit("0120123456")
	assert.Nil(err)
	assert.Equal("0120", result[0])
	assert.Equal("123", result[1])
	assert.Equal("456", result[2])

	result, err = JpTelSplit("08001234567")
	assert.Nil(err)
	assert.Equal("0800", result[0])
	assert.Equal("123", result[1])
	assert.Equal("4567", result[2])

	// invalid telephone number
	result, err = JpTelSplit("00000000000")
	assert.Error(err)
	assert.Equal("", result[0])
	assert.Equal("", result[1])
	assert.Equal("", result[2])

	result, err = JpTelSplit("0312")
	assert.Error(err)
	assert.Equal("", result[0])
	assert.Equal("", result[1])
	assert.Equal("", result[2])

	result, err = JpTelSplit("090")
	assert.Error(err)
	assert.Equal("", result[0])
	assert.Equal("", result[1])
	assert.Equal("", result[2])

	result, err = JpTelSplit("hogefugapiyo")
	assert.Error(err)
	assert.Equal("", result[0])
	assert.Equal("", result[1])
	assert.Equal("", result[2])

	result, err = JpTelSplit("あいうえお")
	assert.Error(err)
	assert.Equal("", result[0])
	assert.Equal("", result[1])
	assert.Equal("", result[2])
}