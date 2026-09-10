package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"my.app/pet-clinic/models"
	"my.app/pet-clinic/repository"
)

func GetOwners(c *gin.Context) {
	var owners []models.Owner
	repository.DB.Find(&owners)
	c.JSON(http.StatusOK, owners)
}

func GetOwner(c *gin.Context) {
	var owner models.Owner
	if err := repository.DB.Preload("Pets").First(&owner, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Owner not found"})
		return
	}
	c.JSON(http.StatusOK, owner)
}

func CreateOwner(c *gin.Context) {
	var owner models.Owner
	if err := c.ShouldBindJSON(&owner); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	repository.DB.Create(&owner)
	c.JSON(http.StatusCreated, owner)
}
