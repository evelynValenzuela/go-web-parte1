package handler

import (
	"errors"
	"fmt"
	"net/http"
	"parte2/internal/domain"
	"parte2/internal/product"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Handler struct {

}

var service = new(product.Service)

func (h *Handler) Init() {
	service.Init()
}


func (h *Handler) GetPong(ctx *gin.Context) {
	//Response
	ctx.String(http.StatusOK, "Pong")
}

func (h *Handler) GetProducts(ctx *gin.Context) {
	//Response
	ctx.JSON(http.StatusOK, service.GetAllProducts() )
}

func (h *Handler) GetProductById(ctx *gin.Context) {
	//Request
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		panic(err)
	}

	//Process
	productSearched, err := service.GetProductById(id)

	//Response
	if err != nil {
		ctx.JSON(http.StatusNotFound, domain.Err(err))
	} else {
		ctx.JSON(http.StatusOK, productSearched)
	}

}

func (h *Handler)  GetProductsByPrice(ctx *gin.Context) {
	//Request
	fmt.Println(ctx.Query("price"))
	priceGt , err:= strconv.ParseFloat(ctx.Query("price"), 8)
	if err != nil {
		panic(err)
	}

	//Process
	productsSearched, err := service.GetProductsByPrice(priceGt)

	//Response 
	if err != nil {
		ctx.JSON(http.StatusNotFound, domain.Err(err))
	} else {
		ctx.JSON(http.StatusOK, productsSearched)
	}

}

func (h *Handler) SaveProduct(ctx *gin.Context) {
	//Request 
	var req domain.Product
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, domain.Err(errors.New("Error en el servidor")))
		return
	}

	validate := validator.New()
	if err := validate.Struct(&req); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, domain.Err(product.ErrorInvalidData))
		return
	}

	//Process
	productSaved, err := service.SaveProduct(req)

	if err != nil {
		if errors.Is(err , product.ErrorCodeValueExist) {
			ctx.JSON(http.StatusConflict, domain.Err(err))
			return 
		} 
		ctx.JSON(http.StatusBadRequest, domain.Err(err))
		return
	}

	//Response 
	ctx.JSON(http.StatusCreated, productSaved)

}