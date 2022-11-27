.PHONY: install test build serve clean pack deploy ship
TAG?=$(shell git rev-list HEAD --max-count=1 --abbrev-commit)

export TAG

install:
	go get .

build: install
	go build -ldflags "-X main.version=$(TAG)" -o ./bin/programming-challenge .

serve: build
	./bin/programming-challenge serve

dev:
	GOOS=linux make build
	docker build -t localhost:5000/programming-challenge .
	docker push localhost:5000/programming-challenge