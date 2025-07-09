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
	"strconv"
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

	db := database.NewDb()

	usdbrl_response.USDBRL.ServerTimeStamp = strconv.FormatInt(time.Now().UnixNano(), 10)

	db.Save(usdbrl_response.USDBRL)

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
		return
	}

	w.Header().Set("Content-Type", "application/json")

	response := struct {
		Bid string `json:"bid"`
	}{
		Bid: data.Bid,
	}

	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/cotacao", clientReqHandler)

	fmt.Println("Servidor rodando localmente na porta 8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Erro ao iniciar o servidor: ", err)
	}
}
