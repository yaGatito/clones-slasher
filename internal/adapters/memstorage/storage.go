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

func (s *ItemStorage) AddItem(ownerID domain.ItemID, item domain.Item) {
	owner, ok := s.GetByID(ownerID)
	if !ok {
		fmt.Println("failed to get owner item" + ownerID)
	}
	owner.Content = append(owner.Content, item.ID)
	s.updateByID(owner)

	err := s.addItem(item)
	if err != nil {
		fmt.Println("add item:" + err.Error())
	}
	err = s.addItemToNamesakes(item)
	if err != nil {
		fmt.Println("add item to namesakes:" + err.Error())
	}
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

func (s *ItemStorage) updateByID(item domain.Item) bool {
	_, ok := s.idOrientedStore[item.ID]
	if !ok {
		return false
	}
	s.idOrientedStore[item.ID] = item
	return true
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

func (s *ItemStorage) dumpMap() map[domain.ItemID]domain.Item {
	return s.idOrientedStore
}

func (s *ItemStorage) dumpNamesakesMap() map[domain.ItemName][]domain.Item {
	return s.namesakesRelStore
}
