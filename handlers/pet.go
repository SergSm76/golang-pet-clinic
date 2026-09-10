package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"my.app/pet-clinic/models"
	"my.app/pet-clinic/repository"
)

func GetPets(c *gin.Context) {
	var pets []models.Pet
	repository.DB.Preload("PetType").Find(&pets)
	c.JSON(http.StatusOK, pets)
}

func CreatePet(c *gin.Context) {
	var pet models.Pet
	if err := c.ShouldBindJSON(&pet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	repository.DB.Create(&pet)
	c.JSON(http.StatusCreated, pet)
}
