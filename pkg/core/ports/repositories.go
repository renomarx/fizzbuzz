package ports

import "github.com/renomarx/fizzbuzz/pkg/core/model"

// RequestsRepository repository to store requests counters
type RequestsRepository interface {
	// Inc increment requests counter by number for params
	Inc(params model.Params, number int) error
	// GetMaxStats get model.Stats for requests with the biggest counter
	GetMaxStats() (model.Stats, error)
}
