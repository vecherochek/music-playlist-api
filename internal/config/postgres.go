package config

import (
	"fmt"
	"os"
)

const (
	postgresEnvHost     = "POSTGRES_HOST"
	postgresEnvPort     = "POSTGRES_PORT"
	postgresEnvUser     = "POSTGRES_USER"
	postgresEnvPassword = "POSTGRES_PASSWORD"
	postgresEnvDbName   = "POSTGRES_DBNAME"
	postgresEnvSslMode  = "POSTGRES_SSLMODE"
)

type PostgresConfig interface {
	Url() string
}

type postgresConfig struct {
	host     string
	port     string
	user     string
	password string
	dbname   string
	sslmode  string
}

func NewPostgresConfig() (PostgresConfig, error) {
	return &postgresConfig{
		host:     os.Getenv(postgresEnvHost),
		port:     os.Getenv(postgresEnvPort),
		user:     os.Getenv(postgresEnvUser),
		password: os.Getenv(postgresEnvPassword),
		dbname:   os.Getenv(postgresEnvDbName),
		sslmode:  os.Getenv(postgresEnvSslMode),
	}, nil
}

func (cfg *postgresConfig) Url() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.host, cfg.port, cfg.user, cfg.password, cfg.dbname, cfg.sslmode,
	)
}
