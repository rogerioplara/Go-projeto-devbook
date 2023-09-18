package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositorios"
	"api/src/respostas"
	"encoding/json"
	"io"
	"net/http"
)

// Login é responsável por autenticar um usuário na API
func Login(w http.ResponseWriter, r *http.Request) {
	// Recebe o que foi passado e armazena na variável, trata o erro se houver
	corpoRequisicao, erro := io.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	// Criação da struct de usuário, passando os dados recebidos em Json
	var usuario models.Usuario
	if erro := json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	// Abertura e fechamento da conexão com a base de dados
	db, erro := database.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	// Criação do repositório de usuários, passando o banco de dados
	repositorio := repositorios.NovoRepositorioDeUsuarios(db)

	// Buscando o usuário por email
	usuarioSalvoNoBanco := repositorio.BuscarPorEmail()
}
