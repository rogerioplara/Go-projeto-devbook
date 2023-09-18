package respostas

import (
	"encoding/json"
	"log"
	"net/http"
)

/*
Esse pacote cuidará do retorno de respostas mais adequadas
*/

// JSON é a função de resposta -> recebe o statusCode, inserir esse statusCode no header. Por fim, pegar os dados que são genéricos e transformar para JSON
func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {
	// definir como será recebido o valor -> content type
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	// Verifica se realmente os dados existem e realiza o encode, se não fizer gera erro NoContet
	if dados != nil {
		if erro := json.NewEncoder(w).Encode(dados); erro != nil {
			log.Fatal(erro)
		}
	}

}

// Erro retorna o erro
func Erro(w http.ResponseWriter, statusCode int, erro error) {
	JSON(w, statusCode, struct {
		Erro string `json:"erro"`
	}{
		Erro: erro.Error(),
	})
}
