package main

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Receivers []string
	Host      string
	Port      int
	Mail      string
	Password  string
	SaveAt    string
}

func DefaultConfig(path string) {
	c := Config{
		Host:   "localhost",
		Port:   25,
		SaveAt: "surveys",
	}
	writeConfig(c, path)
}

func writeConfig(c Config, path string) {
	fmt.Println(c)
	os.Remove(path)
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	enc := toml.NewEncoder(file)
	err = enc.Encode(c)
	if err != nil {
		panic(err)
	}
	fmt.Println("Configure has been generated at", path)
}

func LoadConfig(path string) (c Config, err error) {
	_, err = toml.DecodeFile(path, &c)
	return
}
