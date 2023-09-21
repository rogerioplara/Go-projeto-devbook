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

// Atualizar altera as informações de um usuário no banco de dados - recebe o ID e o struct de usuário, retorna um erro se existir
func (repositorio usuarios) Atualizar(ID uint64, usuario models.Usuario) error {
	// Statement SQL
	statement, erro := repositorio.db.Prepare("update usuarios set nome = ?, nick = ?, email = ? where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	// Execução do statement -> retorna 2 valores, onde o primeiro será ignorado
	if _, erro = statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, ID); erro != nil {
		return erro
	}

	return nil
}

// Deletar apaga as informações de um usuário no banco de dados
func (repositorio usuarios) Deletar(ID uint64) error {
	statement, erro := repositorio.db.Prepare("delete from usuarios where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(ID); erro != nil {
		return erro
	}

	return nil
}

// BuscarPorEmail busca um usuário por email e retorna seu ID e Senha com hash
func (repositorio usuarios) BuscarPorEmail(email string) (models.Usuario, error) {
	// Query string
	linha, erro := repositorio.db.Query("select id, senha from usuarios where email = ?", email)
	if erro != nil {
		return models.Usuario{}, erro
	}
	defer linha.Close()

	var usuario models.Usuario

	// Se existir, faz o scan e já trata o erro
	if linha.Next() {
		if erro = linha.Scan(&usuario.ID, &usuario.Senha); erro != nil {
			return models.Usuario{}, erro
		}
	}

	return usuario, nil
}

// Seguir permite que um usuário siga outro
func (repositorio usuarios) Seguir(usuarioID, seguidorID uint64) error {
	statement, erro := repositorio.db.Prepare("insert ignore into seguidores (usuario_id, seguidor_id) values (?, ?)")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(usuarioID, seguidorID); erro != nil {
		return erro
	}

	return nil
}

// PararDeSeguir permite que um usuário pare de seguir outro
func (repositorio usuarios) PararDeSeguir(usuarioID, seguidorID uint64) error {
	statement, erro := repositorio.db.Prepare("delete from seguidores where usuario_id = ? and seguidor_id = ?")
	if erro != nil {
		return erro
	}

	if _, erro = statement.Exec(usuarioID, seguidorID); erro != nil {
		return erro
	}
	defer statement.Close()

	return nil
}

// BuscarSeguidores traz todos os seguidores de um usuário
func (repositorio usuarios) BuscarSeguidores(usuarioID uint64) ([]models.Usuario, error) {
	linhas, erro := repositorio.db.Query(`
		SELECT u.id, u.nome, u.nick, u.email, u.criadoEm
		FROM usuarios u INNER JOIN seguidores s ON u.id = s.seguidor_id 
		WHERE s.usuario_id = ?
	`, usuarioID)
	if erro != nil {
		return nil, erro
	}
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

// BuscarSeguidores traz todos que um determinado usuário está seguindo
func (repositorio usuarios) BuscarSeguindo(usuarioID uint64) ([]models.Usuario, error) {
	linhas, erro := repositorio.db.Query(`
		SELECT u.id, u.nome, u.nick, u.email, u.criadoEm
		FROM usuarios u INNER JOIN seguidores s ON u.id = s.usuario_id 
		WHERE s.seguidor_id = ?
	`, usuarioID)
	if erro != nil {
		return nil, erro
	}
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

// BuscarSenha traz a senha de um usuário pelo ID
func (repositorio usuarios) BuscarSenha(usuarioID uint64) (string, error) {
	linha, erro := repositorio.db.Query("select senha from usuarios where id = ?", usuarioID)
	if erro != nil {
		return "", erro
	}
	defer linha.Close()

	var usuario models.Usuario

	if linha.Next() {
		if erro = linha.Scan(&usuario.Senha); erro != nil {
			return "", erro
		}
	}

	return usuario.Senha, nil

}

// Atualizar senha altera a senha do usuário
func (repositorio usuarios) AtualizarSenha(usuarioID uint64, senha string) error {
	statement, erro := repositorio.db.Prepare("update usuarios set senha = ? where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(senha, usuarioID); erro != nil {
		return erro
	}

	return nil
}
