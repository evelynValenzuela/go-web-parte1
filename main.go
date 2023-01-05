package main

import (
	"parte2/controllers"

	"github.com/gin-gonic/gin"
)



func main(){
	//products := services.CargarDatos(FILEPATH)
	
	server := gin.Default()
	p := server.Group("/product")



	server.GET("/ping", controllers.GetPong)
	p.GET("", controllers.GetProducts)
	p.GET("/:id", controllers.GetProductById)
	p.GET("/search", controllers.GetProductsWithPrice)
	p.POST("", controllers.SaveProduct)
	
	err := server.Run(":8081")
	if err != nil {
		panic(err)
	}
}