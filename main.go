package main

import (
	"barramento/config"
	"fmt"
)

func main() {
	cfg, err := config.Get()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cfg)
}
