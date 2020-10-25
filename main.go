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
	fmt.Println(cfg)

	fmt.Println("hostname:", cfg.Name)
	udp.Server()
}
