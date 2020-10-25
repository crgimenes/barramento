package config

import (
	"os"
	"sync"

	"github.com/gosidekick/goconfig"
)

type Config struct {
	Node string `json:"node" cfg:"n" cfgDefault:".*" cfgRequired:"true"`
	Name string `json:"name" cfg:"name" cfgDefault:""`
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
		if err != nil {
			return
		}
		if cfg.Name == "" {
			cfg.Name, err = os.Hostname()
		}
	})
	return cfg, err
}
