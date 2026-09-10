package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

// CacheConfig хранит настройки кэша
type CacheConfig struct {
	DefaultExpiration time.Duration
	CleanupInterval   time.Duration
}

// Config — корневая структура конфигурации приложения
type Config struct {
	Cache CacheConfig
}

// Load читает конфигурацию из файла и переменных окружения.
// Если значения отсутствуют, используются дефолтные.
func Load() *Config {
	v := viper.New()

	// 1. Дефолтные значения (используются, если нет ни файла, ни переменных окружения)
	v.SetDefault("cache.default_expiration", "5m")
	v.SetDefault("cache.cleanup_interval", "10m")

	// 2. Чтение из файла конфигурации
	v.SetConfigName("config")   // имя файла без расширения
	v.SetConfigType("yaml")     // формат
	v.AddConfigPath("./config") // путь к папке с конфигом
	v.AddConfigPath(".")        // или в корне проекта

	// 3. Чтение из переменных окружения (префикс, разделитель)
	// Переменная окружения: PETCLINIC_CACHE_DEFAULT_EXPIRATION
	v.SetEnvPrefix("PETCLINIC")
	v.AutomaticEnv()

	// Пытаемся прочитать файл (если его нет — не критично, берём дефолты)
	if err := v.ReadInConfig(); err != nil {
		log.Println("Config file not found, using defaults and env vars:", err)
	}

	// 4. Парсинг значений
	cfg := &Config{}

	expStr := v.GetString("cache.default_expiration")
	cleanupStr := v.GetString("cache.cleanup_interval")

	var err error
	cfg.Cache.DefaultExpiration, err = time.ParseDuration(expStr)
	if err != nil {
		log.Printf("Invalid default_expiration '%s', using 5m: %v", expStr, err)
		cfg.Cache.DefaultExpiration = 5 * time.Minute
	}

	cfg.Cache.CleanupInterval, err = time.ParseDuration(cleanupStr)
	if err != nil {
		log.Printf("Invalid cleanup_interval '%s', using 10m: %v", cleanupStr, err)
		cfg.Cache.CleanupInterval = 10 * time.Minute
	}

	return cfg
}
