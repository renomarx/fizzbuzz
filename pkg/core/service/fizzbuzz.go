package service

import (
	"strconv"

	"github.com/renomarx/fizzbuzz/pkg/core/model"
)

type fizzbuzzSVC struct {
}

func NewFizzbuzzSVC() *fizzbuzzSVC {
	return &fizzbuzzSVC{}
}

func (svc *fizzbuzzSVC) Fizzbuzz(params model.Params) []string {
	res := make([]string, params.Limit)
	for i := 1; i <= params.Limit; i++ {
		res[i-1] = svc.getFizzbuzzString(i, params)
	}
	return res
}

func (svc *fizzbuzzSVC) getFizzbuzzString(i int, params model.Params) string {
	str := ""
	if i%params.Int1 == 0 {
		str = params.Str1
	}
	if i%params.Int2 == 0 {
		str += params.Str2
	}
	if str != "" {
		return str
	}
	return strconv.Itoa(i)
}
