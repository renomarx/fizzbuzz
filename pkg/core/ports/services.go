package ports

import (
	"github.com/renomarx/fizzbuzz/pkg/core/model"
)

type FizzbuzzService interface {
	Fizzbuzz(params model.Params) []string
}
