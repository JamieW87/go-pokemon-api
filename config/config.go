package config

import "github.com/kelseyhightower/envconfig"

type Environment struct {
	// Port is the port to the run server on.
	Port             int    `required:"true"`
	PostgresHost     string `required:"true"`
	PostgresPort     int    `required:"true"`
	PostgresName     string `required:"true"`
	PostgresUser     string `required:"true"`
	PostgresPassword string `required:"true"`
}

func Get() (*Environment, error) {
	var e Environment
	err := envconfig.Process("", &e)
	return &e, err
}
