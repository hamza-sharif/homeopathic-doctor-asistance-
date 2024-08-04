package postgres

import (
	"fmt"
	"log"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/hamza-sharif/homeopathic-doctor-assistant/config"
	"github.com/hamza-sharif/homeopathic-doctor-assistant/db"
	"github.com/hamza-sharif/homeopathic-doctor-assistant/models"
)

// client struct for mysql client.
type client struct {
	conn gorm.DB
}

func NewClient() (db.Client, error) {
	dbUser := viper.GetString(config.DBUser)
	dbPassword := viper.GetString(config.DBPassword)
	dbHost := viper.GetString(config.DBHost)
	dbPort := viper.GetString(config.DBPort)
	dbName := viper.GetString(config.DBName)

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("database Connection error %s", dsn)

		return nil, err
	}

	if err = conn.AutoMigrate(&models.User{}, &models.Medicine{},
		&models.Patient{}, &models.User{}); err != nil {
		return nil, errors.Wrap(err, "failed to create tables")
	}

	return &client{conn: *conn}, nil
}
