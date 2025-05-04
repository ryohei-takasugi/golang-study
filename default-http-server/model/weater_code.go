package model

import "fmt"

// Weather Code
// -- Weather variable documentation
// -- WMO Weather interpretation codes (WW)
// Code	| Description
// 0	| Clear sky
// 1, 2, 3	| Mainly clear, partly cloudy, and overcast
// 45, 48	| Fog and depositing rime fog
// 51, 53, 55	| Drizzle: Light, moderate, and dense intensity
// 56, 57	| Freezing Drizzle: Light and dense intensity
// 61, 63, 65	| Rain: Slight, moderate and heavy intensity
// 66, 67	| Freezing Rain: Light and heavy intensity
// 71, 73, 75	| Snow fall: Slight, moderate, and heavy intensity
// 77	| Snow grains
// 80, 81, 82	| Rain showers: Slight, moderate, and violent
// 85, 86	| Snow showers slight and heavy
// 95 *	| Thunderstorm: Slight or moderate
// 96, 99 *	| Thunderstorm with slight and heavy hail

type WeatherCode int

const (
	ClearSky WeatherCode = iota
	MainlyClear
	PartlyCloudy
	Overcast
	Fog
	DepositingRimeFog
	LightDrizzle
	ModerateDrizzle
	DenseDrizzle
	LightFreezingDrizzle
	DenseFreezingDrizzle
	LightRain
	ModerateRain
	HeavyRain
	LightFreezingRain
	HeavyFreezingRain
	LightSnow
	ModerateSnow
	HeavySnow
	LightSnowGrains
	LightRainShowers
	ModerateRainShowers
	ViolentRainShowers
	LightSnowShowers
	HeavySnowShowers
	ThunderstormSlight
	ThunderstormModerate
	ThunderstormSlightHail
	ThunderstormHeavyHail
)

var weatherCodeDescriptions = map[WeatherCode]string{
	ClearSky:               "Clear sky",
	MainlyClear:            "Mainly clear",
	PartlyCloudy:           "Partly cloudy",
	Overcast:               "Overcast",
	Fog:                    "Fog",
	DepositingRimeFog:      "Depositing rime fog",
	LightDrizzle:           "Light drizzle",
	ModerateDrizzle:        "Moderate drizzle",
	DenseDrizzle:           "Dense drizzle",
	LightFreezingDrizzle:   "Light freezing drizzle",
	DenseFreezingDrizzle:   "Dense freezing drizzle",
	LightRain:              "Light rain",
	ModerateRain:           "Moderate rain",
	HeavyRain:              "Heavy rain",
	LightFreezingRain:      "Light freezing rain",
	HeavyFreezingRain:      "Heavy freezing rain",
	LightSnow:              "Light snow",
	ModerateSnow:           "Moderate snow",
	HeavySnow:              "Heavy snow",
	LightSnowGrains:        "Light snow grains",
	LightRainShowers:       "Light rain showers",
	ModerateRainShowers:    "Moderate rain showers",
	ViolentRainShowers:     "Violent rain showers",
	LightSnowShowers:       "Light snow showers",
	HeavySnowShowers:       "Heavy snow showers",
	ThunderstormSlight:     "Thunderstorm slight",
	ThunderstormModerate:   "Thunderstorm moderate",
	ThunderstormSlightHail: "Thunderstorm slight hail",
	ThunderstormHeavyHail:  "Thunderstorm heavy hail",
}

func (wc WeatherCode) GetWeather() string {
	descriptions, ok := weatherCodeDescriptions[wc]
	if ok {
		return descriptions
	}

	return "Unknown weather code"
}

func (wc WeatherCode) String() string {
	return fmt.Sprintf("%d: %s", wc, wc.GetWeather())
}
