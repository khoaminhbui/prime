BINARY=prime-service

test: 
	go test -v -cover -covermode=atomic ./...

build:
	go build -o ${BINARY} main.go

docker-start:
	docker build -t prime-service-i .
	docker run -p 8000:8000 --entrypoint /app/prime-service --name=prime-service prime-service-i

docker-clean:
	docker stop prime-service
	docker container rm -f prime-service
	docker rmi prime-service-i