BINARY=engine

test: 
	go test -v -cover -covermode=atomic ./...

run-standalone:
	go run main.go

docker:
	docker build -t prime-service-i .

docker-run:
	docker run -d -it --name=prime-service prime-service-i