package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type WeatherResponse struct {
	ResolvedAddress string `json:"resolvedAddress"`
	Description     string `json:"description"`
	CurrentConditions struct {
		Temp     float64 `json:"temp"`
		Conditions string `json:"conditions"`
	} `json:"currentConditions"`
}

func FetchWeather(city string) (*WeatherResponse, error) {
	apiKey := os.Getenv("VISUAL_CROSSING_API_KEY")
	url := fmt.Sprintf("https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/%s?unitGroup=metric&key=%s&contentType=json", city, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch weather data: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("error from weather API: %s", string(body))
	}

	var weatherResp WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&weatherResp); err != nil {
		return nil, err
	}

	return &weatherResp, nil
}
