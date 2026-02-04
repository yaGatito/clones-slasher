package app

import (
	"cloneslasher/internal/adapters/handler"
	"cloneslasher/internal/domain"
	"cloneslasher/internal/ports"
	"fmt"
	"slices"
)

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

// Process scans the given directories paths and stores files and folders in the repository, returning a slice of Items found.
func (cs *CloneSeeker) Process(paths ...string) error {
	cs.fileHandler.AddHandleFunc(func(item domain.Item) {
		err := cs.itemRepo.AddItem(item)
		if err != nil {
			fmt.Println("error add item" + err.Error())
		}
	})
	return cs.fileHandler.Process(paths...)
}

type ItemNamesakes struct {
	Name      domain.ItemName
	Namesakes []domain.Item
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

type ItemClones struct {
	Item   domain.Item
	Clones []domain.Item
}

func (cs *CloneSeeker) GetItemClones() []ItemClones {
	var res []ItemClones = make([]ItemClones, 0)

	names := cs.itemRepo.GetNames()
	for _, name := range names {
		namesakes, ok := cs.itemRepo.GetByName(name)

		if ok && len(namesakes) > 0 {
			folderClones := findClones(namesakes...)
			for path, itemClones := range folderClones {
				item, ok := cs.itemRepo.GetByID(domain.ItemID(path))
				if !ok {
					fmt.Println("warning: didnt find item by path: " + path)
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

func findClones(values ...domain.Item) map[domain.ItemID][]domain.Item {
	res := make(map[domain.ItemID][]domain.Item)

	var target domain.Item

	for j := 0; j < len(values)-1; j++ {
		if len(values) <= 1 {
			break
		}

		target = values[0]
		ix := slices.IndexFunc(values[j+1:], target.Same)
		if ix >= 0 {
			res[target.ID] = append(res[target.ID], values[ix])
		}
	}

	// i := 0
	// for {
	// 	if len(values) <= 1 {
	// 		break
	// 	}

	// 	if i == 0 {
	// 		target = item
	// 		continue
	// 	} else if target.Same(values[i]) {
	// 		_, ok := res[target.ID]
	// 		if !ok {
	// 			res[target.ID] = make([]domain.Item, 2)
	// 		}
	// 		res[target.ID] = append(res[target.ID], item)

	// 		continue
	// 	}

	// 	// If hit last index
	// 	if i == (len(values) - 1) {
	// 		i = 0
	// 		values = values[1:]
	// 		continue
	// 	} else {
	// 		i++
	// 	}
	// }

	return res
}
