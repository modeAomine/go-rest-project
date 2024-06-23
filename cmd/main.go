package main

import (
	"log"
	web_go "web-go"
	"web-go/pkg/handler"
)

func main() {
	handler := new(handler.Handler)
	srv := new(web_go.Server)
	if err := srv.Run("8494", handler.InitRoutes()); err != nil {
		log.Fatal(err)
	}
}
