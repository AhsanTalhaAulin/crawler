package main

import (
	"crawler/handler"
	"log"

	"crawler/conn"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const EchoLogFormat = "time: ${time_rfc3339_nano} || ${method}: ${uri} || status: ${status} || latency: ${latency_human} \n"

func main() {
	// Load configuration
	// Initialize logger

	// Connect to MongoDB and Redis
	err := conn.MongoConnect()
	if err != nil {
		log.Fatal("Mongo connection error : ", err.Error())
	}

	err = conn.RedisConnect()
	if err != nil {
		log.Fatal("Redis connection error : ", err.Error())
	}

	// Create Echo server instance
	e := echo.New()

	// Middleware for logging
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{Format: EchoLogFormat}))

	// Healthcheck endpoint
	e.GET("/healthcheck", handler.Healthcheck)

	// Endpoint for getting child URLs
	e.GET("/urls/:parentUrl", handler.GetChildUrls)

	// Start the server
	e.Logger.Fatal(e.Start(":8080"))
}
