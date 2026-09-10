package repository

import (
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"my.app/pet-clinic/models"
)

// SetupTestDB создаёт in-memory SQLite базу для тестов (аналог @DataJpaTest с H2)
func SetupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		t.Fatalf("Failed to connect to test database: %v", err)
	}

	// Миграция схем
	err = db.AutoMigrate(
		&models.PetType{},
		&models.Specialty{},
		&models.Vet{},
		&models.Owner{},
		&models.Pet{},
		&models.Visit{},
	)
	if err != nil {
		t.Fatalf("Failed to migrate test database: %v", err)
	}

	return db
}
