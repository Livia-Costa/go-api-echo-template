package config

import (
	"fmt"

	"github.com/victorsantosbrazil/go-api-echo-template/pkg/infra/datasource/postgresql"
)

type Config struct {
	Port       int
	DataSource *postgresql.Config
}

func (c *Config) GetAddress() string {
	return fmt.Sprintf(":%d", c.Port)
}
