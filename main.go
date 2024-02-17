package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"

	"github.com/gorilla/mux"
)

type Weather struct {
	TempC float64 `json:"tempC"`
	TempF float64 `json:"tempF"`
	TempK float64 `json:"tempK"`
}

type ViaCEPResponse struct {
	Localidade string `json:"localidade"`
	Erro       bool   `json:"erro"`
}

type WeatherAPIResponse struct {
	Current struct {
		TempC float64 `json:"temp_c"`
	} `json:"current"`
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/{cep}", weatherHandler)

	http.Handle("/", r)

	fmt.Println("Servidor rodando na porta 8080")
	http.ListenAndServe(":8080", nil)
}

func weatherHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	cep := vars["cep"]

	matched, err := regexp.MatchString(`^\d{8}$`, cep)
	if err != nil || !matched {
		http.Error(w, "invalid zipcode - format eg.: 00111222", http.StatusUnprocessableEntity)
		return
	}

	resp, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
	if err != nil {
		http.Error(w, "can not find zipcode", http.StatusNotFound)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "can not find zipcode", http.StatusNotFound)
		return
	}

	var viaCEPResponse ViaCEPResponse
	err = json.Unmarshal(body, &viaCEPResponse)
	if err != nil || viaCEPResponse.Erro {
		http.Error(w, "can not find zipcode", http.StatusNotFound)
		return
	}

	resp, err = http.Get("http://api.weatherapi.com/v1/current.json?key=3841b81037a5427eb51191826241702&q=" + viaCEPResponse.Localidade)
	if err != nil {
		http.Error(w, "can not find weather", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "can not find weather", http.StatusInternalServerError)
		return
	}

	var weatherAPIResponse WeatherAPIResponse
	err = json.Unmarshal(body, &weatherAPIResponse)
	if err != nil {
		http.Error(w, "can not find weather", http.StatusInternalServerError)
		return
	}

	tempC := weatherAPIResponse.Current.TempC
	tempF := tempC*1.8 + 32
	tempK := tempC + 273.15

	// Crie a resposta
	weather := Weather{TempC: tempC, TempF: tempF, TempK: tempK}

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(weather)
}
