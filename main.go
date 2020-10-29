package main

import (
	"barramento/config"
	"barramento/udp"
	"fmt"
)

func main() {
	cfg, err := config.Get()
	if err != nil {
		fmt.Println(err)
		return
	}

	u := udp.New("udp4", cfg)

	if cfg.Message != "" {
		u.Send([]byte(cfg.Message))
		return
	}

	if cfg.ServerMode {
		u.Server()
	}
}
