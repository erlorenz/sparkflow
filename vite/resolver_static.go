package vite

import (
	"fmt"
	"slices"

	"github.com/erlorenz/sparkflow"
)

const (
	defaultPublicDir    = "dist"
	defaultBuildDir     = "assets"
	defaultManifestPath = ".vite/manifest.json"
)

// StaticResolver resolves assets from the vite manifest.json.
type StaticResolver struct {
	Manifest  string
	PublicDir string
	BuildDir  string
	AssetMap  map[string]sparkflow.Asset
}

// A vite manifest chunk.
type ManifestChunk struct {
	File    string   `json:"file"`
	Src     string   `json:"src"`
	IsEntry bool     `json:"isEntry"`
	Imports []string `json:"imports"`
}

// Resolve implements sparkflow.Resolver. Recursively resolves assets.
func (sr *StaticResolver) Resolve(logicalPath string) ([]sparkflow.Asset, error) {
	asset, ok := sr.AssetMap[logicalPath]
	if !ok {
		return nil, fmt.Errorf("cannot resolve %s: %w", logicalPath, sparkflow.ErrNotFound)
	}

	allAssets := []sparkflow.Asset{asset}

	for _, imp := range asset.Imports {
		impAssets, err := sr.Resolve(imp)
		if err != nil {
			return nil, err
		}
		allAssets = slices.Concat(allAssets, impAssets)

	}
	return allAssets, nil
}
