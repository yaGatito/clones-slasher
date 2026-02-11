package app

import (
	"cloneslasher/internal/adapters/formatter"
	"cloneslasher/internal/adapters/handler"
	"cloneslasher/internal/adapters/terminal"
	"cloneslasher/internal/domain"
	"cloneslasher/internal/ports"
	"encoding/json"
	"fmt"
	"maps"
	"os"
	"path/filepath"
	"slices"

	"github.com/yaGatito/slicex"
)

const (
	ClonesReportFile    = "clones.json"
	NamesakesReportFile = "namesakes.json"
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

func (cs *CloneSeeker) ProcessCommand(cmd terminal.Command) error {
	cs.fileHandler.AddHandleFunc(cs.itemRepo.AddItem)

	err := cs.fileHandler.Process(cmd.Paths)
	if err != nil {
		return err
	}

	return nil
}

func (cs *CloneSeeker) ReportClones(outputPathDir string) {
	clonesDTO := slicex.Map(cs.getItemClones(), formatter.MapItemClonesToDTO)
	cs.reportFile(outputPathDir, ClonesReportFile, clonesDTO)
}

func (cs *CloneSeeker) ReportNamesakes(outputPathDir string) {
	namesakes := slicex.Map(cs.getItemNamesakes(), formatter.MapItemNamesakesToDTO)
	cs.reportFile(outputPathDir, NamesakesReportFile, namesakes)
}

func (cs *CloneSeeker) reportFile(outputDit, fileName string, v any) {
	file, err := os.OpenFile(filepath.Join(outputDit, fileName), os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	err = json.NewEncoder(file).Encode(v)
	if err != nil {
		panic(err)
	}
}

func (cs *CloneSeeker) getItemNamesakes() []domain.ItemNamesakes {
	res := make([]domain.ItemNamesakes, 0)

	for _, name := range cs.itemRepo.GetNames() {
		namesakes, ok := cs.itemRepo.GetByName(name)
		if !ok {
			fmt.Println("warning: didnt find folder namesakes")
		}

		if len(namesakes) > 1 {
			res = append(res, domain.ItemNamesakes{
				Name:      name,
				Namesakes: namesakes,
			})
		}
	}

	return res
}

func (cs *CloneSeeker) getItemClones() []domain.ItemClones {
	res := make([]domain.ItemClones, 0)

	names := cs.itemRepo.GetNames()
	for _, name := range names {
		namesakes, ok := cs.itemRepo.GetByName(name)

		if ok && len(namesakes) > 0 {
			clones := findClones(namesakes)
			if len(clones) == 0 {
				continue
			}
			for itemID, itemClones := range clones {
				// if len(itemClones) > 0 {
				item, ok := cs.itemRepo.GetByPath(itemID.UniquePath)
				if !ok {
					fmt.Println("warning: didnt find item by path: " + itemID.UniquePath)
				}
				res = append(res, domain.ItemClones{
					Item:   item,
					Clones: itemClones,
				})
				// }
			}
		} else {
			fmt.Println("warning: didnt folders namesakes by name: " + name)
		}
	}

	return res
}

func findClones(items []domain.Item) map[domain.ItemID][]domain.Item {
	res := make(map[domain.ItemID][]domain.Item)

	var target domain.Item
	for j, item := range items {
		target = item
		exist := slices.ContainsFunc(slices.Collect(maps.Keys(res)), target.ItemID.IsClone)
		if exist {
			continue
		}

		clones := slicex.Filter(items[j:], target.IsClone)
		res[target.ItemID] = clones
	}

	return res
}
