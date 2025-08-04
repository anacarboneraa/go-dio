package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Response struct {
	Nome string `json:"name"`
	Pokemon []Pokemon `json:"pokemon_entries"`
}

type Pokemon struct {
	Numero int `json:"entry_number"`
	Especie PokemonSpecies `json:"pokemon_species"`
}

type PokemonSpecies struct {
	Nome string `json:"name"`
}

func main() {
	res, err := http.Get("https://pokeapi.co/api/v2/pokedex/kanto/") // mapeia

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	defer res.Body.Close()

	responseData, err := io.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(responseData))

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	fmt.Println(responseObject.Nome)
	fmt.Println(responseObject.Pokemon)

	for _, p := range responseObject.Pokemon {
		fmt.Println(p.Especie.Nome)
	}
}


