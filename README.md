# test-examen

## Requisitos

* Go 1.17
* Swaggo( go get -u -v github.com/swaggo/swag/cmd/swag )
* Echo Swagger ( go get -u -v github.com/swaggo/echo-swagger )
* MySQL 5.7.x ( docker pull mysql:5.7.33 )

* `export PATH=$(go env GOPATH)/bin:$PATH`
* `swag init -g api/api.go -o api/docs`


## Ambiente Local ( BD basado en Docker )

* BD: `docker run --name test-db -e MYSQL_ROOT_PASSWORD=123456 -d -p 3306:3306 mysql:5.7.33`

## Creación Esquema y Tablas - Carga datos iniciales ( Basado en Docker)

* Copiar scripts dentro del contenedor : `docker cp ./db/scripts/ test-db:/tmp/`
* Eliminación y creación de esquema y tablas : `docker exec -t test-db /bin/sh -c 'mysql -u root -p123456 </tmp/scripts/create-db.sql'`
* Cargar datos de prueba : `docker exec -t test-db /bin/sh -c 'mysql -u root -p123456 </tmp/scripts/basic-data.sql'`

## Compilación y Ejecución

* `go run api/api.go`

## Link Swagger

* `http://localhost:1323/swagger/index.html`

## Ejecución de Tests

* `go test ./...`

## Docker

Comandos para generación de contenedor de API. No es necesario para ambiente local.

* `docker build -t exam-test:1.0 .`
* `docker run -p 1323:1323 --name exam-test exam-test:1.0`
