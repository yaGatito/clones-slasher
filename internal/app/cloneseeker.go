package app

import (
	"cloneslasher/internal/domain"
	"cloneslasher/internal/ports"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

type CloneSeeker struct {
	fileRepo   ports.ItemRepository
	folderRepo ports.ItemRepository
}

func NewCloneSeeker(fileRepo, folderRepo ports.ItemRepository) *CloneSeeker {
	return &CloneSeeker{fileRepo: fileRepo, folderRepo: folderRepo}
}

// Process scans the given directories paths and stores files and folders in the repository, returning a slice of Items found.
func (cs *CloneSeeker) Process(paths ...string) error {
	for _, path := range paths {
		err := filepath.WalkDir(path,
			func(pathArg string, dirArg fs.DirEntry, errArg error) error {
				if errArg != nil {
					fmt.Printf("preventing walk through %s: %v\n", pathArg, errArg)
					return errArg
				}

				// Process the file or directory
				stat, errArg := os.Stat(pathArg)
				if errArg != nil {
					fmt.Printf("error getting stat for directory %s: %v\n", pathArg, errArg)
					return errArg
				}

				if stat.IsDir() {
					folderItem := domain.NewFolder(pathArg, stat.Name(), stat.Size(), nil)
					cs.folderRepo.AddItem(*folderItem)
				} else {
					fileItem := domain.NewFile(pathArg, stat.Name(), stat.Size(), filepath.Ext(pathArg))
					cs.fileRepo.AddItem(*fileItem)
				}

				return errArg
			})

		if err != nil {
			log.Fatalf("error walking the path %s: %v\n", path, err)
			return err
		}

	}
	return nil
}

func (cs *CloneSeeker) GetFoldersNamesakes() map[string][]domain.Item {
	var res map[string][]domain.Item = make(map[string][]domain.Item)

	for _, name := range cs.folderRepo.GetNames() {
		namesakes, ok := cs.folderRepo.GetByName(name)
		if !ok {
			fmt.Println("didnt find folder namesakes")
		}

		if len(namesakes) > 1 {
			res[name] = namesakes
		}
	}

	return res
}

// type ItemClones struct {
// 	Path      string
// 	Name      string
// 	Size      int64
// 	Extension string
// 	IsFolder  bool
// 	Content   []string
// 	Clones    []string
// }

// TODO: fix the problem: create some struct here to output data. (ItemClones) and some field `clones []string`
// func (cs *CloneSeeker) GetFolderClones() []ItemClones {
// 	var res []ItemClones = make([]ItemClones, 0)

// 	names := cs.folderRepo.GetNames()
// 	for _, name := range names {
// 		folders, err := cs.folderRepo.GetByName(name)
// 		if err != nil {
// 			panic(err)
// 		}

// 		if len(folders) > 1 {
// 			exactFoldersClones := findExactClones(folders...)
// 			for path, itemClones := range exactFoldersClones {

// 				item, ok := cs.folderRepo.GetByPath(path)
// 				if !ok {
// 					fmt.Println("warning: didnt find item by path: " + path)
// 				}
// 			}
// 		}
// 	}

// 	return res
// }

func findExactClones(values ...domain.Item) map[string][]domain.Item {
	res := make(map[string][]domain.Item)

	var target domain.Item
	i := 0
	for {
		if len(values) <= 1 {
			break
		}

		item := values[i]
		if i == 0 {
			target = item
			continue
		} else if target.Equals(values[i]) {
			clones, ok := res[target.Path]
			if !ok {
				clones = make([]domain.Item, 2)
				clones[0] = target
				clones[1] = item
				res[target.Path] = clones
			} else {
				res[target.Path] = append(res[target.Path], item)
			}
			continue
		}

		// If hit last index
		if i == (len(values) - 1) {
			i = 0
			values = values[1:]
			continue
		} else {
			i++
		}
	}

	return res
}

func (cs *CloneSeeker) FileClones() map[string][]domain.Item {
	var res map[string][]domain.Item = make(map[string][]domain.Item)

	keys := cs.fileRepo.GetNames()

	for _, k := range keys {
		files, ok := cs.fileRepo.GetByName(k)
		if !ok {
			fmt.Println("didnt find clones for file name:" + k)
			continue
		}
		res[k] = files
	}

	return res
}
