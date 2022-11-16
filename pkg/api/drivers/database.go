package drivers

import (
	"SwordHealth/pkg/api/models"
	"SwordHealth/pkg/config"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type DatabaseDriver struct {
	dbUrl     string
	migration bool
	DB        *gorm.DB
}

func NewDBDriver() *DatabaseDriver {
	env := config.InitEnvironmentConfig()

	return &DatabaseDriver{
		dbUrl:     env.DB_URL,
		migration: env.MIGRATE,
	}
}

func (dbDriver *DatabaseDriver) Connect() {
	log.Println(dbDriver.dbUrl)
	db, err := gorm.Open("mysql", dbDriver.dbUrl)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database Connected")

	dbDriver.DB = db
}

func (dbDriver *DatabaseDriver) Close() {
	dbDriver.DB.Close()
}

func (dbDriver *DatabaseDriver) Migrate() {
	dbDriver.DB.Debug().AutoMigrate(
		&models.Manager{},
		&models.Technician{},
		&models.Task{},
	)
}
