// +build ignore

package main

import (
	"fmt"
	"os"

	"github.com/99designs/gqlgen/api"
	"github.com/99designs/gqlgen/codegen/config"
	gbgen "github.com/web-ridge/gqlgen-sqlboiler/v2"
)

func main() {
	cfg, err := config.LoadConfigFromDefaultLocations()
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to load config", err.Error())
		os.Exit(2)
	}

	convertHelpersDir := "helpers"
	sqlboilerDir := "models"
	gqlgenModelDir := "graph"
	err = api.Generate(cfg,
		api.AddPlugin(gbgen.NewConvertPlugin(
			convertHelpersDir, // directory where convert.go, convert_input.go and preload.go should live
			sqlboilerDir,      // directory where sqlboiler files are put
			gqlgenModelDir,    // directory where gqlgen models live
		)),
		api.AddPlugin(gbgen.NewResolverPlugin(
			convertHelpersDir,
			sqlboilerDir,
			gqlgenModelDir,
			"", // leave empty if you don't have auth
		)),
	)
	if err != nil {
		fmt.Println("error!!")
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(3)
	}
}
