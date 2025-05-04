package handler

import (
	"net/http"

	"default-http-server/client"
)

type WeatherHandler struct {
	client   *client.WeatherClient
	response http.ResponseWriter
	request  *http.Request
}

func CreateWeatherHandler(w http.ResponseWriter, r *http.Request) *WeatherHandler {
	return &WeatherHandler{
		client:   client.NewWeatherClient(),
		response: w,
		request:  r,
	}
}

func (h WeatherHandler) Handle() {
	weatherCode, err := h.client.GetWeather()
	if err != nil {
		http.Error(h.response, "Failed to get weather", http.StatusInternalServerError)
		return
	}
	h.response.WriteHeader(http.StatusOK)
	h.response.Write([]byte(weatherCode))
}
