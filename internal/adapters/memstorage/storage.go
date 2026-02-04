package storage

import (
	"cloneslasher/internal/domain"
	"cloneslasher/internal/ports"
	"fmt"
	"slices"
)

// ItemStorage name oriented memory storage.
type ItemStorage struct {
	// clonesLoke     sync.RWMutex
	namesakesRelStore map[domain.ItemName][]domain.Item

	// loke            sync.RWMutex
	idOrientedStore map[domain.ItemID]domain.Item
}

var _ ports.ItemRepository = (*ItemStorage)(nil)

func NewItemStorage() *ItemStorage {
	return &ItemStorage{
		// clonesLoke:     sync.RWMutex{},
		namesakesRelStore: make(map[domain.ItemName][]domain.Item),

		// loke:            sync.RWMutex{},
		idOrientedStore: make(map[domain.ItemID]domain.Item),
	}
}

func (s *ItemStorage) AddItem(item domain.Item) error {
	err := s.addItem(item)
	if err != nil {
		return err
	}
	err = s.addItemToNamesakes(item)
	if err != nil {
		return err
	}
	return nil
}

func (s *ItemStorage) addItem(item domain.Item) error {
	// s.loke.RLock()
	existedItem, ok := s.idOrientedStore[item.ID]
	if ok {
		return fmt.Errorf("adding existed item %v", existedItem)
	} else {
		s.idOrientedStore[item.ID] = item
	}
	return nil
	// s.loke.RUnlock()
}

func (s *ItemStorage) addItemToNamesakes(item domain.Item) error {
	// s.clonesLoke.Lock()
	namesakes, ok := s.namesakesRelStore[item.Name]
	if ok {
		ix := slices.IndexFunc(namesakes, item.Equal)
		// fmt.Printf("ix: %d; item: %s; namesakes: %v\n", ix, item.ID, namesakes)
		if ix > 0 {
			return fmt.Errorf("failed to add existed namesake %v", namesakes[ix])
		} else {
			s.namesakesRelStore[item.Name] = append(namesakes, item)
		}
	} else {
		namesakes = make([]domain.Item, 0)
		s.namesakesRelStore[item.Name] = namesakes
	}
	return nil
	// s.clonesLoke.Unlock()
}

func (s *ItemStorage) GetByName(itemName domain.ItemName) ([]domain.Item, bool) {
	// s.clonesLoke.RLock()
	// defer s.clonesLoke.RUnlock()

	items, exists := s.namesakesRelStore[itemName]
	return items, exists
}

func (s *ItemStorage) GetByID(id domain.ItemID) (domain.Item, bool) {
	// s.loke.RLock()
	// defer s.loke.RUnlock()

	item, ok := s.idOrientedStore[id]
	return item, ok
}

func (s *ItemStorage) GetNames() []domain.ItemName {
	// s.clonesLoke.RLock()
	// defer s.clonesLoke.RUnlock()

	var names []domain.ItemName
	for name := range s.namesakesRelStore {
		names = append(names, name)
	}
	return names
}

func (s *ItemStorage) GetIDs() []domain.ItemID {
	// s.clonesLoke.RLock()
	// defer s.clonesLoke.RUnlock()

	var IDs []domain.ItemID
	for ID := range s.idOrientedStore {
		IDs = append(IDs, ID)
	}
	return IDs
}

func (s *ItemStorage) DumpMap() map[domain.ItemID]domain.Item {
	return s.idOrientedStore
}

func (s *ItemStorage) DumpNamesakesMap() map[domain.ItemName][]domain.Item {
	return s.namesakesRelStore
}
