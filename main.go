package main

import (
	"log"
	"twittergo/domain/config"

	"twittergo/handlers"
)

func main() {
	if config.CheckConection() == 0 {
		log.Fatal("Error al conectar a la bbdd")
		return
	}
	handlers.Handlers()
}
