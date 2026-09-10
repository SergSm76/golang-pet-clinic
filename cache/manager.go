package cache

import (
	"time"

	"github.com/patrickmn/go-cache"
)

// Manager — обертка над in-memory кэшем
type Manager struct {
	c *cache.Cache
}

// NewManager создает новый менеджер кэша с параметрами из конфигурации
func NewManager(defaultExpiration, cleanupInterval time.Duration) *Manager {
	return &Manager{
		c: cache.New(defaultExpiration, cleanupInterval),
	}
}

// Get получает значение из кэша. Если его нет, возвращает nil и false.
func (m *Manager) Get(key string) (any, bool) {
	return m.c.Get(key)
}

// Set сохраняет значение в кэш с временем жизни по умолчанию.
func (m *Manager) Set(key string, value any) {
	m.c.SetDefault(key, value)
}

// Delete удаляет значение из кэша
func (m *Manager) Delete(key string) {
	m.c.Delete(key)
}
