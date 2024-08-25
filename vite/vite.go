package vite

import (
	"cmp"

	"github.com/erlorenz/sparkflow"
)

// Vite wraps either a DevResolver or StaticResolver and implements sparkflow.Resolver.
// It also provides a script for hot-module-reload to be injected into the page.
type Vite struct {
	sparkflow.Resolver
	Environment string // Environment (development|production).
	DevURL      string // The base URL of the vite dev server.
}

// Config configures the vite resolvers.
type Config struct {

	// The environment (development|production).
	// Controls whether to use static or dev resolver. Defaults to production.
	Environment string

	// Dev server URL, defaults to http://localhost:5173.
	DevURL string

	// Public directory, defaults to "dist".
	PublicDir string

	// Build directory inside public directory, defaults to "assets".
	BuildDir string

	// Manifest file path inside build directory, defaults to ".vite/manifest.json".
	Manifest string
}

// New chooses either the DevResolver or the StaticResolver.
// It calls ParseManifest on the StaticResolver if environment = "production".
func New(config Config) (*Vite, error) {

	vite := &Vite{
		Environment: parseEnvironment(config.Environment),
	}

	// Dev server
	if config.Environment == "development" {
		vite.DevURL = config.DevURL
		vite.Resolver = &DevResolver{config.DevURL}
		return vite, nil
	}

	// Set defaults
	static := &StaticResolver{
		Manifest:  cmp.Or(config.Manifest, defaultManifestPath),
		PublicDir: cmp.Or(config.PublicDir, defaultPublicDir),
		BuildDir:  cmp.Or(config.BuildDir, defaultBuildDir),
	}

	if err := static.ParseManifest(); err != nil {
		return &Vite{}, err
	}

	vite.Resolver = static
	return vite, nil
}
