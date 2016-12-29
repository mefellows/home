package config

import (
	"log"
	"os"

	"github.com/mefellows/home/db"

	"github.com/hashicorp/logutils"
	"github.com/jinzhu/gorm"
	"github.com/kelseyhightower/envconfig"

	// Auto load .env into env vars
	_ "github.com/joho/godotenv/autoload"
)

// globalConfig contains the current Config for our application
var globalConfig *Config

// Config wraps up the entire application configuration
type Config struct {
	ConnectionString string `envconfig:"DATABASE_URL" default:"sqlite3://app.db"`
	LogLevel         string `envconfig:"LOG_LEVEL" default:"INFO"`
	Port             string `envconfig:"PORT" default:"8000"`
	DB               *gorm.DB
	RedisURL         string `envconfig:"REDIS_URL"`
}

// NewConfig returns an application configuration struct.
func NewConfig() *Config {
	c := Config{}
	envconfig.Process("", &c)

	// Setup logging
	filter := &logutils.LevelFilter{
		Levels:   []logutils.LogLevel{"DEBUG", "INFO", "ERROR"},
		MinLevel: logutils.LogLevel(c.LogLevel),
		Writer:   os.Stderr,
	}
	log.SetOutput(filter)
	log.Println("[DEBUG] config:", &c)

	// Setup DB
	c.DB = db.GetDatabase(c.ConnectionString)
	//db, err := gorm.Open("postgres", "host=myhost user=gorm dbname=gorm sslmode=disable password=mypassword")

	globalConfig = &c

	return &c
}

// GetGlobalConfig gets the current config for the app.
func GetGlobalConfig() *Config {
	return globalConfig
}
