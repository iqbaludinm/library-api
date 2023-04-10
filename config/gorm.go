package config

import (
	"fmt"
	"os"

	"github.com/iqbaludinm/library-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	// "gorm.io/gorm/schema"
)

type Gorm struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string

	DB *gorm.DB
}

type GormDb struct {
	*Gorm
}

var (
	GORM *GormDb
)

func InitGorm() error {
	GORM = new(GormDb)

	GORM.Gorm = &Gorm{
		Username: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
		Database: os.Getenv("POSTGRES_DBNAME"),
	}

	// connect to database
	err := GORM.Gorm.OpenConnection()
	if err != nil {
		return err
	}

	return nil
}

func (g *Gorm) OpenConnection() error {
	// init dsn
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", g.Host, g.Port, g.Username, g.Password, g.Database)

	// if any schema in my db
	dbConnection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		// NamingStrategy: schema.NamingStrategy{
		// 	TablePrefix:   "library.",
		// 	SingularTable: false,
		// },
	})

	// if without schema
	// dbConnection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	g.DB = dbConnection

	err = g.DB.AutoMigrate(models.Book{})
	if err != nil {
		return err
	}

	fmt.Println("Successfully connected to database with Gorm")

	return nil
}
