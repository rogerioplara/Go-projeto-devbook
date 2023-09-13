package models

import "time"

// Usuário representa um usuário utilizando a aplicação
type Usuario struct {
	ID       uint64    `json:"id,omitempty"` // se o id estiver em branco, rejeita
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"CriadoeEm,omitempty"`
}
