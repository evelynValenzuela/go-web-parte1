package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"parte2/models"
	"parte2/repository"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)


func GetPong(ctx *gin.Context) {
	//Response
	ctx.String(http.StatusOK, "Pong")
}

func GetProducts(ctx *gin.Context) {
	//Response
	ctx.JSON(http.StatusOK, repository.CargarDatos())
}

func GetProductById(ctx *gin.Context) {
	//Request
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		panic(err)
	}

	//Process
	productSearched := repository.GetProduct(id)

	//Response
	if productSearched.Id == 0 {
		ctx.JSON(http.StatusNotFound, models.Err(repository.ErrorProductNotFound))
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
	productsSearched := repository.GetProductsWithFilter(priceGt)

	//Response 
	if len(productsSearched) == 0{
		ctx.JSON(http.StatusNotFound, models.Err(repository.ErrorProductsNotFound))
	} else {
		ctx.JSON(http.StatusOK, productsSearched)
	}

}

func SaveProduct(ctx *gin.Context) {
	//Request 
	var req models.Product
	if err := ctx.ShouldBind(&req); err != nil {
		
		ctx.JSON(http.StatusBadRequest, models.Err(errors.New("Error en el servidor")))
		return
	}

	validate := validator.New()
	if err := validate.Struct(&req); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, models.Err(repository.ErrorInvalidData))
		return
	}

	//Process
	product, err := repository.AddProduct(req)

	if err != nil {
		if errors.Is(err , repository.ErrorCodeValueExist) {
			ctx.JSON(http.StatusConflict, models.Err(err))
			return 
		} 
		ctx.JSON(http.StatusBadRequest, models.Err(err))
		return
	}

	//Response 
	ctx.JSON(http.StatusCreated, product)

}