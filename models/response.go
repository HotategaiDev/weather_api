package models

type Response struct {
	StatusCode             int                  `json:"statusCode"`
	Condition              string               `json:"condition"`
	Temperature            string               `json:"temperature"`
	TemperatureUnit        string               `json:"temperatureUnit"`
	TemperatureDescription string               `json:"temperatureDescription"`
	Detail                 *OpenWeatherResponse `json:"detail"`
}

type WrappedResponse struct {
	Response *Response `json:"response"`
	Error    error     `json:"error"`
}

type ResponseError struct {
	DetailError string `json:"error"`
}
