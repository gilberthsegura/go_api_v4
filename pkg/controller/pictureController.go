package controller

import (
	"apiv4/pkg/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddPicture(context *gin.Context) {
	//Parse the json with the object model
	var picture model.Picture
	if err := context.BindJSON(&picture); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//Create the object in the database
	newPicture, err := model.CreatePicture(&picture)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, newPicture)
}

func GetPictures(context *gin.Context) {
	//Get every object from database
	var pictures []model.Picture
	pictures, err := model.ReadPicture()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, pictures)
}

func GetOnePicture(context *gin.Context) {
	//Query the database for a single object
	id := context.Param("id")
	picture, err := model.ReadOnePicture(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, picture)
}

func GetPicturesByIllustrator(context *gin.Context) {
	//Query the database for a single object
	id := context.Param("id")
	pictures, err := model.ReadPictureByIllustrator(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, pictures)
}

func EditPicture(context *gin.Context) {
	//Parse the json with the object model
	id := context.Param("id")
	var picture model.Picture
	if err := context.BindJSON(&picture); err != nil {
		context.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	//Update a single row of the database
	updatedPicture, err := model.UpdatePicture(&picture, id)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, updatedPicture)
}

func RemovePicture(context *gin.Context) {
	id := context.Param("id")
	err := model.DeletePicture(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Picture deleted"})
}
