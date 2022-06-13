package handlers

import (
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
	"go-pokemon-api/config"
	"net/http"
)

func Router(router *mux.Router, log *logrus.Logger, cfg *config.Environment, dbPool *pgxpool.Pool) {

	getPokemonByIDHandler := newGetPokemonByIDHandler(log)

	router.HandleFunc("/isalive", healthCheck).Name("isAlive")

	getPokemonByIDRouter := router.PathPrefix("/pokemon/{id}").Subrouter()
	getPokemonByIDRouter.Handle("", getPokemonByIDHandler).Methods(http.MethodGet).Name("pokemon/id")

}

func healthCheck(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}
