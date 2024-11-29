package main

import (
	"fmt"
	"go-url-shortener-service/configs"
	"go-url-shortener-service/internal/auth"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	router := http.NewServeMux()

	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: conf,
	})

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Sever on 8081")
	server.ListenAndServe()

}
