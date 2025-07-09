package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"pos_go_desafio_1/internal/domain"
	"time"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)

	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)

	if err != nil {
		panic(err)
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println("Erro ao chamar o servidor: ", err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		fmt.Println("Erro ao ler o corpo da resposta: ", err)
	}

	var cotacao domain.Cotacao

	if err = json.Unmarshal(body, &cotacao); err != nil {
		fmt.Println("Erro ao realizar o unmarshal!")
	}

	file, err := os.OpenFile("cotacao.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Erro ao criar ou abrir o arquivo: ", err)
	}
	defer file.Close()

	line := fmt.Sprintf("Dólar: %s\n", string(cotacao.Bid))
	_, err = file.Write([]byte(line))
	if err != nil {
		fmt.Println("Erro ao escrever dados no arquivo cotacao.txt: ", err)
	}

	fmt.Println("Resposta do servidor: " + string(body))

}
