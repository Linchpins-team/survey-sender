package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var path string
	flag.StringVar(&path, "c", "config.toml", "use -c to choose configure file.")
	flag.Parse()

	config, err := LoadConfig(path)
	if err != nil {
		fmt.Println("Configure file not found. Generating...")
		InitConfig(path)
		return
	}
	fmt.Println(config)
}