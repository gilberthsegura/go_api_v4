package model

import (
	"apiv4/pkg/database"
	"strconv"

	"gorm.io/gorm"
)

type Picture struct {
	gorm.Model
	Title         string `json:"title"`
	IllustratorID uint   `json:"illustrator_id"`
}

// CREATE
func CreatePicture(picture *Picture) (*Picture, error) {
	err := database.DB.Create(&picture).Error
	if err != nil {
		return &Picture{}, err
	}
	return picture, nil
}

// READ
func ReadPicture() ([]Picture, error) {
	var pictures []Picture
	err := database.DB.Find(&pictures).Error
	if err != nil {
		return pictures, err
	}
	return pictures, nil
}

func ReadOnePicture(id string) (Picture, error) {
	var picture Picture
	idInt, _ := strconv.Atoi(id)
	err := database.DB.Where("id = ?", idInt).First(&picture).Error
	if err != nil {
		return picture, err
	}
	return picture, nil
}

func ReadPictureByIllustrator(id string) ([]Picture, error) {
	var pictures []Picture
	idInt, _ := strconv.Atoi(id)
	err := database.DB.Where("illustrator_id = ?", idInt).Find(&pictures).Error
	if err != nil {
		return pictures, err
	}
	return pictures, nil
}

// UPDATE
func UpdatePicture(picture *Picture, id string) (*Picture, error) {
	idInt, _ := strconv.Atoi(id)
	err := database.DB.Model(&Picture{}).
		Where("id = ?", idInt).
		Updates(picture).Error

	if err != nil {
		return &Picture{}, err
	}
	return picture, nil
}

// DELETE
func DeletePicture(id string) error {
	idInt, _ := strconv.Atoi(id)
	err := database.DB.Delete(&Picture{}, idInt).Error
	if err != nil {
		return err
	}
	return nil
}
