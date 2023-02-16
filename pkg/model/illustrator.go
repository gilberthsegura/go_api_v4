package model

import (
	"apiv4/pkg/database"
	"strconv"

	"gorm.io/gorm"
)

type Illustrator struct {
	gorm.Model
	Name     string    `json:"name"`
	Pictures []Picture `json:"pictures" gorm:"-"`
}

// CREATE
func CreateIllustrator(illustrator *Illustrator) (*Illustrator, error) {
	err := database.DB.Create(&illustrator).Error
	if err != nil {
		return &Illustrator{}, err
	}
	return illustrator, nil
}

// READ
func ReadIllustrators() ([]Illustrator, error) {
	var illustrators []Illustrator
	err := database.DB.Find(&illustrators).Error
	if err != nil {
		return illustrators, err
	}
	return illustrators, nil
}

func ReadOneIllustrators(id string) (Illustrator, error) {
	var illustrator Illustrator
	idInt, _ := strconv.Atoi(id)
	err := database.DB.Where("id = ?", idInt).First(&illustrator).Error
	if err != nil {
		return illustrator, err
	}
	return illustrator, nil
}

// UPDATE
func UpdateIllustrator(illustrator *Illustrator, id string) (*Illustrator, error) {
	idInt, _ := strconv.Atoi(id)
	err := database.DB.Model(&Illustrator{}).
		Where("id = ?", idInt).
		Updates(illustrator).Error

	if err != nil {
		return &Illustrator{}, err
	}
	return illustrator, nil
}

// DELETE
func DeleteIllustrator(id string) error {
	idInt, _ := strconv.Atoi(id)
	err := database.DB.Delete(&Illustrator{}, idInt).Error
	if err != nil {
		return err
	}
	return nil
}
