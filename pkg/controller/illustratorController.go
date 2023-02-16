package controller

import (
	"apiv4/pkg/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddIllustrator(context *gin.Context) {
	//Parse the json with the object model
	var illustrator model.Illustrator
	if err := context.BindJSON(&illustrator); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//Create the object in the database
	newIllustrator, err := model.CreateIllustrator(&illustrator)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, newIllustrator)
}

func GetIllustrators(context *gin.Context) {
	//Get every illustrator from database
	var illustrators []model.Illustrator
	illustrators, err := model.ReadIllustrators()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, illustrators)
}

func GetOneIllustrator(context *gin.Context) {
	//Query the database for a single illustrator
	id := context.Param("id")
	illustrator, err := model.ReadOneIllustrators(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, illustrator)
}

func EditIllustrator(context *gin.Context) {
	//Parse the json with the object model
	id := context.Param("id")
	var illustrator model.Illustrator
	if err := context.BindJSON(&illustrator); err != nil {
		context.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	//Update a single row of the database
	updatedIllustrator, err := model.UpdateIllustrator(&illustrator, id)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, updatedIllustrator)
}

func RemoveIllustrator(context *gin.Context) {
	id := context.Param("id")
	err := model.DeleteIllustrator(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Illustrator deleted"})
}
