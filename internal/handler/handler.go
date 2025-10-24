package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"weather-api/internal/cache"
	"weather-api/internal/weather"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Cache *cache.RedisCache
}

func (h *Handler) GetWeather(c *gin.Context) {
	city := c.Query("city")
	if city == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "city parameter is required"})
		return
	}

	// Check Redis cache
	data, err := h.Cache.Get(city)
	if err == nil {
		var cached weather.WeatherResponse
		json.Unmarshal([]byte(data), &cached)
		c.JSON(http.StatusOK, gin.H{"source": "cache", "data": cached})
		return
	}

	// Fetch from API
	weatherData, err := weather.FetchWeather(city)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Cache the result for 12 hours
	jsonData, _ := json.Marshal(weatherData)
	h.Cache.Set(city, string(jsonData), 12*time.Hour)

	c.JSON(http.StatusOK, gin.H{"source": "api", "data": weatherData})
}
