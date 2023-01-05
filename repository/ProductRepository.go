package repository

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"parte2/models"
	"parte2/services"
	"strings"
)

var (
	ErrorInvalidDate = errors.New("Error: El formato de fecha es inválido")
	ErrorProductNotFound = errors.New("Error: El producto NO ya existe")
	ErrorProductsNotFound = errors.New("No se encontraron productos con el criterio establecido")
	ErrorCodeValueExist = errors.New("Error: El code value del producto ya existe")
	ErrorInvalidData= errors.New("Error: No se permiten valores vacíos")

)

var products   = CargarDatos()
var ItemActual = 500

func CargarDatos() []models.Product {
	file, err := os.ReadFile("products.json")

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

func GetProduct(id int) (productSearched models.Product) {

	for _, product := range products {
		if(product.Id == id) {
			productSearched = product
			break
		}
	}
	return
}

func GetProductsWithFilter(priceGt float64) (productsSearched []models.Product) {
	for _, product := range products {
		if(product.Price > priceGt) {
			productsSearched = append(productsSearched, product)
		}
	}
	return 
}

func AddProduct(product models.Product) (productStoraged models.Product, err error) {
	if services.ValidateCodeValue(products, product.CodeValue) {
		err = ErrorCodeValueExist
		return 
	}

	if services.ValidateDate(products, product.Expiration) {
		err = ErrorInvalidDate
		return
	}

	ItemActual++
	product.Id = ItemActual
	productStoraged = product

	products = append(products, productStoraged)
	
	return 

}