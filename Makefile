.PHONY: build
build:rm
	docker build -t go-app-img .

.PHONY: rm
rm:
	docker rm go-app-container

.PHONY: up
up:
	docker run -d -p 3333:3000 --name go-app-container go-app-img