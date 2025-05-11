package main

import (
	"MilTrace/config"
	"MilTrace/delivery"
	"MilTrace/repository"
	"MilTrace/services"
	"log"

	"github.com/gin-gonic/gin"
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
		log.Fatalf("Failed to Boot DB ⛔, err: %s", err)
	}

	tracerRepo := repository.NewDeviceRepository(db)
	tracerService := services.NewDeviceService(tracerRepo)

	// NetHTTP Router
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	engine.SetTrustedProxies(nil)
	delivery.NewDeviceHandler(engine, tracerService)
	log.Println("Starting HTTP server on :8080 🌐")
	log.Println("Endpoint Server Check: http://localhost:8080/ping")
	if err := engine.Run(":8080"); err != nil {
		log.Fatalf("Failed to start HTTP server ⛔, err: %s", err)
	}

}
