package config

var C = &Config{
	Apricot: struct {
		Listen            string
		DisableStartupLog bool
	}{
		Listen:            ":8132",
		DisableStartupLog: true,
	},
	Log: struct {
		File        string
		MaxAge      int
		MaxSize     int
		MaxBackup   int
		Compress    bool
		ForceColors bool
	}{
		File:        "logs/apricot.log",
		MaxAge:      30,
		MaxSize:     3,
		MaxBackup:   100,
		Compress:    true,
		ForceColors: false,
	},
	Database: struct {
		Path string
	}{
		Path: "apricot.sqlite",
	},
}
