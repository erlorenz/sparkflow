package mock

import (
	"cmp"
	"encoding/json"
	"path/filepath"
	"strings"
)

type chunk struct {
	File    string   `json:"file"`
	Src     string   `json:"src"`
	IsEntry bool     `json:"isEntry"`
	Imports []string `json:"imports"`
}

type FakeManifest struct {
	ChunkMap map[string]chunk
}

func (fm *FakeManifest) AddChunk(file, src string, imports ...string) {
	if fm.ChunkMap == nil {
		fm.ChunkMap = map[string]chunk{}
	}

	baseFile := filepath.Base(file)
	if strings.HasPrefix(baseFile, "_") {
		fm.ChunkMap[baseFile] = chunk{File: file}
		return
	}

	fm.ChunkMap[cmp.Or(src, file)] = chunk{
		File:    file,
		Src:     src, // shared chunks are empty here
		Imports: imports,
		IsEntry: true,
	}
}

func (fm *FakeManifest) MustJSON() []byte {
	b, err := json.Marshal(fm.ChunkMap)
	if err != nil {
		panic(err)
	}

	return b
}
