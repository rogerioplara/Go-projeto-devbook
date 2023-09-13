# Aplicação Devbook

Serão desenvolvidos 2 componentes para a aplicação:
- API (<i>Back-End</i>)
- Web App (<i>Front-End</i>)

## Estrutura da aplicação
### USUÁRIOS
```
- CRUD
- Seguir outro usuário
- Parar de seguir outro usuário
- Buscar todos os usuários que segue
- Buscar todos os usuários que são seguidos
- Atualizar senha

Tabelas
- Usuários
- Seguidores
```

### PUBLICAÇÕES
```
- CRUD
- Buscar publicações de acordo com os usuários que segue
- Curtir publicações
```

### PACOTES

Os pacotes da aplicação podem ser divididos em dois tipos:


#### Pacotes Principais
```
- Main
- Router
- Controllers
- Modelos
- Repositórios
```

#### Pacotes Auxiliares
```
- Config
- Banco (abrir conexão)
- Autenticação
- Middleware
- Segurança
- Respostas
```