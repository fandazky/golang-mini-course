package main

import (
	"go-restful-demo/config"
	"go-restful-demo/routers"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config.InitDB()
	e := routers.New()
	e.Logger.Fatal(e.Start(":8000"))
}
