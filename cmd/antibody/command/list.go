package command

import (
	"fmt"

	"github.com/caarlos0/antibody"
	"github.com/caarlos0/antibody/bundle"
	"github.com/codegangsta/cli"
)

// List all downloaded bundles
var List = cli.Command{
	Name:  "list",
	Usage: "list all currently downloaded bundles",
	Action: func(ctx *cli.Context) {
		for _, b := range bundle.List(antibody.Home()) {
			fmt.Println(b.Name())
		}
	},
}