package provider

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/erlorenz/sparkflow"
)

func (p *Provider) assetToTag(asset sparkflow.Asset) string {
	path := asset.Filepath

	// Add prefix to static
	isStatic := !strings.HasPrefix(asset.Filepath, "http")
	if isStatic {
		path = filepath.Join(p.Prefix, path)
	}

	switch {
	case asset.IsChunk:
		return fmt.Sprintf(`<link rel="modulepreload" href="%s"></link>`, path)

	case asset.Ext == ".js" || asset.Ext == ".ts":
		return fmt.Sprintf(`<script type="module" src="%s"></script>`, path)

	case asset.Ext == ".css":
		return fmt.Sprintf(`<link rel="stylesheet" href="%s"></link>`, path)

	default:
		return ""
	}
}
