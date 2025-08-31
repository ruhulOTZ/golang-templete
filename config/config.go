package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var configurations Config

type Config struct {
	Version     string
	ServiceName string
	HttpPort    int
	DBName      string
	DBHost      string
	DBUser      string
	DBPassword  string
	DBPort      int
}

func loadConfig() Config {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file:", err)
		os.Exit(1)
	}

	version := os.Getenv("VERSION")
	serviceName := os.Getenv("SERVICE_NAME")
	httpPort := os.Getenv("HTTP_PORT")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbPortStr := os.Getenv("DB_PORT")

	port, err := strconv.ParseInt(httpPort, 10, 64)
	dbPort, dbPortErr := strconv.ParseInt(dbPortStr, 10, 64)

	if err != nil || dbPortErr != nil {
		fmt.Println("Port must be a number")
		os.Exit(1)
	}

	if version == "" || serviceName == "" || httpPort == "" || dbName == "" || dbHost == "" || dbUser == "" || dbPassword == "" || dbPort == 0 {
		fmt.Println("Missing required environment variables")
		os.Exit(1)
	}

	configurations = Config{
		Version:     version,
		ServiceName: serviceName,
		HttpPort:    int(port),
		DBName:      dbName,
		DBHost:      dbHost,
		DBUser:      dbUser,
		DBPassword:  dbPassword,
		DBPort:      int(dbPort),
	}

	return configurations
}

func (c *Config) DatabaseUrl() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d ?sslmode=disable",
		c.DBHost, c.DBUser, c.DBPassword, c.DBName, c.DBPort)
}

func GetConfig() Config {
	return loadConfig()
}
