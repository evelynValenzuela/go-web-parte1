package services

import (
	"encoding/json"
	"io"
	"os"
	"parte2/models"
	"strings"
)

func CargarDatos(filepath string) []models.Product {
	file, err := os.ReadFile(filepath)

	if(err != nil ) {
		panic(err)
	}
	streaming := strings.NewReader(string(file))
	decoder := json.NewDecoder(streaming)
	for {
		var product []models.Product

		if err := decoder.Decode(&product); err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		} 
		return product
	}
	return nil


}