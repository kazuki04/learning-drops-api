package config

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
)

type ConfigList struct {
	Port    int
	LogFile string
}

var Config ConfigList

func init() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	Config = ConfigList{
		Port:    cfg.Section("web").Key("port").MustInt(),
		LogFile: cfg.Section("log").Key("log_file").String(),
	}
}
