package config

import (
	"errors"
	"fmt"

	"github.com/BurntSushi/toml"
)

type database struct {
	Path string
}

type log struct {
	File        string
	MaxAge      int
	MaxSize     int
	MaxBackup   int
	Compress    bool
	ForceColors bool
}

type apricot struct {
	Listen            string
	DisableStartupLog bool
}

func newDefaultApricot() apricot {
	return apricot{
		Listen:            ":8999",
		DisableStartupLog: true,
	}
}

func newDefaultLog() log {
	return log{
		File:        "logs/apricot.log",
		MaxAge:      30,
		MaxSize:     3,
		MaxBackup:   100,
		Compress:    true,
		ForceColors: false,
	}
}
func newDefaultDatabase() database {
	return database{
		Path: "apricot.sqlite",
	}
}

type Config struct {
	Apricot  apricot
	Log      log
	Database database
}

var c Config

func Get() Config {
	return c
}

func LoadConfig(file string) error {
	tmp := new(Config)
	tmp.Log = newDefaultLog()
	tmp.Apricot = newDefaultApricot()
	tmp.Database = newDefaultDatabase()
	if _, err := toml.DecodeFile(file, tmp); err != nil {
		return errors.New(fmt.Sprintf("dump toml config fail, error : %s", err))
	}
	c = *tmp

	return nil
}
