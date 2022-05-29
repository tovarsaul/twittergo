package main

import (
	"log"
	"twittergo/bd"

	"twittergo/handlers"
)

func main() {
	if bd.CheckConection() == 0 {
		log.Fatal("Error al conectar a la bbdd")
		return
	}
	handlers.Handlers()
	log.Println("Hola")
}
