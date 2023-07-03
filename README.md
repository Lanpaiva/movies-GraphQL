## Go GraphQL

Este é o backend da aplicação, onde foi desenvolvida toda a lógica do negócio. Responsável por fornecer os endpoints e os algorítimos da API GraphQL.

## Pré-requisitos

- Go (versão 1.20)

- Banco de dados SQLite

## Instalação

- Clone este repositório.
- Execute o comando `go mod tidy` para baixar as dependências.
- Execute o comando `go run server.go` para iniciar o servidor.
- Certifique-se de ter o banco de dados SQLite intalado e configurado corretamente.

## Uso

- Playground GraphQL: `http://localhost:8080/`
- Endpoint GraphQL: `http://localhost:8080/query`

## Observações

**1 -** Inicie o servidor da API GraphQL antes de iniciá-lo
no app React.JS.
