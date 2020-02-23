all: build

build:
	go build -o app main.go

parcel:
	(cd js; npm run parcel &)

run-with-parcel:
	TODO_DEV_RUN=true ./app

run:
	./app

run-dev: build run-with-parcel


.PHONY: run-dev

