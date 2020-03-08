all: build

build:
	go build -o app main.go

node-get-deps:
	docker run -v "$(shell pwd)/js:/home/node/" node:10.19.0-alpine3.11 /bin/sh -c "cd /home/node && npm install"

parcel:
	docker run -it  -v "$(shell pwd)/js:/home/node/" -p 1234:1234 node:10.19.0-alpine3.11 /bin/sh -c "cd /home/node && npm run parcel"

run-with-parcel:
	TODO_DEV_RUN=true ./app

run:
	./app

run-dev: build run-with-parcel

postgres:
	docker run -it -e "POSTGRES_USER=task_user" -e "POSTGRES_DB=task_db" -e "POSTGRES_PASSWORD=Muonoo4p" -v "$(shell pwd)/db:/docker-entrypoint-initdb.d/" -p 5433:5432 postgres:12-alpine 

.PHONY: run-dev

