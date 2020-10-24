package config

import (
	"sync"

	"github.com/gosidekick/goconfig"
)

type Config struct {
	Node string `json:"node" cfg:"n" cfgDefault:".*" cfgRequired:"true"`
}

var (
	once sync.Once
	cfg  *Config
)

func Get() (*Config, error) {
	var err error
	once.Do(func() {
		goconfig.PrefixEnv = "br"
		cfg = &Config{}
		err = goconfig.Parse(cfg)
	})
	return cfg, err
}
