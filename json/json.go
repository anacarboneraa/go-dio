package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Users struct {
	Users []User `json:"usuarios"`
}

type User struct {
	Nome   string `json:"Nome"`
	Tipo   string `json:"Tipo"`
	Idade  int    `json:"Idade"`
	Social Social `json:"Social"`
}

type Social struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
}

func main() {
	jsonFile, err := os.Open("usuarios.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Arquivo aberto com sucesso")
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	var usuarios Users

	err = json.Unmarshal(byteValue, &usuarios)
	if err != nil {
		fmt.Println("Erro ao fazer Unmarshal:", err)
		return
	}

	for i := 0; i < len(usuarios.Users); i++ {
		fmt.Println("Usuario Tipo: " + usuarios.Users[i].Tipo)
		fmt.Println("Usuario Idade: " + strconv.Itoa(usuarios.Users[i].Idade))
		fmt.Println("Usuario Nome: " + usuarios.Users[i].Nome)
		fmt.Println("Facebook: " + usuarios.Users[i].Social.Facebook)
		fmt.Println("Twitter: " + usuarios.Users[i].Social.Twitter)
	}
}
