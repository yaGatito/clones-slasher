package storage

import (
	"cloneslasher/internal/domain"
	"cloneslasher/internal/ports"
	"sync"
)

// ItemStorage name oriented memory storage.
type ItemStorage struct {
	nameLoke          sync.RWMutex
	nameOrientedStore map[domain.ItemName][]domain.ItemID

	pathLoke          sync.RWMutex
	pathOrientedStore map[domain.ItemID]domain.Item
}

var _ ports.ItemRepository = (*ItemStorage)(nil)

func NewItemStorage() *ItemStorage {
	return &ItemStorage{
		nameLoke:          sync.RWMutex{},
		nameOrientedStore: make(map[domain.ItemName][]domain.ItemID),

		pathLoke:          sync.RWMutex{},
		pathOrientedStore: make(map[domain.ItemID]domain.Item),
	}
}

func (s *ItemStorage) AddItem(item domain.Item) {
	s.pathLoke.RLock()
	_, ok := s.pathOrientedStore[item.ID]
	if ok {
		panic("adding existed item")
	} else {
		s.pathOrientedStore[item.ID] = item
	}
	s.pathLoke.RUnlock()

	s.nameLoke.Lock()
	namesakesIDs, ok := s.nameOrientedStore[item.Name]
	if ok {
		s.nameOrientedStore[item.Name] = append(namesakesIDs, item.ID)
	} else {
		namesakesIDs = make([]domain.ItemID, 0)
		namesakesIDs = append(namesakesIDs, item.ID)
		s.nameOrientedStore[item.Name] = namesakesIDs
	}
	s.nameLoke.Unlock()
}

func (s *ItemStorage) GetByName(itemName domain.ItemName) ([]domain.ItemID, bool) {
	s.nameLoke.RLock()
	defer s.nameLoke.RUnlock()

	items, exists := s.nameOrientedStore[itemName]
	return items, exists
}

func (s *ItemStorage) GetByID(itemPath domain.ItemID) (domain.Item, bool) {
	s.pathLoke.RLock()
	defer s.pathLoke.RUnlock()

	item, exists := s.pathOrientedStore[itemPath]
	return item, exists
}

func (s *ItemStorage) GetNames() []domain.ItemName {
	s.nameLoke.RLock()
	defer s.nameLoke.RUnlock()

	var names []domain.ItemName
	for name := range s.nameOrientedStore {
		names = append(names, name)
	}
	return names
}

func (s *ItemStorage) GetIDs() []domain.ItemID {
	s.nameLoke.RLock()
	defer s.nameLoke.RUnlock()

	var IDs []domain.ItemID
	for ID := range s.pathOrientedStore {
		IDs = append(IDs, ID)
	}
	return IDs
}
