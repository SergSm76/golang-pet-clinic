package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"my.app/pet-clinic/models"
	"my.app/pet-clinic/repository"
)

func CreateVisit(c *gin.Context) {
	var visit models.Visit
	if err := c.ShouldBindJSON(&visit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	repository.DB.Create(&visit)
	c.JSON(http.StatusCreated, visit)
}
