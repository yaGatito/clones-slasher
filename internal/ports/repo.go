package ports

import "cloneslasher/internal/domain"

type ItemRepository interface {
	AddItem(item domain.Item)
	GetByName(key string) ([]domain.Item, error)
	GetByPath(key string) ([]domain.Item, error)
	GetNames() []string
	GetPaths() []string
}
