package vite

import (
	"cmp"
	"fmt"
)

// HMR returns a script for the hot-module-reload in vite dev server.
func (v *Vite) HMRScript() string {
	addr := cmp.Or(v.DevURL, defaultDevURL)
	if v.Environment == "development" {
		return fmt.Sprintf(`<script type="module" src="%s/@vite/client"></script>`, addr)
	}

	return ""
}
