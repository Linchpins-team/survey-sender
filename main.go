package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var (
	Configure Config
)

func initConfig() {
	var path string
	flag.StringVar(&path, "c", "config.toml", "use -c to choose configure file.")
	flag.Parse()

	var err error
	Configure, err = LoadConfig(path)
	if err != nil {
		fmt.Println("Configure file not found. Generating...")
		DefaultConfig(path)
		return
	}
}

func main() {
	initConfig()
	handler := NewHandler()
	http.HandleFunc("/", handler.newSurvey)
	fmt.Println("Listen on http://localhost:1510")
	log.Fatal(http.ListenAndServe("localhost:1510", nil))
}
