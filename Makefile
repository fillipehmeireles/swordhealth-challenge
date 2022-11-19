MAIN_PATH = ./cmd/main.go
DOCKER_DB_SERVICE = mysql
DOCKER_MQ_SERVICE = rabbit-mq

run:
	go run $(MAIN_PATH)
test:
	go test ./... -v --cover

docker-up-db:
	sudo docker-compose up -d $(DOCKER_DB_SERVICE)

docker-up-mq:
	sudo docker-compose up -d $(DOCKER_MQ_SERVICE)

docker-build-run:
	sudo docker-compose build && sudo docker-compose up -d


