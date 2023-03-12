package main

import (
	"log"

	"restplaceholder/config"
	"restplaceholder/internal/app/server"
	"restplaceholder/internal/app/server/handler"
	"restplaceholder/internal/app/server/service"
)

func main() {
	cfg := config.NewConfig()
	if err := cfg.ReadConfig(); err != nil {
		log.Fatal(err)
	}

	serv := service.NewService()
	hand := handler.NewHandler(serv)

	s := server.NewServer()
	if err := s.Run(cfg.GetServerPort(), hand.InitRoutes()); err != nil {
		log.Fatal(err)
	}
}
