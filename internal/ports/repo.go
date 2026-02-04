package ports

import "cloneslasher/internal/domain"

type ItemRepository interface {
	AddItem(item domain.Item) error

	// GetByName gets namesakes items by specified name.
	GetByName(key domain.ItemName) ([]domain.Item, bool)

	// GetByPath gets item by its path.
	GetByID(key domain.ItemID) (domain.Item, bool)

	GetNames() []domain.ItemName

	GetIDs() []domain.ItemID

	DumpMap() map[domain.ItemID]domain.Item

	DumpNamesakesMap() map[domain.ItemName][]domain.Item
}
