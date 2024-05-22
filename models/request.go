package models

type Request struct {
	Lat  string `json:"url" validate:"required,latitude"`
	Long string `json:"long" validate:"required,longitude"`
}
