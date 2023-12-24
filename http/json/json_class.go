package main

import ( 
	"fmt"
	"os"
	// for json :
	"io"
	"encoding/json"
	// "strconv"
)

// type Objeto1 struct {
// 	Objetos []string `json: "objeto"` // identificador de leitura para json
// 	Idade int `json: "idade"`
// }

type Usuario struct {
	Nome string `json : "nome"`
	Idades []int `json : "idades"`
}

type Usuarios struct {
	Usuarios []Usuario `json: "usuarios"`
}

func main() {
	jsonFile, err := os.Open("file.json")

	if err != nil {
		fmt.Println("Arquivo nao existente")
		panic(err)
	}

	fmt.Println("Arquivo aberto com sucesso")

	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)

	if err != nil {
		panic(err)
	}

	// criar uma estrutura para cara objeto dentro do json, incluindo o principal
	// var objeto1 Objeto1
	var usuarios Usuarios

	// json.Unmarshal(byteValue, &objeto1)
	json.Unmarshal(byteValue, &usuarios)

	// strconv.Itoa(10) // returns "10"

	for i := 0 ; i < len(usuarios.Usuarios); i++ {
		fmt.Println("Nome: ", usuarios.Usuarios[i].Nome)
		fmt.Println("Idades: ", usuarios.Usuarios[i].Idades)
	}
}