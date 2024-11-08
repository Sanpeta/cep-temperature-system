package main

import (
	"net/http"

	"github.com/Sanpeta/cep-temperature-system/internal/api"
)

func main() {
	http.HandleFunc("/weather", api.HandlerWeather)

	http.ListenAndServe(":8080", nil)
}
