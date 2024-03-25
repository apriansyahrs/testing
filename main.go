package main

import (
	"gin-railway/database"
	"gin-railway/handler"
	"gin-railway/repository"
	"gin-railway/service"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	postgres := database.NewPostgres()
	if postgres.Err != nil {
		panic(postgres.Err)
	}

	productRepo := repository.NewProductRepo(postgres)
	productService := service.NewProductService(productRepo)
	productHandler := handler.NewProductHandler(productService)

	r := gin.Default()

	r.GET("/product/:id", productHandler.FindOneProduct)

	r.Run(":" + os.Getenv("PORT"))
}
