package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// response

// pokemon

// pokemonSpecies

var url string = "http://pokeapi.co/api/v2/pokedex/kanto/"

type response struct {
	Id      int       `json:"id"` 
	Name    string    `json:"name"`
	Pokemon []pokemon `json:"pokemon_entries"`
}

type pokemon struct {
	Number int            `json:"entry_number"`
	Spiece pokemonSpecies `json:"pokemon_species"`
}

type pokemonSpecies struct {
	Name string `json:"name"` // for json reading, the field Name will be read from the key name on json
	Url  string `json:"url"`
}

func main() {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	// we need to close the body response (if err == nil) always
	defer resp.Body.Close()

	responseData, err := io.ReadAll(resp.Body) // reads the data but put it in bytes in responseData

	if err != nil {
		log.Fatal(err) // a pretty way of closing the program after printing it
	}

	// fmt.Println(string(responseData))

	var responseObject response

	json.Unmarshal(responseData, &responseObject) // desserializando para objeto

	fmt.Println("nome da requisicao principal: ", responseObject.Name)
	fmt.Println("array de pokemon entries: ", responseObject.Pokemon)
	fmt.Println("id: ", responseObject.Id)

	for _, pokem := range responseObject.Pokemon {
		fmt.Println(pokem.Spiece.Name)
	}

}
