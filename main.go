package main

import (
	"log"
	"os"
	"strconv"

	"weather-api/internal/cache"
	"weather-api/internal/handler"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	redisDB, _ := strconv.Atoi(os.Getenv("REDIS_DB"))
	cacheClient := cache.NewRedisCache(
		os.Getenv("REDIS_ADDR"),
		os.Getenv("REDIS_PASSWORD"),
		redisDB,
	)

	h := handler.Handler{Cache: cacheClient}

	r := gin.Default()
	r.GET("/weather", h.GetWeather)

	log.Println("🚀 Server running on :8080")
	r.Run(":8080")
}
