package vite

import (
	"testing"

	"github.com/erlorenz/sparkflow"
	"github.com/erlorenz/sparkflow/internal/mock"
	"github.com/erlorenz/sparkflow/internal/testhelpers"
	"github.com/google/go-cmp/cmp"
)

// MANIFEST AT testdata/
func TestParseManifest_ErrorsWithMissingFile(t *testing.T) {
	t.Parallel()

	if _, err := manifestToAssetMap(nil); err == nil {
		t.Errorf("expected error with missing manifest")
	}
}

func TestParseManifest(t *testing.T) {
	t.Parallel()

	fm := mock.FakeManifest{}
	fm.AddChunk("assets/_shared-HASH.js", "")
	fm.AddChunk("assets/index-HASH0.js", "resources/js/other/index.ts", "_shared-HASH.js")
	fm.AddChunk("assets/index-HASH1.js", "resources/js/thing/index.ts", "_shared-HASH.js")
	fm.AddChunk("assets/style-HASH.css", "resources/css/style.css")

	assetMap, err := manifestToAssetMap(fm.MustJSON())
	if err != nil {
		t.Error(err)
	}

	t.Run("Parse shared", func(t *testing.T) {
		t.Parallel()

		got, ok := assetMap["_shared-HASH.js"]
		if !ok {
			t.Errorf("missing key _shared-HASH.js")
			t.Logf("%v", testhelpers.JSON(t, assetMap))
			t.FailNow()
		}
		want := sparkflow.Asset{
			Filepath:    "assets/_shared-HASH.js",
			Ext:         ".js",
			IsChunk:     true,
			LogicalPath: "",
			Imports:     nil,
		}

		if !cmp.Equal(got, want) {
			t.Error(cmp.Diff(want, got))
			t.Logf("%#+v", assetMap["_shared-HASH.js"])
		}
	})
	t.Run("Parse entry", func(t *testing.T) {
		t.Parallel()

		got := assetMap["resources/js/thing/index.ts"]
		want := sparkflow.Asset{
			Filepath:    "assets/index-HASH1.js",
			Ext:         ".js",
			IsChunk:     false,
			LogicalPath: "resources/js/thing/index.ts",
			Imports:     []string{"_shared-HASH.js"},
		}

		if !cmp.Equal(got, want) {
			t.Error(cmp.Diff(want, got))
			t.Logf("%#+v", assetMap["resources/js/thing/index.ts"])
		}
	})
	t.Run("Parse css", func(t *testing.T) {
		t.Parallel()

		got := assetMap["resources/css/style.css"]
		want := sparkflow.Asset{
			Filepath:    "assets/style-HASH.css",
			Ext:         ".css",
			LogicalPath: "resources/css/style.css",
		}

		if !cmp.Equal(got, want) {
			t.Error(cmp.Diff(want, got))
			t.Logf("%#+v", assetMap["resources/css/style.css"])
		}
	})
}
