package service

import (
	"shop/application/libs/utils"
	"shop/application/models"
)

func Create(param SCreate) (models.Model, error) {
	model := models.Model{}

	return model, nil
}
func Detail(param SDetail) (models.Model, error) {
	model := models.Model{}

	return model, nil
}
func Count(param SAll) (int64, error) {
	var count int64

	return count, nil
}
func Pages(param SAll, page utils.Pages) ([]models.Model, error) {
	var list []models.Model

	return list, nil
}
func Update(param SUpdate) (models.Model, error) {
	model := models.Model{}

	return model, nil
}
func Del(param SUpdate) (models.Model, error) {
	model := models.Model{}

	return model, nil
}
