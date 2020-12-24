docker-start:
	docker build -t prime-web-client-i .
	docker run -p 3000:3000 --name=prime-web-client prime-web-client-i

docker-clean:
	docker stop prime-web-client
	docker container rm -f prime-web-client
	docker rmi prime-web-client-i