package main

import (
	"log"
	"net/http"

	"github.com/Sanpeta/cep-temperature-system/internal/api"
	"github.com/Sanpeta/cep-temperature-system/internal/config"
)

func main() {
	config, err := config.LoadConfig("../")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	http.HandleFunc("/weather", api.HandlerWeather)

	http.ListenAndServe(config.SERVER_ADDRESS, nil)
}
