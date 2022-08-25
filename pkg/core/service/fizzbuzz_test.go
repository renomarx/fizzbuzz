package service

import (
	"testing"

	"github.com/renomarx/fizzbuzz/pkg/core/model"
	"github.com/stretchr/testify/assert"
)

func TestFizzbuzzStandard(t *testing.T) {
	svc := NewFizzbuzzSVC()
	res := svc.Fizzbuzz(model.Params{
		Int1:  3,
		Int2:  5,
		Limit: 20,
		Str1:  "fizz",
		Str2:  "buzz",
	})
	expected := []string{"1", "2", "fizz", "4", "buzz", "fizz", "7", "8", "fizz", "buzz", "11", "fizz", "13", "14", "fizzbuzz", "16", "17", "fizz", "19", "buzz"}
	assert.Equal(t, expected, res)
}

func TestFizzbuzzNoTonicWithoutGin(t *testing.T) {
	svc := NewFizzbuzzSVC()
	// Int2 is a multiple of Int1
	// Also testing different strings than fizz and buzz, just to be sure
	res := svc.Fizzbuzz(model.Params{
		Int1:  3,
		Int2:  6,
		Limit: 20,
		Str1:  "gin",
		Str2:  "tonic",
	})
	expected := []string{"1", "2", "gin", "4", "5", "gintonic", "7", "8", "gin", "10", "11", "gintonic", "13", "14", "gin", "16", "17", "gintonic", "19", "20"}
	assert.Equal(t, expected, res)
}
