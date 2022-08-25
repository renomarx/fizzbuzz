package model

import (
	"fmt"
	"strconv"

	"github.com/sirupsen/logrus"
)

type Params struct {
	Int1  int
	Int2  int
	Limit int
	Str1  string
	Str2  string
}

// Parse parse params from strings, return the first error encountered
func (params *Params) Parse(int1, int2, limit, str1, str2 string) error {
	iint1, err := strconv.Atoi(int1)
	if err != nil {
		logrus.Error(err)
		return fmt.Errorf("Impossible to parse to int int1 %s", int1)
	}
	params.Int1 = iint1

	iint2, err := strconv.Atoi(int2)
	if err != nil {
		logrus.Error(err)
		return fmt.Errorf("Impossible to parse to int int2 %s", int2)
	}
	params.Int2 = iint2

	ilimit, err := strconv.Atoi(limit)
	if err != nil {
		logrus.Error(err)
		return fmt.Errorf("Impossible to parse to int limit %s", limit)
	}
	params.Limit = ilimit

	params.Str1 = str1
	params.Str2 = str2
	return nil
}

// Validate validate params format (not really necessary here but anyway)
func (params *Params) Validate() error {
	if params.Limit < 1 {
		return fmt.Errorf("Limit cannot be lower than 1")
	}
	return nil
}
