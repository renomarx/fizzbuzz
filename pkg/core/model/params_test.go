package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParamsParse(t *testing.T) {
	p := Params{}
	err := p.Parse("toto", "5", "16", "fizz", "buzz")
	assert.NotNil(t, err)

	err = p.Parse("3", "toto", "16", "fizz", "buzz")
	assert.NotNil(t, err)

	err = p.Parse("3", "5", "toto", "fizz", "buzz")
	assert.NotNil(t, err)

	err = p.Parse("", "", "", "", "")
	assert.NotNil(t, err)

	err = p.Parse("3", "5", "16", "fizz", "buzz")
	assert.Nil(t, err)
	assert.Equal(t, Params{
		Int1:  3,
		Int2:  5,
		Limit: 16,
		Str1:  "fizz",
		Str2:  "buzz",
	}, p)
}

func TestParamsValidate(t *testing.T) {
	p := Params{
		Int1:  3,
		Int2:  5,
		Limit: 0,
		Str1:  "fizz",
		Str2:  "buzz",
	}
	err := p.Validate()
	assert.NotNil(t, err)

	p.Limit = 16
	err = p.Validate()
	assert.Nil(t, err)
}
