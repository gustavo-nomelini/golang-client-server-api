package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type CotacaoResponse struct {
	Bid string `json:"bid"`
}

func main() {
	// Context with timeout for client request - 300ms
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		log.Fatalln("Error creating request:", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalln("Error making request:", err)
	}
	defer resp.Body.Close()

	var cotacao CotacaoResponse
	err = json.NewDecoder(resp.Body).Decode(&cotacao)
	if err != nil {
		log.Fatalln("Error decoding response:", err)
	}

	err = saveToFile(cotacao.Bid)
	if err != nil {
		log.Fatalln("Error saving to file:", err)
	}

	fmt.Println("Cotação do dólar salva com sucesso! Valor:", cotacao.Bid)
}

func saveToFile(bid string) error {
	content := fmt.Sprintf("Dólar: %s", bid)
	return ioutil.WriteFile("cotacao.txt", []byte(content), 0644)
}