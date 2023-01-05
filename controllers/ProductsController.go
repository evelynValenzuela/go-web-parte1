package controllers

import (
	"fmt"
	"net/http"
	"parte2/models"
	"parte2/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

var products = services.CargarDatos("products.json")

func GetPong(ctx *gin.Context) {
	//Response
	ctx.String(http.StatusOK, "Pong")
}

func GetProducts(ctx *gin.Context) {
	//Response
	ctx.JSON(http.StatusOK, products)
}

func GetProductById(ctx *gin.Context) {
	//Request
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		panic(err)
	}

	//Process
	var productSearched models.Product
	for _, product := range products {
		if(product.Id == id) {
			productSearched = product
			break
		}
	}

	//Response
	if productSearched.Id == 0 {
		ctx.String(http.StatusNotFound, "El producto ingresado no se encuentra en la base de datos")
	} else {
		ctx.JSON(http.StatusOK, productSearched)
	}

}

func GetProductsWithPrice(ctx *gin.Context) {
	//Request
	fmt.Println(ctx.Query("price"))
	priceGt , err:= strconv.ParseFloat(ctx.Query("price"), 8)
	if err != nil {
		panic(err)
	}

	//Process
	var productsSearched []models.Product
	for _, product := range products {
		if(product.Price > priceGt) {
			productsSearched = append(productsSearched, product)
		}
	}

	//Response 
	if len(productsSearched) == 0{
		ctx.String(http.StatusNotFound, "No se encontraron productos con el criterio establecido")
	} else {
		ctx.JSON(http.StatusOK, productsSearched)
	}

}