
# SwordHealth Challenge API

SwordHealth Task Manager.



## Tech Stack

**API:** Golang GIN, Golang GORM

**Notification Service:** RabbitMQ


## Screenshots

![Dev Screenshot](https://user-images.githubusercontent.com/36938063/202851589-1aea18fa-b5fe-43d1-95b7-588f4b23f9ed.png)



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

![event-notification](https://user-images.githubusercontent.com/36938063/202851459-b400333d-6ea2-4b20-8d77-51566e4b1298.png)

