package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"my.app/pet-clinic/models"
)

func TestCreateOwner(t *testing.T) {
	db := SetupTestDB(t)

	owner := models.Owner{
		FirstName: "George",
		LastName:  "Franklin",
		Address:   "110 W. Liberty St.",
		City:      "Madison",
		Telephone: "6085551023",
	}

	result := db.Create(&owner)
	require.NoError(t, result.Error)
	assert.NotZero(t, owner.ID)
}

func TestFindOwnerByID(t *testing.T) {
	db := SetupTestDB(t)

	// Создаём владельца
	owner := models.Owner{
		FirstName: "Betty",
		LastName:  "Davis",
		Address:   "638 Cardinal Ave.",
		City:      "Sun Prairie",
		Telephone: "6085551749",
	}
	db.Create(&owner)

	// Ищем по ID
	var found models.Owner
	result := db.First(&found, owner.ID)
	require.NoError(t, result.Error)
	assert.Equal(t, "Betty", found.FirstName)
	assert.Equal(t, "Davis", found.LastName)
}

func TestFindAllOwners(t *testing.T) {
	db := SetupTestDB(t)

	// Создаём нескольких владельцев
	owners := []models.Owner{
		{FirstName: "George", LastName: "Franklin"},
		{FirstName: "Betty", LastName: "Davis"},
		{FirstName: "Eduardo", LastName: "Rodriquez"},
	}
	for _, o := range owners {
		db.Create(&o)
	}

	// Ищем всех
	var found []models.Owner
	result := db.Find(&found)
	require.NoError(t, result.Error)
	assert.Len(t, found, 3)
}

func TestUpdateOwner(t *testing.T) {
	db := SetupTestDB(t)

	owner := models.Owner{
		FirstName: "Harold",
		LastName:  "Davis",
		Address:   "563 Friendly St.",
		City:      "Windsor",
		Telephone: "6085553198",
	}
	db.Create(&owner)

	// Обновляем
	owner.City = "New Windsor"
	db.Save(&owner)

	// Проверяем
	var updated models.Owner
	db.First(&updated, owner.ID)
	assert.Equal(t, "New Windsor", updated.City)
}

func TestDeleteOwner(t *testing.T) {
	db := SetupTestDB(t)

	owner := models.Owner{FirstName: "Test", LastName: "Delete"}
	db.Create(&owner)

	db.Delete(&owner)

	var found models.Owner
	result := db.First(&found, owner.ID)
	assert.Error(t, result.Error) // запись не найдена
}
