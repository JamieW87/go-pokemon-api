package handlers

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

type getPokemonByID struct {
	log *logrus.Logger
}

func newGetPokemonByIDHandler(log *logrus.Logger) *getPokemonByID {
	return &getPokemonByID{
		log: log,
	}

}

func (gp *getPokemonByID) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
