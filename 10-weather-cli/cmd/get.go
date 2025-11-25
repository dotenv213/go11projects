package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get [city]",
	Short: "Get current weather for a city",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		city := args[0]
		getWeather(city)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}

type GeoResponse struct {
	Results []struct {
		Name      string  `json:"name"`
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
		Country   string  `json:"country"`
	} `json:"results"`
}

type WeatherResponse struct {
	CurrentWeather struct {
		Temperature float64 `json:"temperature"`
		WindSpeed   float64 `json:"windspeed"`
	} `json:"current_weather"`
}

func getWeather(city string) {
	fmt.Printf("🔍 Searching for %s...\n", city)

	lat, long, country, err := getCoordinates(city)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	temp, wind, err := fetchWeather(lat, long)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Println("-----------------------------")
	fmt.Printf("📍 Location: %s, %s\n", city, country)
	fmt.Printf("🌡️  Temperature: %.1f°C\n", temp)
	fmt.Printf("💨 Wind Speed:  %.1f km/h\n", wind)
	fmt.Println("-----------------------------")
}

func getCoordinates(city string) (float64, float64, string, error) {
	url := fmt.Sprintf("https://geocoding-api.open-meteo.com/v1/search?name=%s&count=1&language=en&format=json", city)
	
	resp, err := http.Get(url)
	if err != nil {
		return 0, 0, "", err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	
	var geo GeoResponse
	json.Unmarshal(body, &geo)

	if len(geo.Results) == 0 {
		return 0, 0, "", fmt.Errorf("city not found")
	}

	return geo.Results[0].Latitude, geo.Results[0].Longitude, geo.Results[0].Country, nil
}

func fetchWeather(lat, long float64) (float64, float64, error) {
	url := fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%f&longitude=%f&current_weather=true", lat, long)

	client := http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return 0, 0, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var weather WeatherResponse
	json.Unmarshal(body, &weather)

	return weather.CurrentWeather.Temperature, weather.CurrentWeather.WindSpeed, nil
}