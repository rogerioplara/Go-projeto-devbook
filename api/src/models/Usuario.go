package models

import (
	"api/src/seguranca"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

// Usuário representa um usuário utilizando a aplicação
type Usuario struct {
	ID       uint64    `json:"id,omitempty"` // se o id estiver em branco, rejeita
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"CriadoeEm,omitempty"`
}

func (usuario *Usuario) Preparar(etapa string) error {
	// Se houver erro na validação, retorna o erro
	if erro := usuario.validar(etapa); erro != nil {
		return erro
	}

	// Se não houver erro, formata os campos e retorna o erro nil. Realiza o hash de senha
	if erro := usuario.formatar(etapa); erro != nil {
		return erro
	}

	return nil
}

// Validar verifica se os campos passados não estão vazios
func (usuario *Usuario) validar(etapa string) error {
	if usuario.Nome == "" {
		return errors.New("o nome é obrigatório")
	}

	if usuario.Nick == "" {
		return errors.New("o nick é obrigatório")
	}

	if usuario.Email == "" {
		return errors.New("o email é obrigatório e não pode estar em branco")
	}

	// Validação do email com o package github.com/badoux/checkmail
	if erro := checkmail.ValidateFormat(usuario.Email); erro != nil {
		return errors.New("o e-mail inserido é inválido")
	}

	// Condição etapa verifica se a etapa é de cadastro, impedindo assim o cadastro da senha vazia
	// Essa condição serve para realizar a atualização do cadastro e senha por rotas diferentes
	if etapa == "cadastro" && usuario.Senha == "" {
		return errors.New("a senha é obrigatória")
	}

	return nil
}

// Formatar retira os espaços do início e do fim dos campos Nome, Nick e Email
func (usuario *Usuario) formatar(etapa string) error {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)

	// Validação da senha
	if etapa == "cadastro" {
		senhaComHash, erro := seguranca.Hash(usuario.Senha)
		if erro != nil {
			return erro
		}

		usuario.Senha = string(senhaComHash)
	}

	return nil
}
