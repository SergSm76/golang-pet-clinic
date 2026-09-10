package repository

import (
	_ "embed"
	"log"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

//go:embed db/sqlite/schema.sql
var schemaSQL string

//go:embed db/sqlite/data.sql
var dataSQL string

var DB *gorm.DB

func InitDB() {
	var err error

	// Подключение к PostgreSQL (каноничный вариант для Petclinic)
	// DSN: host=localhost user=petclinic password=petclinic dbname=petclinic port=5432 sslmode=disable
	// DB, err = gorm.Open(postgres.Open("host=localhost user=postgres password=postgres dbname=petclinic port=5432 sslmode=disable"), &gorm.Config{})

	// Подключение к  SQLite:
	DB, err = gorm.Open(sqlite.Open("petclinic.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Выполнение скрипта создания схемы
	// Ошибки игнорируются, если таблицы уже существуют (благодаря IF NOT EXISTS в schema.sql)
	if err := DB.Exec(schemaSQL).Error; err != nil {
		log.Println("Schema execution note:", err)
	}

	// Выполнение скрипта заполнения данными
	// В реальном проекте для этого лучше использовать библиотеки миграций, например, golang-migrate
	if err := DB.Exec(dataSQL).Error; err != nil {
		log.Println("Data execution note (may already exist):", err)
	}

	log.Println("Database initialized successfully from embedded SQL files.")
}
