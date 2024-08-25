package main

import (
	"fmt"

	sf "github.com/erlorenz/sparkflow/provider"
	"github.com/erlorenz/sparkflow/vite"
)

func main() {

	cfg := vite.Config{Environment: "development"}
	vite, _ := vite.New(cfg) // Dev does not error
	provider := sf.New(vite, "static")

	html := provider.HTML("resources/js/index.ts", "resources/css/style.css")

	fmt.Println(html) // HTML tags

	fmt.Println(vite.HMRScript()) // Client script
}
