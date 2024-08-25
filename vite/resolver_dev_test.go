package vite_test

import (
	"testing"

	"github.com/erlorenz/sparkflow/vite"
	"github.com/google/go-cmp/cmp"
)

func TestDevResolver_ExtensionMatches(t *testing.T) {
	t.Parallel()
	testvdr := &vite.DevResolver{}

	assets, err := testvdr.Resolve("js/index.ts")
	if err != nil {
		t.Fatal(err)
	}

	got := assets[0].Ext
	want := ".ts"
	if got != want {
		t.Fatalf("got, want\n%s", cmp.Diff(got, want))
	}

}

func TestDevResolver_AddsDefaultServerToFilepath(t *testing.T) {
	t.Parallel()
	testvdr := &vite.DevResolver{}

	assets, err := testvdr.Resolve("resources/js/index.ts")
	if err != nil {
		t.Fatalf("expected no error: %s", err)
	}

	got := assets[0].Filepath
	want := "http://localhost:5173/resources/js/index.ts"
	if got != want {
		t.Fatalf("got, want\n%s", cmp.Diff(got, want))
	}
}
func TestDevResolver_AddsServerToFilepath(t *testing.T) {
	t.Parallel()
	testvdr := &vite.DevResolver{URL: "http://localhost:5000"}

	assets, err := testvdr.Resolve("resources/js/index.ts")
	if err != nil {
		t.Fatalf("expected no error: %s", err)
	}

	got := assets[0].Filepath
	want := "http://localhost:5000/resources/js/index.ts"
	if got != want {
		t.Fatalf("got, want\n%s", cmp.Diff(got, want))
	}
}

func TestDevResolver_ErrorsOnBadExt(t *testing.T) {
	t.Parallel()
	testvdr := &vite.DevResolver{}

	assets, err := testvdr.Resolve("resources/js/index.jpg")
	if err == nil {
		t.Logf("assets: %#v", assets)
		t.Fatalf("expected error")
	}

}
