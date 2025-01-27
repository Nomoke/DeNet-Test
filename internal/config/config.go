package config

import (
	"fmt"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

type Config struct {
	Database    Database
	Server      Server
	Logger      Logger
	Mail        Mail
	DecoderKeys DecoderKeys
}

type Database struct {
	PgUser     string `env:"POSTGRES_USER"`
	PgPassword string `env:"POSTGRES_PASSWORD"`
	PgHost     string `env:"POSTGRES_HOST"`
	PgPort     uint16 `env:"POSTGRES_PORT"`
	PgDatabase string `env:"POSTGRES_DB"`
	PgSSLMode  string `env:"POSTGRES_SSL_MODE"`
}

type Server struct {
	Port        string `env:"PORT" envDefault:":8080"`
	Host        string `env:"HOST"`
	LogFilePath string `env:"LOG_FILE_PATH" envDefault:"logs/logs.log"`
	Mode        string `env:"SERVER_MODE" envDefault:"debug"`
	TesterToken string `env:"TESTER_TOKEN"`
}

type Logger struct {
	LogFilePath string `env:"LOG_FILE_PATH"`
	Level       string `env:"LOG_LVL" envDefault:"debug"`
}

type Mail struct {
	From     string `env:"MAIL_FROM"`
	Password string `env:"MAIL_PASSWORD"`
	SmtpPort string `env:"MAIL_PORT"`
	SmtpHost string `env:"MAIL_HOST"`
}

type DecoderKeys struct {
	ServiceAuthSingingKey  string `env:"SINGING_KEY,notEmpty"`
	PasswordHashSalt       string `env:"PASSWORD_HASH_SALT,notEmpty"`
	SecretKeyEncryptUserID string `env:"SECRET_KEY_ENCRYPT_USER_ID,notEmpty"`
}

func New() (Config, error) {
	if err := godotenv.Load(); err != nil {
		return Config{}, fmt.Errorf("failed to load compose.env file: %w", err)
	}

	var cfg Config

	if err := env.Parse(&cfg); err != nil {
		return Config{}, fmt.Errorf("failed to parse config from environment variables: %w", err)
	}

	return cfg, nil
}
