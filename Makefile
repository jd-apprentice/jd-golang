all: lint test imports build start

start:
	./bin/app_name

dev: path
	gowatch

fmt:
	@if [ -n "$$(go fmt ./src)" ]; then \
        echo "Code is not properly formatted"; \
        exit 1; \
    fi

pre-commit:
	pre-commit run --all

imports: path
	goimports -w ./src

lint: path
	golangci-lint run

lint-fix: path
	golangci-lint run --fix

clear:
	rm -rf ./bin

## https://github.com/golang-standards/project-layout/issues/113#issuecomment-1336514449
build: clear fmt
	GOARCH=amd64 go build -o ./bin/app_name ./src/main.go

build-arm: clear fmt
	GOARCH=arm64 go build -o ./bin/app_name ./src/main.go

test:
	go test -v ./tests

path:
	export PATH=$$PATH:$$HOME/go/bin

### Kubernetes && Docker

## https://www.digitalocean.com/community/tutorials/how-to-use-minikube-for-local-kubernetes-development-and-testing
k8s-up:
	minikube start

k8s-down:
	minikube stop

k8s-apply:
	kubectl apply -f ./infra/kubernetes

pods:
	kubectl get pods -A

compose-up:
	docker compose -f ./compose.yaml up -d --build

compose-down:
	docker compose -f ./compose.yaml down