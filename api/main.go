package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {

	// Carregar as configurações
	config.Carregar()
	// Recebendo algum tipo de retorno da configuração
	r := router.Gerar()

	fmt.Printf("Escutando na porta %d", config.Porta)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))

}
