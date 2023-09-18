package seguranca

import "golang.org/x/crypto/bcrypt"

/*
Utilização do bcrypt para criar o hash de senha

golang.org/x/crypto/bcrypt
*/

// Hash recebe a senha como string e coloca um hash nela. Retorna um slice de bytes e um erro
func Hash(senha string) ([]byte, error) {
	// Função para realizar o hash
	// Deve ser passado na função um slice de bytes (realizar casting) e o custo do hash
	return bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
}

// VerificarSenha compara uma senha com um hash e retorna se elas são iguais
func VerificarSenha(senhaComHash, senhaString string) error {
	return bcrypt.CompareHashAndPassword([]byte(senhaComHash), []byte(senhaString))
}
