package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Port          string
	DBDriver      string
	DBUrl         string
	JWTSecretKey  string
	JWTExpiration time.Duration
}

var AppConfig *Config

func LoadConfig(path ...string) {
	envPath := ".env"
	if len(path) > 0 {
		envPath = path[0]
	}
	err := godotenv.Load(envPath)
	if err != nil {
		log.Println("No .env file found or error loading .env, using environment variables directly. Error: ", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	dbDriver := os.Getenv("DB_DRIVER")
	dbURL := os.Getenv("DB_URL")
	jwtSecret := os.Getenv("JWT_SECRET_KEY")
	jwtExpHoursStr := os.Getenv("JWT_EXPIRATION_HOURS")

	jwtExpHours, err := strconv.Atoi(jwtExpHoursStr)
	if err != nil || jwtExpHours <= 0 {
		jwtExpHours = 72
	}

	AppConfig = &Config{
		Port:          port,
		DBDriver:      dbDriver,
		DBUrl:         dbURL,
		JWTSecretKey:  jwtSecret,
		JWTExpiration: time.Duration(jwtExpHours) * time.Hour,
	}

	if AppConfig.DBUrl == "" || AppConfig.JWTSecretKey == "" {
		log.Fatal("Database URL and JWT Secret Key must be set in environment variables or .env file")
	}
}
