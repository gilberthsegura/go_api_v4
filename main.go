package main

import (
	"apiv4/pkg/controller"
	"apiv4/pkg/database"
	"apiv4/pkg/model"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	cors "github.com/rs/cors/wrapper/gin"
)

func main() {
	loadENV()
	loadDB()
	serveApp()
}

func loadENV() {
	err := godotenv.Load(".env.local")
	if err != nil {
		panic(err)
	}
}

func loadDB() {
	database.Connect()
	database.DB.AutoMigrate(&model.Illustrator{})
	database.DB.AutoMigrate(&model.Picture{})
}

func serveApp() {
	router := gin.Default()
	router.Use(cors.AllowAll())
	router.POST("/illustrator", controller.AddIllustrator)
	router.GET("/illustrator", controller.GetIllustrators)
	router.GET("/illustrator/:id", controller.GetOneIllustrator)
	router.POST("/illustrator/:id", controller.EditIllustrator)
	router.DELETE("/illustrator/:id", controller.RemoveIllustrator)

	router.POST("/picture", controller.AddPicture)
	router.GET("/picture", controller.GetPictures)
	router.GET("/picture/:id", controller.GetOnePicture)
	router.GET("/pictures/:id", controller.GetPicturesByIllustrator)
	router.POST("/picture/:id", controller.EditPicture)
	router.DELETE("/picture/:id", controller.RemovePicture)

	router.Run("localhost:8080")
}
