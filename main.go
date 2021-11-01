package main

import (
	"log"

	"github.com/arfan21/golang-mygram/server"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	err := server.Start()
	if err != nil {
		log.Fatal(err)
	}

}
