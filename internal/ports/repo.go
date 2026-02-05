package ports

import "cloneslasher/internal/domain"

type ItemRepository interface {
	AddItem(ownerPath string, item domain.Item)

	// GetByName gets namesakes items by specified name.
	GetByName(key domain.ItemName) ([]domain.Item, bool)

	// GetByPath gets item by its path.
	GetByPath(path string) (domain.Item, bool)

	GetNames() []domain.ItemName

	GetIDs() []domain.ItemID

	// UpdateByID(item domain.Item) bool

	// DumpMap() map[domain.ItemID]domain.Item

	// DumpNamesakesMap() map[domain.ItemName][]domain.Item
}
