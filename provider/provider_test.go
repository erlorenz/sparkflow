package provider_test

import (
	"strings"
	"testing"

	"github.com/erlorenz/sparkflow"
	"github.com/erlorenz/sparkflow/internal/mock"
	"github.com/erlorenz/sparkflow/provider"
	"github.com/google/go-cmp/cmp"
)

func TestAssetTags_ReturnsEmptyOnError(t *testing.T) {
	res := mock.Resolver{Error: sparkflow.ErrNotFound}
	prov := provider.New(res, "")

	got := prov.HTML("anything", "anything")
	want := ""
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestAssetTags_ReturnsSingleTagWithNoImports(t *testing.T) {
	res := mock.Resolver{
		Assets: []sparkflow.Asset{
			{
				Filepath:    "assets/index-HASH.js",
				LogicalPath: "resources/js/index.ts",
				Ext:         ".js",
			}},
	}
	prov := provider.New(res, "")

	got := prov.HTML("resources/js/index.ts")
	want := `<script type="module" src="assets/index-HASH.js"></script>`
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestAssetTags_AddsPrefixStatic(t *testing.T) {
	res := mock.Resolver{
		Assets: []sparkflow.Asset{
			{
				Filepath:    "assets/index-HASH.js",
				LogicalPath: "resources/js/index.ts",
				Ext:         ".js",
			}},
	}
	prov := provider.New(res, "static")

	got := prov.HTML("resources/js/index.ts")
	want := `<script type="module" src="static/assets/index-HASH.js"></script>`
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestAssetTags_ReturnsSingleTagWithImport(t *testing.T) {
	res := mock.Resolver{
		Assets: []sparkflow.Asset{
			{
				Filepath:    "assets/index-HASH.js",
				LogicalPath: "resources/js/index.ts",
				Ext:         ".js",
				Imports:     []string{"_shared-HASH.js"},
			},
			{
				Filepath: "assets/_shared-HASH.js",
				Ext:      ".js",
				IsChunk:  true,
			},
		},
	}
	prov := provider.New(res, "")

	got := prov.HTML("resources/js/index.ts")
	want := `<script type="module" src="assets/index-HASH.js"></script>
<link rel="modulepreload" href="assets/_shared-HASH.js"></link>`
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestAssetTags_ReturnsMultipleTags(t *testing.T) {
	res := mock.Resolver{
		Assets: []sparkflow.Asset{
			{
				Filepath:    "assets/index-HASH.js",
				LogicalPath: "resources/js/index.ts",
				Ext:         ".js",
				Imports:     []string{"_shared-HASH.js"},
			},
			{
				Filepath: "assets/_shared-HASH.js",
				Ext:      ".js",
				IsChunk:  true,
			},
			{
				Filepath:    "assets/style-HASH.css",
				LogicalPath: "resources/css/style.css",
				Ext:         ".css",
			},
		},
	}
	prov := provider.New(res, "")

	got := prov.HTML("resources/js/index.ts")
	tags := []string{`<script type="module" src="assets/index-HASH.js"></script>`,
		`<link rel="modulepreload" href="assets/_shared-HASH.js"></link>`,
		`<link rel="stylesheet" href="assets/style-HASH.css"></link>`}
	want := strings.Join(tags, "\n")

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
