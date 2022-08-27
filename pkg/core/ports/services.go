package ports

import (
	"github.com/renomarx/fizzbuzz/pkg/core/model"
)

// FizzbuzzService service implementing fizzbuzz agorithm
type FizzbuzzService interface {
	// Fizzbuzz fizzbuzz params to a []string
	Fizzbuzz(params model.Params) []string
}
