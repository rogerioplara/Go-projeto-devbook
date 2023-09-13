package rotas

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Criação de uma struct para as rotas
// Independente da rota, o que elas fazem é igual

// Rota é a struct que representa todas as rotas da API
type Rota struct {
	URI                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}

// Configurar recebe o router sem nenhuma rota dentro e configura todas as rotas e retorna o router pronto
func Configurar(r *mux.Router) *mux.Router {
	rotas := rotasUsuarios

	// Insere cada rota já definida dentro do router
	for _, rota := range rotas {
		r.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)
	}

	return r
}
