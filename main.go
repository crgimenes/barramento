package main

import (
	"barramento/command"
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

	cmd := command.New()
	u := udp.New("udp4", cfg, cmd)

	if cfg.Message != "" {
		u.Send([]byte(cfg.Message))
		return
	}

	if cfg.ServerMode {
		u.Server()
	}
}
