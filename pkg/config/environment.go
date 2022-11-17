package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Environment struct {
	DB_URL     string
	MIGRATE    bool
	SECRET_KEY []byte
	RMQ_URL    string
	RMQ_QUEUE  string
}

func InitEnvironmentConfig() *Environment {
	f, _ := os.Getwd()
	err := godotenv.Load(fmt.Sprintf("%s/.env", f))
	if err != nil {
		log.Println(".env file not found. Using system env variables")
	}

	migrate, err := strconv.ParseBool(os.Getenv("MIGRATE"))
	if err != nil {
		log.Println("MIGRATE not setted on env. Using default value (false).")
		migrate = false
	}

	secret_key := os.Getenv("SECRET_KEY")

	db_url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=Local", os.Getenv("DB_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_ROOT_HOST"), os.Getenv("DB_PORT"), os.Getenv("MYSQL_DATABASE"))
	fmt.Println(db_url)

	return &Environment{
		DB_URL:     db_url,
		MIGRATE:    migrate,
		SECRET_KEY: []byte(secret_key),
		RMQ_URL:    os.Getenv("RMQ_URL"),
		RMQ_QUEUE:  os.Getenv("RMQ_QUEUE"),
	}
}
