package service

import "go-pokemon-api/persistence"

type PokemonService interface {
	GetPokemonById(id string) error
}

type PokemonServiceImpl struct {
	db *persistence.PostgresDB
}

func NewPokemonService(db *persistence.PostgresDB) PokemonService {
	return &PokemonServiceImpl{
		db: db,
	}
}

func (ps *PokemonServiceImpl) GetPokemonById(id string) error {

	//TODO implement
	return nil
}
