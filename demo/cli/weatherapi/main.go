package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const (
	APIKey = "760a1ba957fdba2e0a8796ee21e7f863"
)

type openWeatherMap struct{}

func (w openWeatherMap) tempreature(city string) (float64, error) {
	// REST call
	resp, err := http.Get(fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?APPID=%v&zip=%v", APIKey, city))
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var d struct {
		Sys struct {
			Sunrise int64 `json:"sunrise"`
			Sunset  int64 `json:"sunset"`
		}
		Weather []struct {
			ID          int64
			Main        string
			Description string
		}
		Main struct {
			Kelvin float64 `json:"temp"`
		} `json:"main"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		return 0, err
	}
	log.Printf("openWeatherMap: %s: %0.2f", city, d.Main.Kelvin)
	return d.Main.Kelvin, nil
}
func main() {
	var weathermap openWeatherMap
	temp, err := weathermap.tempreature("94568,US")
	if err != nil {
		log.Fatalf("Error thrown %v", err)
	}
	fmt.Printf("temp:%0.2f", temp)
}
