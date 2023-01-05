package main

import (
	"parte2/controllers"
	"github.com/gin-gonic/gin"
)



func main(){
	//products := services.CargarDatos(FILEPATH)
	server := gin.Default()
	server.GET("/ping", controllers.GetPong)
	server.GET("/products", controllers.GetProducts)
	server.GET("/products/:id", controllers.GetProductById)
	server.GET("/products/search", controllers.GetProductsWithPrice)
	
	err := server.Run()
	if err != nil {
		panic(err)
	}
}