package handlers

import (
	"github.com/sirupsen/logrus"
	"go-pokemon-api/persistence"
	"go-pokemon-api/service"
	"net/http"
)

type getPokemonByID struct {
	log     *logrus.Logger
	service service.PokemonService
}

func newGetPokemonByIDHandler(log *logrus.Logger, db *persistence.PostgresDB) *getPokemonByID {
	return &getPokemonByID{
		log:     log,
		service: service.NewPokemonService(db),
	}

}

func (gp *getPokemonByID) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	//TODO
	//Get ID from url
	//Send to service
	//Return OK and return correct pokemon

}
