package config

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/caarlos0/env/v8"
	"github.com/joho/godotenv"
)

type Config struct {
	AppServerPort string   `env:"PORT,required"`
	Database      Database `env:"DATABASE"`
}

type Database struct {
	Host                string `env:"DATABASE_HOST,required"`
	User                string `env:"DATABASE_USER,required"`
	Password            string `env:"DATABASE_PASSWORD,required"`
	SSLMode             string `env:"DATABASE_SSL_MODE,required"`
	Name                string `env:"DATABASE_NAME,required"`
	Port                int    `env:"DATABASE_PORT,required"`
	ConnectTimeout      int    `env:"DATABASE_CONNECT_TIMEOUT" required:"true"`
	MaxPoolConns        int    `env:"DATABASE_MAX_POOL_CONNS" required:"true"`
	DefaultQueryTimeout int    `env:"DEFAULT_DATABASE_QUERY_TIMEOUT" required:"true"`
}

func init() {
	currentDir := extractCurrentDir()
	err := godotenv.Load(currentDir + "/.env")
	if err != nil {
		log.Fatalf("unable to load .env file: %e", err)
	}
}

func extractCurrentDir() string {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	if _, err := os.Stat(currentDir + "/.env"); err != nil {
		tokens := strings.Split(currentDir, "/")
		var buf bytes.Buffer
		for i := 0; i < len(tokens)-1; i++ {
			buf.WriteString(tokens[i])
			if i != len(tokens)-2 {
				buf.WriteString("/")
			}
		}
		return buf.String()
	}
	return currentDir
}

func (cfg Config) GetDBString() string {
	dbString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s", cfg.Database.Host, cfg.Database.User,
		cfg.Database.Password, cfg.Database.Name, cfg.Database.Port, cfg.Database.SSLMode)
	return dbString
}

func Load() Config {
	log.Default().SetPrefix("\r")

	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("unable to parse .env file: %e", err)
	}
	return cfg
}
