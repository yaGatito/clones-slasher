package app

import (
	"cloneslasher/internal/adapters/handler"
	"cloneslasher/internal/domain"
	"cloneslasher/internal/ports"
	"cloneslasher/pkg/slicex"
	"fmt"
	"maps"
	"slices"
)

type ItemNamesakes struct {
	Name      domain.ItemName
	Namesakes []domain.Item
}

type ItemClones struct {
	Item   domain.Item
	Clones []domain.Item
}

type CloneSeeker struct {
	itemRepo    ports.ItemRepository
	fileHandler *handler.FileHandler
}

func NewCloneSeeker(
	itemRepo ports.ItemRepository,
	fileHandler *handler.FileHandler,
) *CloneSeeker {
	return &CloneSeeker{itemRepo: itemRepo, fileHandler: fileHandler}
}

func (cs *CloneSeeker) GetItemNamesakes() []ItemNamesakes {
	var res []ItemNamesakes = make([]ItemNamesakes, 0)

	for _, name := range cs.itemRepo.GetNames() {
		namesakes, ok := cs.itemRepo.GetByName(name)
		if !ok {
			fmt.Println("didnt find folder namesakes")
		}

		if len(namesakes) > 1 {
			res = append(res, ItemNamesakes{
				Name:      name,
				Namesakes: namesakes,
			})
		}
	}

	return res
}

func (cs *CloneSeeker) GetItemClones() []ItemClones {
	var res []ItemClones = make([]ItemClones, 0)

	names := cs.itemRepo.GetNames()
	for _, name := range names {
		namesakes, ok := cs.itemRepo.GetByName(name)

		if ok && len(namesakes) > 0 {
			clones := findClones(namesakes...)
			for itemID, itemClones := range clones {
				item, ok := cs.itemRepo.GetByPath(itemID.UniquePath)
				if !ok {
					fmt.Println("warning: didnt find item by path: " + itemID.UniquePath)
				}
				res = append(res, ItemClones{
					Item:   item,
					Clones: itemClones,
				})
			}
		} else {
			fmt.Println("warning: didnt folders namesakes by name: " + name)
		}
	}

	return res
}

func findClones(items ...domain.Item) map[domain.ItemID][]domain.Item {
	res := make(map[domain.ItemID][]domain.Item)

	var target domain.Item
	for j, item := range items {
		target = item
		exist := slices.ContainsFunc(slices.Collect(maps.Keys(res)), target.ItemID.Same)
		if exist {
			continue
		}

		sameItems := slicex.Filter(items[j:], target.Same)
		res[target.ItemID] = sameItems
	}

	return res
}
