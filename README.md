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

