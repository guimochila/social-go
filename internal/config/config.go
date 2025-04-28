// Copyright (c) 2025. guimochila.com. Continuous Learning.

package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port string `default:"8080" envconfig:"PORT"`
}

func (c *Config) GetAddress() string {
	return fmt.Sprintf(":%s", c.Port)
}

func FromEnv(config *Config) error {
	return envconfig.Process("config", config)
}
