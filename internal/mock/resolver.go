package mock

import (
	"github.com/erlorenz/sparkflow"
)

type Resolver struct {
	Assets []sparkflow.Asset
	Error  error
}

func (r Resolver) Resolve(path string) ([]sparkflow.Asset, error) {
	if r.Error != nil {
		return nil, r.Error
	}

	return r.Assets, nil
}
