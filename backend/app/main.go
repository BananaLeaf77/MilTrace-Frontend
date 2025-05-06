package main

import (
	"MilTrace/config"
	"MilTrace/delivery"
	"MilTrace/repository"
	"MilTrace/services"
	"log"

	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env")
	}

	StartHTTP()
}

func StartHTTP() {
	db, err := config.BootDB()
	if err != nil {
		log.Fatalf("Failed to Boot DB, err: %s", err)
	}

	tracerRepo := repository.NewDeviceRepository(db)
	tracerService := services.NewDeviceService(tracerRepo)

	// NetHTTP Router
	netHttp := http.NewServeMux()
	delivery.NewDeviceHandler(netHttp, tracerService)
	log.Println("Starting HTTP server on :8080")

	err = http.ListenAndServe(":8080", netHttp)
	if err != nil {
		log.Fatalf("Failed to start HTTP server, err: %s", err)
	}
	log.Println("HTTP server started on :8080")

}
