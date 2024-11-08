package entity

type Address struct {
	CEP        string `json:"cep"`
	Localidade string `json:"localidade"`
	UF         string `json:"uf"`
	Estado     string `json:"estado"`
}
