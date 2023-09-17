package repositorios

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

// usuários representa um repositório de usuários
type usuarios struct {
	// Struct que recebe o banco
	db *sql.DB
}

// NovoRepositorioDeUsuarios recebe o banco de dados do controller e passa a informação para dentro da struct, que fará a comunicação com as tabelas do banco
func NovoRepositorioDeUsuarios(db *sql.DB) *usuarios {
	return &usuarios{db}
}

// Criar insere um usuário no banco de dados
func (repositorio usuarios) Criar(usuario models.Usuario) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"insert into usuarios (nome, nick, email, senha) values(?, ?, ?, ?)")
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil

}

// Buscar retorna todos os usuários que atendam um filtro de nome ou nick
func (repositorio usuarios) Buscar(nomeOuNick string) ([]models.Usuario, error) {
	// variável que vai receber o valor buscado na url
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick) // %nomeOuNick%

	// query de pesquisa no banco de dados, buscando os usuários por nome ou nick, podendo ser parcial
	linhas, erro := repositorio.db.Query(
		"select id, nome, nick, email, criadoEm from usuarios where nome like ? or nick like ?",
		nomeOuNick, nomeOuNick)

	if erro != nil {
		return nil, erro
	}
	// fechamento do banco de dados
	defer linhas.Close()

	// criação do slice de usuários para armazenar os resultados da pesquisa
	var usuarios []models.Usuario

	// para cada resultado encontrado, inserir dentro do slice de pesquisa com seus respectivos valores
	for linhas.Next() {
		var usuario models.Usuario

		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)
	}

	// retorna o resultado em um slice de struct, sem erro
	return usuarios, nil
}

// BuscarPorID traz um usuário do banco de dados pelo ID
func (repositorio usuarios) BuscarPorID(ID uint64) (models.Usuario, error) {
	linhas, erro := repositorio.db.Query(
		"select id, nome, nick, email, criadoEm from usuarios where id = ?", ID)
	if erro != nil {
		// Não é possível passar somente o nil, necessário passar um usuário vazio
		return models.Usuario{}, erro
	}
	defer linhas.Close()

	var usuario models.Usuario

	if linhas.Next() {
		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return models.Usuario{}, erro
		}
	}

	return usuario, nil
}
