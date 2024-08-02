# go-finance
Projeto fullstack de finanças  utilizando Golang, SQLC, Next, Postgres e React. Aplicando-se CRUD, login, autenticação, download de CSV e mais...


Golang Migrate: https://github.com/golang-migrate/migrate
Golang Migrate CMD para instalação: https://github.com/golang-migrate/migrate/tree/master/cmd/migrate


Comandos:
````
migrate create -ext sql -dir db/migration -seq initial_tables
docker pull postgres:14-alpine
docker images
docker run --name postgres -p 5432:5432 -e POSTGRES_PASSWORD=postgres -d postgres:14-alpine
docker exec -it postgres psql -U postgres
docker start postgres
docker exec -it postgres /bin/sh

docker exec -it postgres /bin/sh
/ # pwd
/ # ls -l
/ # createdb --username=postgres --owner=postgres go_finance


migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/go_finance?sslmode=disable" -verbose up

/gera código para inserir, listar dados
docker pull kjconroy/sqlc
docker run --rm -v "$(pwd)":/src -w /src kjconroy/sqlc generate

go mod init github.com/Jennyads/go-finance
go mod tidy

````

up: adicionar algo
down: remover algo do projeto



## SQLC:
O sqlc é uma ferramenta que gera código fortemente tipado a partir de consultas SQL. Você escreve suas consultas em SQL, e o sqlc compila essas consultas em funções e estruturas de dados na linguagem de programação de sua escolha (como Go, Kotlin, Python, ou TypeScript). Isso permite interações seguras e eficientes com o banco de dados, aproveitando a segurança de tipo e melhorando a manutenibilidade do código.

````
-- query/get_user.sql
SELECT id, name, email FROM users WHERE id = $1;
````

````
package main

import (
    "database/sql"
    "log"

    "your_project/internal/db"
)

func main() {
    conn, err := sql.Open("postgres", "your_connection_string")
    if err != nil {
        log.Fatal(err)
    }

    queries := db.New(conn)
    user, err := queries.GetUser(ctx, 1)
    if err != nil {
        log.Fatal(err)
    }

    log.Println(user.Name)
}

````


