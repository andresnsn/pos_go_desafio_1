package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"pos_go_desafio_1/internal/domain"
	"pos_go_desafio_1/pkg/database"
	"time"
)

func getCotation() (*domain.USDBRL, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)

	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)

	if err != nil {
		panic(err)
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println("Erro ao chamar a API de cotação: ", err)
	}

	defer res.Body.Close()

	var usdbrl_response domain.USDBRLResponse

	err = json.NewDecoder(res.Body).Decode(&usdbrl_response)

	if err != nil {
		fmt.Println("Erro ao converter para JSON: ", err)
	}

	ctxDb, cancelDb := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancelDb()

	db := database.NewDb()

	err = db.Save(ctxDb, usdbrl_response.USDBRL)

	if err != nil {
		fmt.Println("Erro ao salvar no banco de dados: ", err)
	}

	io.Copy(os.Stdout, res.Body)

	return &usdbrl_response.USDBRL, nil
}

func clientReqHandler(w http.ResponseWriter, r *http.Request) {
	data, err := getCotation()

	if err != nil {
		http.Error(w, "Erro ao buscar dados", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func main() {
	http.HandleFunc("/", clientReqHandler)

	fmt.Println("Servidor rodando localmente na porta 8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Erro ao iniciar o servidor: ", err)
	}
}
