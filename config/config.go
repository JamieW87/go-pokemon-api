package config

import "github.com/kelseyhightower/envconfig"

type Environment struct {
}

func Get() (*Environment, error) {
	var e Environment
	err := envconfig.Process("", &e)
	return &e, err
}
