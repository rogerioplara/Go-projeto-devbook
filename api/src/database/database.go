package database

import (
	"api/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // Driver
)

// Conectar abre a conexão com o banco de dados e a retorna
func Conectar() (*sql.DB, error) {
	db, erro := sql.Open("mysql", config.StringConexaoBanco)
	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		// Se ocorrer erro, fecha a conexão
		db.Close()
		return nil, erro
	}

	// Se não houver erro, retorna a conexão aberta sem erro
	return db, nil
}
