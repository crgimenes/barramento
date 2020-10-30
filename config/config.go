package config

import (
	"os"
	"sync"

	"github.com/gosidekick/goconfig"
)

type Config struct {
	Node          string `json:"node" cfg:"n" cfgDefault:".*" cfgRequired:"true"`
	Name          string `json:"name" cfg:"name" cfgDefault:""`
	Message       string `json:"message" cfg:"msg"`
	IPv4          string `json:"ipv4" cfg:"ipv4" cfgDefault:"224.0.0.1:2222" cfgRequired:"true"`
	IPv6          string `json:"ipv6" cfg:"ipv6" cfgDefault:"[ff02::1%en0]:2222" cfgRequired:"true"`
	ServerAddress string `json:"server_address" cfg:"sa" cfgDefault:":2222" cfgRequired:"true"`
	ServerMode    bool   `json:"server_mode" cfg:"s"`
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
