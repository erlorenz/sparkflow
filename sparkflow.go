package sparkflow

import (
	"errors"
)

// Asset represents the
type Asset struct {
	Filepath    string   // True filepath
	Ext         string   // File extension
	IsChunk     bool     // If it is a chunk that should be preloaded
	LogicalPath string   // The path that is requested by the view
	Imports     []string // The list of dependencies
}

// Resolver represents the mechanism for retrieving
// assets and their dependencies.
type Resolver interface {
	Resolve(string) ([]Asset, error)
}

var ErrInvalidExt = errors.New("invalid extension")
var ErrNotFound = errors.New("not found")
