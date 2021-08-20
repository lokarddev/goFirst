package main

import (
	"goFirst"
	"goFirst/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(goFirst.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("Error occured while: %s", err.Error())
	}
}
