package middlewares

import (
	"fmt"
	"log"
	"net/http"
)

/*
Camada que fica entre a requisição e a resposta
Muito utilizado quando existe alguma função que deve ser aplicado para todas as rotas
Por exemplo autenticação, criação de logs
*/

// Logger escreve informações da requisição no terminal
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}

// Autenticar verifica se o usuário que está fazendo a requisição está autenticado
func Autenticar(next http.HandlerFunc) http.HandlerFunc {
	// recebe uma função e retorna outra função, que atende aos requisitos da rota
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Autenticando...")
		next(w, r)
	}

}
