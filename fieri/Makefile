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


.PHONY: run-dev

