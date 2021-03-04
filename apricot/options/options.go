package options

import (
	"github.com/urfave/cli/v2"
)

type Options struct {
	ConfigFile   string
	DatabaseFile string
}

func NewDefaultOptions() *Options {
	return &Options{}
}

func (o *Options) AddFlags(flags *[]cli.Flag) {
	*flags = append(*flags, &cli.StringFlag{
		Name:        "config",
		Aliases:     []string{"c"},
		Usage:       "config file path",
		Destination: &o.ConfigFile,
		Required:    true,
	})

	*flags = append(*flags, &cli.StringFlag{
		Name:        "database",
		Aliases:     []string{"d"},
		Usage:       "sqlite database file path",
		Destination: &o.DatabaseFile,
		Required:    true,
	})
}
