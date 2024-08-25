package vite

import (
	"errors"
	"testing"

	"github.com/erlorenz/sparkflow"
	"github.com/erlorenz/sparkflow/internal/mock"
	"github.com/google/go-cmp/cmp"
)

func TestStaticResolver_TestResolve(t *testing.T) {
	t.Parallel()

	fm := mock.FakeManifest{}
	fm.AddChunk("assets/_shared-HASH.js", "")
	fm.AddChunk("assets/index-HASH0.js", "resources/js/other/index.ts", "_shared-HASH.js")
	fm.AddChunk("assets/index-HASH1.js", "resources/js/thing/index.ts", "_shared-HASH.js")
	fm.AddChunk("assets/index-HASH.css", "resources/css/style.css")

	static := &StaticResolver{
		PublicDir: "dist",
		BuildDir:  defaultBuildDir,
		Manifest:  defaultManifestPath,
	}

	// Dont use method
	assetMap, err := manifestToAssetMap(fm.MustJSON())
	if err != nil {
		t.Error(err)
	}
	static.AssetMap = assetMap

	t.Run("Error on missing", func(t *testing.T) {
		t.Parallel()

		_, err := static.Resolve("something.js")
		if err == nil {
			t.Fatal("expected error on missing 'something.js'")
		}

		want := sparkflow.ErrNotFound
		if !errors.Is(err, want) {
			t.Error(cmp.Diff(want, err))
		}

	})
	t.Run("Resolve single asset", func(t *testing.T) {
		t.Parallel()

		assets, err := static.Resolve("resources/css/style.css")
		if err != nil {
			t.Fatal(err)
		}
		want := 1
		got := len(assets)
		if want != got {
			t.Logf("wanted length of 1")
			t.Error(cmp.Diff(want, got))
		}

	})

	t.Run("Resolve asset and import", func(t *testing.T) {
		t.Parallel()

		assets, err := static.Resolve("resources/js/other/index.ts")
		if err != nil {
			t.Fatal(err)
		}
		want := 2
		got := len(assets)
		if want != got {
			t.Logf("wanted length of 2")
			t.Error(cmp.Diff(want, got))
		}

	})

}
