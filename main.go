package main

import (
	"fmt"
	"os"

	"github.com/alex-techs/apricot/apricot"
	"github.com/alex-techs/apricot/apricot/options"
	"github.com/urfave/cli/v2"
)

func main() {
	opts := options.NewDefaultOptions()
	app := &cli.App{
		Name:        "apricot",
		Description: "apricot is println.org server",
		Flags:       []cli.Flag{},
		Action: func(cmdline *cli.Context) error {
			var err error
			var engine *apricot.Apricot
			if engine, err = apricot.NewApricotEngine(opts); err != nil {
				return err
			}

			return engine.Listen()
		},
	}

	opts.AddFlags(&app.Flags)
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
