package main

import (
	"barramento/config"
	"barramento/udpserver"
	"fmt"
)

func main() {
	cfg, err := config.Get()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cfg)
	udpserver.Server()
}
