package main

import (
	"barbot/internal/router"
	"log"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	r := router.New()

	if err := r.Run(); err != nil {
		log.Fatalln(err)
	}
}
