package ports

import "cloneslasher/internal/domain"

type ItemRepository interface {
	AddItem(item domain.Item)

	// GetByName gets namesakes items by specified name.
	GetByName(key domain.ItemName) ([]domain.ItemID, bool)

	// GetByPath gets item by its path.
	GetByID(key domain.ItemID) (domain.Item, bool)

	GetNames() []domain.ItemName

	GetIDs() []domain.ItemID
}
