package repository

import (
	"api/models"
	"encoding/csv"
	"os"
	"strconv"
)

const dataFilePath = "./data/pokemon.csv"

func FetchPokemonData() (data []models.Pokemon, err error) {
	csvFile, err := os.Open(dataFilePath)
	if err != nil {
		return
	}
	defer csvFile.Close()
	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		return
	}
	for _, line := range csvLines {
		id, err := strconv.Atoi(line[0])
		if err != nil {
			break
		}
		item := models.Pokemon{
			Id:   id,
			Name: line[1],
		}
		data = append(data, item)
	}
	return
}
