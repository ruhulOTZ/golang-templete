package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var configurations *Config

type Config struct {
	Version      string
	ServiceName  string
	HttpPort     int
	JwtSecretKey string
}

func loadConfig() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file:", err)
		os.Exit(1)
	}

	version := os.Getenv("VERSION")
	serviceName := os.Getenv("SERVICE_NAME")
	httpPort := os.Getenv("HTTP_PORT")

	if version == "" || serviceName == "" || httpPort == "" {
		fmt.Println("Missing required environment variables")
		os.Exit(1)
	}

	port, err := strconv.ParseInt(httpPort, 10, 64)
	if err != nil {
		fmt.Println("Port must be a number")
		os.Exit(1)
	}

	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	if jwtSecretKey == "" {
		fmt.Println("JWT secret key is required")
		os.Exit(1)
	}

	configurations = &Config{
		Version:      version,
		ServiceName:  serviceName,
		HttpPort:     int(port),
		JwtSecretKey: jwtSecretKey,
	}
}

func GetConfig() *Config {
	if configurations == nil {
		loadConfig()
	}
	return configurations
}
