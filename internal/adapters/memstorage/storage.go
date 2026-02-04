package storage

import (
	"cloneslasher/internal/domain"
	"cloneslasher/internal/ports"
	"sync"
)

type Item struct {
	Path      string
	Name      string
	Size      int64
	Extension string
	IsFolder  bool
}

// ItemStorage name oriented memory storage.
type ItemStorage struct {
	nameLoke          sync.RWMutex
	nameOrientedStore map[string][]domain.Item
	pathLoke          sync.RWMutex
	pathOrientedStore map[string]Item
}

var _ ports.ItemRepository = (*ItemStorage)(nil)

func NewItemStorage() *ItemStorage {
	return &ItemStorage{
		nameLoke:          sync.RWMutex{},
		nameOrientedStore: make(map[string][]domain.Item),
		pathLoke:          sync.RWMutex{},
		pathOrientedStore: make(map[string]Item),
	}
}

func (is *ItemStorage) AddItem(item domain.Item) {
	is.pathLoke.RLock()
	is.pathOrientedStore[item.Path] = mapDomainToItem(item)
	is.pathLoke.RUnlock()

	is.nameLoke.RLock()
	items, exists := is.nameOrientedStore[item.Name]
	is.nameLoke.RUnlock()

	is.nameLoke.Lock()
	if exists {
		is.nameOrientedStore[item.Name] = append(items, item)
	} else {
		items = make([]domain.Item, 1)
		items[0] = item
		is.nameOrientedStore[item.Name] = items
	}
	is.nameLoke.Unlock()
}

func (is *ItemStorage) GetByName(itemName string) ([]domain.Item, bool) {
	is.nameLoke.RLock()
	defer is.nameLoke.RUnlock()

	items, exists := is.nameOrientedStore[itemName]
	return items, exists
}

func (is *ItemStorage) GetByPath(itemPath string) (domain.Item, bool) {
	is.pathLoke.RLock()
	defer is.pathLoke.RUnlock()

	item, exists := is.pathOrientedStore[itemPath]
	return mapItemToDomain(item), exists
}

func (is *ItemStorage) GetNames() []string {
	is.nameLoke.RLock()
	defer is.nameLoke.RUnlock()

	var names []string
	for name := range is.nameOrientedStore {
		names = append(names, name)
	}
	return names
}

func (is *ItemStorage) GetPaths() []string {
	is.nameLoke.RLock()
	defer is.nameLoke.RUnlock()

	var names []string
	for name := range is.pathOrientedStore {
		names = append(names, name)
	}
	return names
}

func mapDomainToItem(item domain.Item) Item {
	return Item{
		Path:      item.Path,
		Name:      item.Name,
		Size:      item.Size,
		Extension: item.Extension,
		IsFolder:  item.IsFolder,
	}
}

func mapItemToDomain(item Item) domain.Item {
	return domain.Item{
		Path:      item.Path,
		Name:      item.Name,
		Size:      item.Size,
		Extension: item.Extension,
		IsFolder:  item.IsFolder,
	}
}
