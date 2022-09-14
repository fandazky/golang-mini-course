package main

import (
	"go-restful-demo/config"
	"go-restful-demo/routers"
)

func main() {
	config.InitDB()
	e := routers.New()
	e.Logger.Fatal(e.Start(":8000"))
}
