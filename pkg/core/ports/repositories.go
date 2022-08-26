package ports

import "github.com/renomarx/fizzbuzz/pkg/core/model"

// RequestsRepository repository to store requests counters
type RequestsRepository interface {
	Inc(params model.Params, number int) error
	GetMaxStats() (model.Stats, error)
}
