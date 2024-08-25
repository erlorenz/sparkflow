package testhelpers

import (
	"encoding/json"
	"testing"
)

func JSON(t *testing.T, v any) string {
	t.Helper()

	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(b)
}
