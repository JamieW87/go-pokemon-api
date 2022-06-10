package main

import (
	"github.com/sirupsen/logrus"
	"go-pokemon-api/config"
)

func main() {

	env, err := config.Get()
	if err != nil {
		panic(err)
	}

	log := logrus.New()

}
