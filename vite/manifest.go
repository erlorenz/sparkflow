package vite

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/erlorenz/sparkflow"
)

// ParseManifest reads the manifest.json and sets it as the AssetMap.
func (sr *StaticResolver) ParseManifest() error {
	manifestPath := filepath.Join(sr.PublicDir, sr.BuildDir, sr.Manifest)

	fBytes, err := os.ReadFile(manifestPath)
	if err != nil {
		return fmt.Errorf("unable to read manifest: %s", err)
	}

	sr.AssetMap, err = manifestToAssetMap(fBytes)
	if err != nil {
		return err
	}

	return nil
}

func manifestToAssetMap(manifest []byte) (map[string]sparkflow.Asset, error) {

	chunkMap := map[string]ManifestChunk{}
	if err := json.Unmarshal(manifest, &chunkMap); err != nil {
		return nil, err
	}

	assetMap := map[string]sparkflow.Asset{}
	for file, chunk := range chunkMap {
		assetMap[file] = sparkflow.Asset{
			Filepath:    chunk.File,
			LogicalPath: chunk.Src,
			Ext:         filepath.Ext(chunk.File),
			IsChunk:     strings.HasPrefix(file, "_"),
			Imports:     chunk.Imports,
		}
	}

	return assetMap, nil
}
