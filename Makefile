MAIN_PATH = ./cmd/api/main.go
DOCKER_DB_SERVICE = mysql
run:
	go run $(MAIN_PATH)

docker-up-db:
	sudo docker-compose up -d $(DOCKER_DB_SERVICE)
