BINARY=prime-service

test: 
	go test -v -cover -covermode=atomic ./...

run-standalone:
	go run main.go

build:
	go build -o ${BINARY} main.go

docker-build:
	docker build -t prime-service-i .
	docker run -p 8000:8000 --entrypoint /app/prime-service --name=prime-service prime-service-i

docker-clean:
	docker stop prime-service
	docker container rm -f prime-service
	docker rmi prime-service-i