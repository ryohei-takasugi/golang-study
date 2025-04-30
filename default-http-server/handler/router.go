package handler

import (
	"io"
	"net/http"
	"net/url"
)

type EchoHandler struct{}

func (h EchoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Handle the request here
	// For example, you can write a simple response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, Echo! Path:" + r.URL.Path))
}

type WeatherEchoHandler struct{}

func (h WeatherEchoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// API リクエストURI
	baseURL := "https://api.open-meteo.com/v1/forecast"
	params := map[string]string{
		"latitude":      "35.6895",
		"longitude":     "139.6917",
		"current":       "weather_code",
		"forecast_days": "1",
	}

	uri, err := buildURI(baseURL, params)
	if err != nil {
		http.Error(w, "Failed to build URI", http.StatusInternalServerError)
		return
	}

	resp, err := http.Get(uri)
	if err != nil {
		http.Error(w, "Failed to fetch weather data", http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read response body", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(body))
}

func buildURI(baseURL string, params map[string]string) (string, error) {
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
