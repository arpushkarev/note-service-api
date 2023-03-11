package config

import (
	"encoding/json"
	"net"
	"os"
	"strings"

	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	dbPassEscSeq = "{password}"
	password     = "note-service-password"
)

// DB structure
type DB struct {
	DSN                string `json:"dsn"`
	MaxOpenConnections int32  `json:"max_open_connections"`
}

// GRPC structure
type GRPC struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

// HTTP structure
type HTTP struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

// IConfig ...
type IConfig interface {
	GetDBConfig() (*pgxpool.Config, error)
	GetGRPCAddress() string
	GetHTTPAddress() string
}

// Config structure
type Config struct {
	DB   DB   `json:"db"`
	GRPC GRPC `json:"grpc"`
	HTTP HTTP `json:"http"`
}

// NewConfig starts config
func NewConfig(path string) (*Config, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	config := &Config{}
	err = json.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

// GetDBConfig starts pgxpool config
func (c *Config) GetDBConfig() (*pgxpool.Config, error) {
	dbDSN := strings.ReplaceAll(c.DB.DSN, dbPassEscSeq, password)

	poolConfig, err := pgxpool.ParseConfig(dbDSN)
	if err != nil {
		return nil, err
	}

	poolConfig.ConnConfig.BuildStatementCache = nil
	poolConfig.ConnConfig.PreferSimpleProtocol = true
	poolConfig.MaxConns = c.DB.MaxOpenConnections

	return poolConfig, err
}

// GetGRPCAddress GRPC generates address from config
func (c *Config) GetGRPCAddress() string {
	return net.JoinHostPort(c.GRPC.Host, c.GRPC.Port)
}

// GetHTTPAddress HTTP generates address from config
func (c *Config) GetHTTPAddress() string {
	return net.JoinHostPort(c.HTTP.Host, c.HTTP.Port)
}
