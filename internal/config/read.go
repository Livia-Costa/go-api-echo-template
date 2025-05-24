package config

import (
	cmnconfig "github.com/victorsantosbrazil/go-api-echo-template/pkg/app/config"
)

func ReadConfig() (*Config, error) {
	var cfg Config

	readCfg := cmnconfig.ReadConfigOptions{
		ConfigName: "config",
		ConfigType: "yaml",
		ConfigPath: ".",
	}

	err := cmnconfig.ReadConfig(&cfg, readCfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
