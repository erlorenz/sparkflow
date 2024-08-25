package main

import (
	"cmp"
	"fmt"
	"log/slog"
	"os"

	"github.com/erlorenz/sparkflow/provider"
	"github.com/erlorenz/sparkflow/vite"
)

func main() {

	cfg := vite.Config{} // Use defaults
	vite, err := vite.New(cfg)
	if err != nil {
		slog.Error("failed to init", "error", err)
		os.Exit(1)
	}

	sf := provider.New(vite, "static")

	html := sf.HTML("resources/js/index.ts", "resources/js/other.ts", "resources/css/style.css")

	fmt.Println(html) // HTML tags

	fmt.Println(cmp.Or(vite.HMRScript(), "no hmr script, is static.")) // ""

}
