package ports

import "cloneslasher/internal/domain"

type ItemRepository interface {
	AddItem(item domain.Item)

	// GetByName gets namesakes items by specified name.
	GetByName(key string) ([]domain.Item, bool)

	// GetByPath gets item by its path.
	GetByPath(key string) (domain.Item, bool)

	GetNames() []string

	GetPaths() []string
}
