package client

import (
	"default-http-server/model"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type WeatherClient struct {
	// HTTP クライアント
	client       *http.Client
	latitude     string
	longitude    string
	current      string
	forecastDays string
}

func NewWeatherClient() *WeatherClient {
	return &WeatherClient{
		client:       &http.Client{},
		latitude:     "35.6895",
		longitude:    "139.6917",
		current:      "weather_code",
		forecastDays: "1",
	}
}

func (c WeatherClient) GetWeather() (string, error) {

	// API リクエストURI
	baseURL := "https://api.open-meteo.com/v1/forecast"
	params := map[string]string{
		"latitude":      c.latitude,
		"longitude":     c.longitude,
		"current":       c.current,
		"forecast_days": c.forecastDays,
	}

	uri, err := c.buildURI(baseURL, params)
	if err != nil {
		return "", err
	}

	fmt.Println("GET ", uri)

	resp, err := http.Get(uri)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	fmt.Println("Response Body:", string(body))

	var weather model.Weather

	err = json.Unmarshal(body, &weather)
	if err != nil {
		return "", err
	}

	return weather.GetWeather(), nil
}

func (c WeatherClient) buildURI(baseURL string, params map[string]string) (string, error) {
	// ベースURLを解析
	u, err := url.Parse(baseURL)
	if err != nil {
		return "", err
	}

	// クエリパラメータを設定
	query := u.Query()
	for key, value := range params {
		query.Set(key, value)
	}
	u.RawQuery = query.Encode()

	return u.String(), nil
}
