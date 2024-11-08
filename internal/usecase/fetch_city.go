package usecase

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Sanpeta/cep-temperature-system/internal/entity"
)

func FetchCity(cep string) (entity.Address, error) {
	cepAPI := "https://viacep.com.br/ws/%s/json/"
	cepAPI = fmt.Sprintf(cepAPI, cep)

	resp, err := http.Get(cepAPI)
	if err != nil {
		fmt.Println("Error:", err)
		return entity.Address{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return entity.Address{}, err
	}

	var address entity.Address
	err = json.Unmarshal(body, &address)
	if err != nil {
		fmt.Println("Error:", err)
		return entity.Address{}, err
	}

	return address, nil
}
