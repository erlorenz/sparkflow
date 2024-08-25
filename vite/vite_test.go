package vite_test

import (
	"testing"

	"github.com/erlorenz/sparkflow/vite"
	"github.com/google/go-cmp/cmp"
)

func TestConfig_ErrorsOnInvalidEnvironment(t *testing.T) {
	t.Parallel()
	cfg := vite.Config{Environment: "something"}

	_, err := vite.New(cfg)
	if err == nil {
		t.Fatalf("expected error with bad environment")
	}
}

func TestConfig_ReturnsHMRScriptWithDevEnv(t *testing.T) {
	t.Parallel()
	cfg := vite.Config{Environment: "development"}

	res, err := vite.New(cfg)
	if err != nil {
		t.Fatal(err)
	}

	want := "development"
	got := res.Environment
	if want != got {
		t.Error(cmp.Diff(want, got))
	}

	want = `<script type="module" src="http://localhost:5173/@vite/client"></script>`
	got = res.HMRScript()

	if got != want {
		t.Error(cmp.Diff(want, got))
	}
}

func TestConfig_ReturnsEmptyHMRScriptWithProd(t *testing.T) {
	t.Parallel()
	cfg := vite.Config{Environment: "development"}

	res, err := vite.New(cfg)
	if err != nil {
		t.Fatal(err)
	}

	res.Environment = "production" // manually set to prod

	want := ""
	got := res.HMRScript()

	if got != want {
		t.Error(cmp.Diff(want, got))
	}
}

func TestStaticInit_SuccesfullyInits(t *testing.T) {
	t.Parallel()
	cfg := vite.Config{
		PublicDir: "testdata",
	}

	res, err := vite.New(cfg)
	if err != nil {
		t.Fatal(err)
	}

	got := res.HMRScript()
	want := ""
	if got != want {
		t.Error(cmp.Diff(want, got))
	}
}
