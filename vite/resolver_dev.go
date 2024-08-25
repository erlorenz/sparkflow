package vite

import (
	"cmp"
	"net/url"

	"github.com/erlorenz/sparkflow"
)

const defaultDevURL = "http://localhost:5173"

// DevResolver resolves paths to the vite dev server.
// Defaults to :5173.
type DevResolver struct {
	URL string // The vite server URL
}

// Resolve implements sparkflow.Resolver.
func (dr *DevResolver) Resolve(path string) ([]sparkflow.Asset, error) {
	dr.URL = cmp.Or(dr.URL, defaultDevURL)

	serverURL, err := url.Parse(dr.URL)
	if err != nil {
		return nil, err
	}

	ext, err := parseExt(path)
	if err != nil {
		return nil, err
	}

	viteURL := serverURL.JoinPath(path).String()

	return []sparkflow.Asset{{
		Filepath:    viteURL,
		LogicalPath: path,
		Ext:         ext,
	}}, nil
}
