package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// StringConexaoBanco é a string de conexão com o banco MySQL
	StringConexaoBanco = ""

	// Porta que a api vai rodar
	Porta = 0

	// SecretKey é a chave que vai ser utilizada para assinar o token
	SecretKey []byte
)

// Carregar vai inicializar as variáveis de ambiente
func Carregar() {
	var erro error

	// Importação do pacote godotenv para carregar as informações de ambientes presentes no arquivo .env
	if erro = godotenv.Load(); erro != nil {
		// Se o carregamento falhar, mata o programa
		log.Fatal(erro)
	}

	// Conversão da string de configuração "API_PORT" para inteiro
	Porta, erro = strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil {
		Porta = 9000
	}

	// Configurando a string de conexão com os dados do .env
	StringConexaoBanco = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USUARIO"),
		os.Getenv("DB_SENHA"),
		os.Getenv("DB_NOME"),
	)

	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}
