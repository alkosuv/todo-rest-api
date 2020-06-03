package main

import (
	"fmt"
	"log"

	v1 "github.com/gen95mis/todo-rest-api/internal/api/v1"
	"github.com/gen95mis/todo-rest-api/internal/api/v1/config"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	c := config.NewConfig()
	apiserver := v1.NewAPIServer(
		c.BindAddr,
		c.LogLevel,
		c.SessionKey,
		c.DB,
	)

	fmt.Println("Server is listening...")
	if err := apiserver.Start(); err != nil {
		log.Fatal(err)
	}
}
