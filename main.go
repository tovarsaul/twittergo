package main

import (
	"log"
	"twittergo/domain"

	"twittergo/handlers"
)

func main() {
	if domain.CheckConection() == 0 {
		log.Fatal("Error al conectar a la bbdd")
		return
	}
	handlers.Handlers()
	log.Println("Hola")
}
