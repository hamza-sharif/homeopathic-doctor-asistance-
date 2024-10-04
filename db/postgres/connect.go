package postgres

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"os"

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
	Fee  int
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

	if err = conn.AutoMigrate(&models.User{}, &models.Medicine{}, &models.Patient{}, &models.User{}, &models.Price{}); err != nil {
		return nil, errors.Wrap(err, "failed to create tables")
	}

	fee, err := addMedicine(context.TODO(), conn)

	return &client{conn: *conn, Fee: fee}, nil
}

func addMedicine(ctx context.Context, db *gorm.DB) (int, error) {

	fee := models.Price{}
	err := db.Model(&models.Price{}).First(&fee).Error
	if err != nil {
		fmt.Println("database not have data on price table %s", err)
		fee.Fee = viper.GetInt(config.PatientsFee)
		if errfee := db.Create(&fee).Error; errfee != nil {
			return 0, err
		}
	}

	var med int64
	db.Model(&models.Medicine{}).Count(&med)

	if med < 100 {
		if _, err := db.ConnPool.ExecContext(ctx, fmt.Sprintf("TRUNCATE TABLE medicines")); err != nil {
			return 0, err
		}

		// Open CSV file
		file, err := os.Open("medicines.csv")
		if err != nil {
			fmt.Println("Failed to open file:", err)
			return 0, err
		}

		defer file.Close()

		// Read the CSV file
		reader := csv.NewReader(file)
		records, err := reader.ReadAll()
		if err != nil {
			fmt.Println("Error reading CSV file:", err)
			return 0, err
		}

		// Iterate over the CSV rows and insert data into the database
		for _, row := range records {
			user := models.Medicine{
				Description: row[0],
				Name:        row[1],
			}

			// Insert user into the database
			result := db.Create(&user)
			if result.Error != nil {
				fmt.Println("Error inserting record:", result.Error)
			}
		}

		fmt.Println("Data successfully inserted!")
	}

	return fee.Fee, nil
}
