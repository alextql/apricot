package config

import (
	"errors"
	"fmt"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Apricot struct {
		Listen            string
		DisableStartupLog bool
	}

	Log struct {
		File        string
		MaxAge      int
		MaxSize     int
		MaxBackup   int
		Compress    bool
		ForceColors bool
	}

	Database struct {
		Path string
	}
}

func LoadConfig() error {
	if _, err := toml.DecodeFile("config/apricot.toml", &C); err != nil {
		return errors.New(fmt.Sprintf("dump toml config fail, error : %s", err))
	}

	return nil
}
