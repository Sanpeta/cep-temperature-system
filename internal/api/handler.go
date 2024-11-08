package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Sanpeta/cep-temperature-system/internal/config"
	usecase "github.com/Sanpeta/cep-temperature-system/internal/usecase"
	"github.com/Sanpeta/cep-temperature-system/pkg/utils"
)

func HandlerWeather(w http.ResponseWriter, r *http.Request) {
	config, err := config.LoadConfig("../../")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	cep := r.URL.Query().Get("cep")
	if !utils.CheckCEP(cep) {
		http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
		return
	}

	address, err := usecase.FetchCity(cep)
	if address.CEP == "" || err != nil {
		http.Error(w, "can not find zipcode", http.StatusNotFound)
		return
	}

	weather, err := usecase.FetchTemperature(address, config.TOKEN_WEATHER_API)
	if weather.TempK == 0 || err != nil {
		http.Error(w, "error fetching temperature", http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(weather)
	if err != nil {
		http.Error(w, "error to marshal", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
