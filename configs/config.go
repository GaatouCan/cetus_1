package configs

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"sync"
)

type Config struct {
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
	Security SecurityConfig `yaml:"security"`
}

type ServerConfig struct {
	Port int `yaml:"port"`
}

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}

type SecurityConfig struct {
	JWTToken string `yaml:"jwt_token"`
}

var (
	instance *Config
	once     sync.Once
)

func GetConfig() *Config {
	once.Do(func() {
		instance = loadConfig("configs/config.yaml")
	})
	return instance
}

func loadConfig(filename string) *Config {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to open config file: %s", err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Printf("Failed to close config file: %s", err)
		}
	}()

	config := &Config{}
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(config); err != nil {
		log.Fatalf("Failed to decode config file: %s", err)
	}

	log.Println("Loaded config")
	return config
}

func (c Config) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.Database.Username, c.Database.Password, c.Database.Host, c.Database.Port, c.Database.Name)
}
