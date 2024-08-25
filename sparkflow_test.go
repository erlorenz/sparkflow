package sparkflow_test

import (
	"strings"
	"testing"

	"github.com/erlorenz/sparkflow/provider"
	"github.com/erlorenz/sparkflow/vite"
)

func TestGetTags_DevVite(t *testing.T) {

	cfg := vite.Config{Environment: "development"}
	vresolver, err := vite.New(cfg)
	if err != nil {
		t.Fatal(err)
	}

	hmr := vresolver.HMRScript()
	if hmr == "" {
		t.Error("hmr is supposed to be a script in dev mode")

	}

	sf := provider.New(vresolver, "static")
	html := sf.HTML("resources/asset.ts", "resources/asset.css")

	lines := strings.Split(html, "\n")
	if len(lines) != 2 {
		t.Errorf("want 2 lines, got %d", len(lines))
		t.Log(html)
	}
}

func TestGetTags_StaticVite(t *testing.T) {

	cfg := vite.Config{PublicDir: "vite/testdata"}
	vresolver, err := vite.New(cfg)
	if err != nil {
		t.Fatal(err)
	}

	hmr := vresolver.HMRScript()
	if hmr != "" {
		t.Errorf("no hmr in static, got %s", hmr)
	}

	sf := provider.New(vresolver, "static")
	html := sf.HTML("resources/js/thing/index.ts", "resources/css/index.css")

	lines := strings.Split(html, "\n")
	if len(lines) != 3 {
		t.Errorf("want 3 lines, got %d", len(lines))
		t.Log(html)
	}
}
