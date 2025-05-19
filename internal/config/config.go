// Copyright (c) 2025. guimochila.com. Continuous Learning.

package config

import (
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type DBConfig struct {
	Host         string `default:"localhost" envconfig:"DB_HOST"`
	Username     string `default:""          envconfig:"DB_USERNAME"`
	Password     string `default:""          envconfig:"DB_PASSWORD"`
	DatabaseName string `default:""          envconfig:"DB_NAME"`
	MaxOpenConns int    `default:"30"        envconfig:"DB_MAX_OPEN_CONNS"`
	MaxIdleConns int    `default:"30"        envconfig:"DB_MAX_IDLE_CONNS"`
	MaxIdleTime  int    `default:"30"        envconfig:"DB_MAX_IDLE_TIME"`
}

func (d *DBConfig) DSN() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s/%s?sslmode=disable",
		d.Username,
		d.Password,
		d.Host,
		d.DatabaseName,
	)
}

func (d *DBConfig) GetMaxIdleTime() time.Duration {
	return time.Duration(d.MaxIdleTime)
}

type Config struct {
	Port string `default:"8080" envconfig:"PORT"`
	DB   DBConfig
}

func (c *Config) GetAddress() string {
	return fmt.Sprintf(":%s", c.Port)
}

func FromEnv(config *Config) error {
	return envconfig.Process("config", config)
}
