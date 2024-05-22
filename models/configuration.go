package models

type Configuration struct {
	Server      Server      `yaml:"server"`
	OpenWeather OpenWeather `yaml:"openWeather"`
}

type Server struct {
	Port int      `yaml:"port"`
	Cors []string `yaml:"cors"`
}

type OpenWeather struct {
	Host                   string `yaml:"host"`
	CurrentWeatherEndpoint string `yaml:"currentWeatherEndpoint"`
	ApiKey                 string `yaml:"apiKey"`
	Units                  string `yaml:"units"`
	Timeout                int    `yaml:"timeout"`
}
