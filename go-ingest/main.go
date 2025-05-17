package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/joho/godotenv"
	"github.com/segmentio/kafka-go"
)

type WeatherData struct {
	Location string  `json:"location"`
	TempC    float64 `json:"temp_c"`
	Time     string  `json:"time"`
}

func fetchWeather(apiKey, city string) (*WeatherData, error) {
	client := resty.New()
	resp, err := client.R().
		SetQueryParams(map[string]string{
			"key": apiKey,
			"q":   city,
		}).
		SetResult(map[string]interface{}{}).
		Get("http://api.weatherapi.com/v1/current.json")
	if err != nil {
		return nil, err
	}

	result := resp.Result().(map[string]interface{})
	loc := result["location"].(map[string]interface{})
	current := result["current"].(map[string]interface{})

	return &WeatherData{
		Location: loc["name"].(string),
		TempC:    current["temp_c"].(float64),
		Time:     loc["localtime"].(string),
	}, nil
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	apiKey := os.Getenv("API_KEY")
	city := "Lagos"
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"kafka:9092"},
		Topic:    "weather",
		Balancer: &kafka.LeastBytes{},
	})

	for {
		data, err := fetchWeather(apiKey, city)
		if err != nil {
			fmt.Println("Error fetching weather:", err)
			continue
		}

		jsonData, _ := json.Marshal(data)
		err = writer.WriteMessages(context.Background(),
			kafka.Message{
				Key:   []byte(data.Location),
				Value: jsonData,
			},
		)

		if err != nil {
			fmt.Println("error writing data:", err)
		}

		fmt.Println("Pushed to Kafka:", string(jsonData))
		time.Sleep(60 * time.Second)
	}
}
