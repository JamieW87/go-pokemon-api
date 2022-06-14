package handlers

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"go-pokemon-api/config"
	"go-pokemon-api/persistence"
	"net/http"
)

func Router(router *mux.Router, log *logrus.Logger, cfg *config.Environment, db *persistence.PostgresDB) {

	getPokemonByIDHandler := newGetPokemonByIDHandler(log, db)

	router.HandleFunc("/isalive", healthCheck).Name("isAlive")

	getPokemonByIDRouter := router.PathPrefix("/pokemon/{id}").Subrouter()
	getPokemonByIDRouter.Handle("", getPokemonByIDHandler).Methods(http.MethodGet).Name("pokemon/id")

}

func healthCheck(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}
