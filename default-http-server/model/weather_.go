package model

import "encoding/json"

// example model record
//
//	{
//		"latitude": 35.7,
//		"longitude": 139.6875,
//		"generationtime_ms": 0.0232458114624023,
//		"utc_offset_seconds": 0,
//		"timezone": "GMT",
//		"timezone_abbreviation": "GMT",
//		"elevation": 40,
//		"current_units": {
//		  "time": "iso8601",
//		  "interval": "seconds",
//		  "weather_code": "wmo code"
//		},
//		"current": {
//		  "time": "2025-05-04T07:15",
//		  "interval": 900,
//		  "weather_code": 1
//		}
//	  }

type Weather struct {
	Latitude           float64      `json:"latitude"`
	Longitude          float64      `json:"longitude"`
	Timezone           string       `json:"timezone"`
	Generationtime_ms  float64      `json:"generationtime_ms"`
	Utc_offset_seconds float64      `json:"utc_offset_seconds"`
	Elevation          float64      `json:"elevation"`
	Current_units      CurrentUnits `json:"current_units"`
	Current            Current      `json:"current"`
}

type CurrentUnits struct {
	Time        string `json:"time"`
	Interval    string `json:"interval"`
	WeatherCode string `json:"weather_code"`
}

type Current struct {
	Time        string      `json:"time"`
	Interval    int         `json:"interval"`
	WeatherCode WeatherCode `json:"weather_code"`
}

func (currentUnits *CurrentUnits) UnmarshalJSON(data []byte) error {
	type Alias CurrentUnits
	aux := &Alias{}
	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}
	*currentUnits = CurrentUnits(*aux)
	return nil
}

func (current *Current) UnmarshalJSON(data []byte) error {
	type Alias Current
	aux := &Alias{}
	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}
	*current = Current(*aux)
	return nil
}

func (w Weather) GetWeather() string {
	return w.Current.WeatherCode.GetWeather()
}

func (w Weather) GetWeatherTime() string {
	return w.Current.Time
}

func (w Weather) GetWeatherInterval() int {
	return w.Current.Interval
}
