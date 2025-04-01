include .env

all: build

build: generate test
	if [ ! -d "build" ]; then mkdir build; fi
	echo "Building main.go"
	go build -o build/main ./cmd/main.go

run: build
	echo "Now running main"
	./build/main

test:
	go test ./...

test-k8s:
	GRPC_HOST=10.0.0.108:30003 go test ./... --ginkgo.label-filter "integration"

install-go-deps:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

clean:
	rm -r ./build

generate:
	protoc --go_out=.  --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative "api/easy-housing.proto"

deploy: docker-build docker-push k8s-rollout-all-updates

docker-build: docker-build-server

docker-push: docker-push-server

docker-build-server:
	docker build --platform=linux/amd64 -t 100.69.236.43:32000/easy-housing-server:latest -f Dockerfile .

docker-push-server:
	docker push 100.69.236.43:32000/easy-housing-server:latest

k8s-rollout-all-updates: k8s-rollout-server-updates

k8s-rollout-server-updates:
	kubectl rollout restart deployment -n mtg-mana-sim-app easy-housing-server-deployment
