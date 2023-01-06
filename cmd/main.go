package main

import (
	"parte2/cmd/handler"
	"github.com/gin-gonic/gin"
)

func main(){
	
	server := gin.Default()
	p := server.Group("/products")

	h := new(handler.Handler)
	h.Init()

	server.GET("/ping", h.GetPong)
	p.GET("", h.GetProducts)
	p.GET("/:id", h.GetProductById)
	p.GET("/search", h.GetProductsByPrice)
	p.POST("", h.SaveProduct)
	
	err := server.Run(":8081")
	if err != nil {
		panic(err)
	}
}