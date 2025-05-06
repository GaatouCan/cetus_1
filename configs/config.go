package configs

import (
	"fmt"
	"os"
	"sync"
)

type Config struct {
	DBHost string
	DBPort string
	DBUser string
	DBPass string
	DBName string

	JWTToken []byte
}

var (
	instance *Config
	once     sync.Once
)

func GetConfig() *Config {
	once.Do(func() {
		instance = loadConfig()
	})
	return instance
}

func loadConfig() *Config {
	return &Config{
		DBHost:   getEnv("DB_HOST", "localhost"),
		DBPort:   getEnv("DB_PORT", "3306"),
		DBUser:   getEnv("DB_USER", "root"),
		DBPass:   getEnv("DB_PASS", "12345678"),
		DBName:   getEnv("DB_NAME", "cetus"),
		JWTToken: []byte("64da497e5e8ae4b16de3d9a6782993b728115cd621606ed74ff995e92f9e7994"),
	}
}

func (c Config) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.DBUser, c.DBPass, c.DBHost, c.DBPort, c.DBName)
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
