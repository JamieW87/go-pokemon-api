package handlers

import "net/http"

type getAllPokemonHandler struct {
}

func newGetAllPokemonHandler() *getAllPokemonHandler {
	return &getAllPokemonHandler{}

}

func (gap *getAllPokemonHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
