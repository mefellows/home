package db

import (
	"fmt"
	"log"
	"regexp"

	// Postgres driver
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/mefellows/home/models"
)

func init() {
}

// Manager is an abstraction to get different DB implementations
type Manager interface {
	Setup(connectionString string) *gorm.DB
}

// GetDatabase returns a database handler for a given type.
func GetDatabase(connectionString string) (db *gorm.DB) {
	var err error
	adapter, gormString := ConnectionStringToGormString(connectionString)
	if db, err = gorm.Open(adapter, gormString); err != nil {
		log.Fatal("could not create database connection:", err)
	}

	log.Println("[DEBUG] setting default table name to singular")
	db.SingularTable(true)

	log.Println("[DEBUG] setting auto migrate on models")
	db.AutoMigrate(&models.List{}, &models.Item{})
	return
}

// Convert Postgres connection string to Gorm string
func ConnectionStringToGormString(connectionString string) (adapter string, gormString string) {

	r := regexp.MustCompile(`([a-zA-Z]+)\:\/\/([a-zA-Z0-9_]+)\:([a-zA-Z0-9_]+)\@([0-9a-zA-Z_-\.]+)\:([0-9]{4,5})/([a-zA-Z0-9_]+)(.*)`)
	matches := r.FindStringSubmatch(connectionString)

	return matches[1], fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%s sslmode=disable", matches[4], matches[2], matches[6], matches[3], matches[5])
}
