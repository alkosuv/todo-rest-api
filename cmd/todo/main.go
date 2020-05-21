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

	config := v1.NewConfigAPIServer()
	if _, err := toml.DecodeFile(configPath, config); err != nil {
		log.Fatal(err)
	}

	if err := v1.InitAPIServer(config); err != nil {
		log.Fatal(err)
	}
}
