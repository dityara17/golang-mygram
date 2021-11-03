package main

import (
	"log"

	"github.com/arfan21/golang-mygram/server"
	_ "github.com/joho/godotenv/autoload"
)

// @title Hacktiv8 TODOS API
// @version 1.0
// @description This is API for completing final project 1 hacktiv8
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /

func main() {
	err := server.Start()
	if err != nil {
		log.Fatal(err)
	}
}
