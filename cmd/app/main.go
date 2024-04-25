package main

import (
	"log"
	"net/http"
	"time"

	"github.com/mystpen/cryptocurrency-rates/internal/delivery"
	"github.com/mystpen/cryptocurrency-rates/internal/repository/api"
	"github.com/mystpen/cryptocurrency-rates/internal/service"
)

func main() {
	apiClient := api.NewApiClient()

	service := service.NewService(apiClient)

	handler := delivery.NewHandler(service)

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      handler.Routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	log.Println("the server is running on http://localhost" + srv.Addr)

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
