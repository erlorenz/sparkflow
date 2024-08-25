package vite

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParseEnvironment_ReturnsProdOnEmpty(t *testing.T) {
	t.Parallel()

	got := parseEnvironment("")

	want := "production"
	if want != got {
		t.Error(cmp.Diff(want, got))
	}
}

func TestParseEnvironment_ReturnsDevOnDev(t *testing.T) {
	t.Parallel()

	got := parseEnvironment("development")

	want := "development"
	if want != got {
		t.Error(cmp.Diff(want, got))
	}
}

func TestParseExt_ErrorsOnBad(t *testing.T) {
	t.Parallel()

	got, err := parseExt("index.mm")
	if err == nil {
		t.Error(err)
		t.Logf("got %s", got)
	}

}

func TestParseExt_ReturnsJSonJS(t *testing.T) {
	t.Parallel()

	got, err := parseExt("index.js")
	if err != nil {
		t.Error(err)
	}
	want := ".js"
	if want != got {
		t.Error(cmp.Diff(want, got))

	}
}

func TestParseExt_ReturnsTSonTS(t *testing.T) {
	t.Parallel()

	got, err := parseExt("index.ts")
	if err != nil {
		t.Error(err)
	}
	want := ".ts"
	if want != got {
		t.Error(cmp.Diff(want, got))
	}
}

func TestParseExt_ReturnsCSSOnCSS(t *testing.T) {
	t.Parallel()

	got, err := parseExt("index.css")
	if err != nil {
		t.Error(err)
	}
	want := ".css"
	if want != got {
		t.Error(cmp.Diff(want, got))
	}
}
