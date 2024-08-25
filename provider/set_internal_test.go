package provider

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSet_RemovesTags(t *testing.T) {
	strs := []string{"john", "jane", "jane", "fred"}

	got := tagSet(strs)

	want := []string{"john", "jane", "fred"}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}

}
