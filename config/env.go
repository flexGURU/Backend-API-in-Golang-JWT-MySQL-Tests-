package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost string
	Port string
	DBUser string
	DBPassword string
	DBAdrress string
	DBName string
}

var Envs = initiateConfig()

func initiateConfig() Config{
	godotenv.Load()
	return Config{
		DBHost: getEnv("PUBLIC_HOST", "localhost"),
		Port: getEnv("PORT", "5432"),
		DBUser: getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", "secret"),
		DBAdrress: fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "3306") ),
		DBName: getEnv("DB_NAME", "demogo"),
	
	}
}

func getEnv(key, fallback string) string  {

	value, ok := os.LookupEnv(key)
	if ok {
		return value
	}

	return fallback

	
}


