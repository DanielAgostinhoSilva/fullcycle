package main

import (
	"fmt"
	"net/http"
)

func main() {

	client := &http.Client{}
	for i := 0; i < 20; i++ {
		req, err := http.NewRequest("GET", "http://localhost:8080", nil)

		if err != nil {
			fmt.Println("Falha ao criar a solicitação:", err)
			return
		}

		req.Header.Add("API_KEY", "12345")

		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Erro ao fazer a solicitação:", err)
			return
		}
		defer resp.Body.Close()
		fmt.Println("Resposta recebida, código de status:", resp.StatusCode)

	}
}
