package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost string
	Port string
	DBUser string
	DBPassword string
	DBAdrress string
	DBName string
	JWTDuration int64	 
	JWTSecret string
}

var Envs = initiateConfig()

func initiateConfig() Config{
	godotenv.Load()
	return Config{
		DBHost: getEnv("PUBLIC_HOST", "localhost"),
		Port: getEnv("PORT", "5432"),
		DBUser: getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", "secret"),
		DBAdrress: fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "5432") ),
		DBName: getEnv("DB_NAME", "demogo"),
		JWTDuration: getEnvInt("JWTDuration", 3600*24*7),
		JWTSecret: getEnv("JWTSecret", "NOSECRET"),
	
	}
}

func getEnv(key, fallback string) string  {

	value, ok := os.LookupEnv(key)
	if ok {
		return value
	}

	return fallback

}

func getEnvInt(key string, fallback int64) int64 {
	 
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				return fallback
			}
			return i
		}

		return fallback
	}

	


