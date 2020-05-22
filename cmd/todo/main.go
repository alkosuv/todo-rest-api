package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"

	v1 "github.com/gen95mis/todo-rest-api/internal/api/v1"
)

var (
	configPath string
)

func init() {
	flag.StringVar(
		&configPath,
		"config-path",
		"config/apiserver.toml",
		"path to config file",
	)
}

func main() {
	flag.Parse()

	apiserver := v1.NewAPIServer()
	if _, err := toml.DecodeFile(configPath, apiserver); err != nil {
		log.Fatal(err)
	}

	if err := apiserver.Start(); err != nil {
		log.Fatal(err)
	}
}
