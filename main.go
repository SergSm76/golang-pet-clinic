package main

import (
	"log"

	"my.app/pet-clinic/cache"
	"my.app/pet-clinic/config"
	"my.app/pet-clinic/handlers"
	"my.app/pet-clinic/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Загрузка конфигурации
	cfg := config.Load()

	// 2. Инициализация базы данных
	repository.InitDB()

	// 3. Инициализация менеджера кэша с параметрами из конфига
	cm := cache.NewManager(cfg.Cache.DefaultExpiration, cfg.Cache.CleanupInterval)
	handlers.InitVetHandler(cm)

	// 4. Настройка веб-сервера
	r := gin.Default()

	// 5. Определение маршрутов (REST API)
	r.GET("/owners", handlers.GetOwners)
	r.GET("/owners/:id", handlers.GetOwner)
	r.POST("/owners", handlers.CreateOwner)

	r.GET("/pets", handlers.GetPets)
	r.POST("/pets", handlers.CreatePet)

	r.GET("/vets", handlers.GetVets)
	r.POST("/vets", handlers.CreateVet)

	r.POST("/visits", handlers.CreateVisit)

	// 6. Запуск сервера на порту 8085
	if err := r.Run(":8085"); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
