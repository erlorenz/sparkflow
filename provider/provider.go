package provider

import (
	"slices"
	"strings"

	"github.com/erlorenz/sparkflow"
)

// Provider resolves assets and returns HTML.
type Provider struct {
	Resolver    sparkflow.Resolver
	Environment string
	Prefix      string
}

// New accepts a resolver and a pathPrefix (for use with http.StripPrefix).
func New(resolver sparkflow.Resolver, pathPrefix string) *Provider {
	return &Provider{Resolver: resolver, Prefix: pathPrefix}
}

// Returns HTML for the assets with all dependencies.
// On error returns an empty string.
func (p *Provider) HTML(paths ...string) string {
	allAssets := []sparkflow.Asset{}

	for _, path := range paths {
		assets, err := p.Resolver.Resolve(path)
		if err != nil {
			continue
		}
		allAssets = slices.Concat(allAssets, assets)
	}
	if len(allAssets) == 0 {
		return ""
	}

	tags := []string{}
	for _, asset := range allAssets {
		tags = append(tags, p.assetToTag(asset))
	}

	return strings.Join(tagSet(tags), "\n")
}
