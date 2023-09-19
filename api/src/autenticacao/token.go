package autenticacao

import (
	"api/src/config"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// CriarToken retorna um token assinado com as permissões do usuário
func CriarToken(usuarioID uint64) (string, error) {
	// Map que contém as permissões dentro do token
	permissoes := jwt.MapClaims{}
	permissoes["authorized"] = true
	// expiração do token -> setando para 6 horas a partir de agora - Unix() retorna a quantidade de segundos que passou desde 01/01/1970
	permissoes["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissoes["usuarioId"] = usuarioID

	// secret e assinatura do token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes)
	return token.SignedString([]byte(config.SecretKey)) // secret de 64 bytes utilizando os pacotes do go
}
