package handlers

import (
	"net/http"

	"my.app/pet-clinic/cache"
	"my.app/pet-clinic/models"
	"my.app/pet-clinic/repository"

	"github.com/gin-gonic/gin"
)

// CacheManager Глобальный экземпляр менеджера кэша (можно пробрасывать через DI в реальном проекте)
var CacheManager *cache.Manager

func InitVetHandler(cm *cache.Manager) {
	CacheManager = cm
}

func GetVets(c *gin.Context) {
	const cacheKey = "vets"

	// 1. Пытаемся получить данные из кэша (аналог @Cacheable("vets") в Spring)
	if cachedVets, found := CacheManager.Get(cacheKey); found {
		c.JSON(http.StatusOK, cachedVets)
		return
	}

	// 2. Если в кэше нет, идем в БД
	var vets []models.Vet
	repository.DB.Preload("Specialties").Find(&vets)

	// 3. Сохраняем результат в кэш
	CacheManager.Set(cacheKey, vets)

	c.JSON(http.StatusOK, vets)
}

// CreateVet с инвалидацией кэша (аналог @CacheEvict)
func CreateVet(c *gin.Context) {
	var vet models.Vet
	if err := c.ShouldBindJSON(&vet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	repository.DB.Create(&vet)

	// 4. Инвалидация кэша при изменении данных
	CacheManager.Delete("vets")

	c.JSON(http.StatusCreated, vet)
}
