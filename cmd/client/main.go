package main

import (
	"fmt"
	"io"
	"net/http"
	"pos_go_desafio_1/pkg/database"
)

func main() {

	db := database.NewDb()

	resp, err := http.Get("http://localhost:8080")
	//
	if err != nil {
		fmt.Println("Erro na requisição: ", err)
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Erro ao ler o corpo da resposta: ", err)
	}

	fmt.Println("Resposta do servidor: ")
	fmt.Println(string(body))
}
