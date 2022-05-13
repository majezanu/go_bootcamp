package services

import (
	"api/dataclass"
	"api/repository"
	"fmt"
	"net/http"
	"strconv"
)

func RetrievePokemon(quantityInput string) (output []dataclass.PokemonDataclass, errorOutput dataclass.ErrorDataclass) {
	quantity, err := strconv.Atoi(quantityInput)

	if err != nil {
		quantity = -1
	}

	output = make([]dataclass.PokemonDataclass, 0)
	input, err := repository.FetchPokemonData()

	if err != nil {
		errorOutput = dataclass.ErrorDataclass{
			Code:        http.StatusNotFound,
			Description: fmt.Sprint(err),
		}
		return
	}
	inputSize := len(input)
	if inputSize < quantity {
		errorOutput = dataclass.ErrorDataclass{
			Code:        http.StatusUnprocessableEntity,
			Description: "Max quantity",
		}
		return
	}

	for index, item := range input {
		data := dataclass.PokemonDataclass{
			Id:   item.Id,
			Name: item.Name,
		}
		if quantity == index {
			break
		}
		output = append(output, data)
	}
	return
}
