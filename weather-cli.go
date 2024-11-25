package main

import (
	"fmt"
	"net/http"
	"io"
	"encoding/json"
	"time"
	"os"

	"github.com/joho/godotenv"
)

type WeatherResponse struct { 
	Location struct { 
		City string `json:"name"`
		Country string `json:"country"`
	} `json:"location"`
	Current struct { 
		TempCalc float64 `json:"temp_c"`
		Condition struct { 
			Text string `json:"text"`
		} `json:"condition"`
	}
	Forecast struct { 
		ForecastDay [] struct { 
			Hour []struct {
				HourTime string `json:"time"`
				HourTempCalc float64 `json:"temp_c"`
				Condition struct { 
					Text string `json:"text"`
				} `json:"condition"`
			}`json:"hour"`
		} `json:"forecastday"`
	} `json:"forecast"`
}

func main() { 
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	apiKey := os.Getenv("API_KEY")
    if apiKey == "" {
        panic("API key not found! Set API_KEY environment variable.")
    }

	// default city
	q := "Athens"

	if len(os.Args) >= 2 {
		q = os.Args[1]
	}
	res, err := http.Get("http://api.weatherapi.com/v1/forecast.json?key=" + apiKey + "&q=" + q)

	if err != nil { 
		panic(err)
	}

	
	defer res.Body.Close()
	
	if res.StatusCode != 200 { 	
		panic("Weather API not available!")
	}
	
	body, err := io.ReadAll(res.Body)

	if err != nil { 
		panic(err)
	}

	var weather WeatherResponse
	err = json.Unmarshal(body, &weather)

	if err != nil { 
		panic(err)
	}
	
	now := time.Now()
	currentHour := now.Format("2006-01-02 15:00")

	green := "\033[32m"
	reset := "\033[0m"

	fmt.Printf("Forecast for %s,%s:\n",
		weather.Location.City,
		weather.Location.Country,
	)
	count := 0

	for _, day := range weather.Forecast.ForecastDay { 
		for _, hour := range day.Hour { 
			if hour.HourTime >= currentHour && count < 7 {
				if hour.HourTime == currentHour {
					// Print the current hour in green
					fmt.Printf(green+"%s: %.1f°C %s"+reset+"\n", hour.HourTime, hour.HourTempCalc, hour.Condition.Text)
				} else {
					// Print other hours normally
					fmt.Printf("%s: %.1f°C %s\n", hour.HourTime, hour.HourTempCalc, hour.Condition.Text)
				}
				count++
			}
		}
	}
}