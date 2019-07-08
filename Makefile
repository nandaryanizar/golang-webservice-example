test:
	go test ./...

build:
	go build -o ./bin/golang-webservice-example ./cmd/...

run: build
	./bin/golang-webservice-example
	
start:
	docker-compose -f ./docker-compose.yml up --build --abort-on-container-exit

destroy: 
	docker-compose -f ./docker-compose.yml down --volumes