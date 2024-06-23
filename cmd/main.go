package main

import (
	"log"
	web_go "web-go"
	"web-go/pkg/handler"
	"web-go/pkg/repo"
	"web-go/pkg/service"
)

func main() {
	repos := repo.NewRepo()
	services := service.NewService(repos)
	handler := handler.NewHandler(services)

	srv := new(web_go.Server)
	if err := srv.Run("8494", handler.InitRoutes()); err != nil {
		log.Fatal(err)
	}
}
