
# SwordHealth Challenge API

SwordHealth Task Manager.



## Tech Stack

**API:** Golang GIN, Golang GORM

**Notification Service:** RabbitMQ


## Screenshots

![App Screenshot](https://via.placeholder.com/468x300?text=App+Screenshot+Here)


## Deployment

To deploy this project run

```bash
  make docker-build-run
  or
  docker-compose up -d --build
```


## API Reference

There are 2 JSON files in the ./postman directory (postman collection and collection env) that you must import into Postman to test the API.


## Environment Variables

To run this project, you will need to add the following environment variables to your .env file

`DB_USER`

`MYSQL_DATABASE`

`MYSQL_PASSWORD`

`MYSQL_ROOT_PASSWORD`

`MYSQL_ROOT_HOST`

`DB_PORT`

`MIGRATE (You must set 'True' in the first run)`

`RMQ_URL`

`RMQ_QUEUE`


## Performed Task Notification

When a task is performed, a notification event is triggered:
![App Screenshot](https://via.placeholder.com/468x300?text=App+Screenshot+Here)

