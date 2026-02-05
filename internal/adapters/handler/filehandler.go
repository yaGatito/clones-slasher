package handler

import (
	"cloneslasher/internal/domain"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

type FileHandleFunc func(ownerID domain.ItemID, item domain.Item)

type FileHandler struct {
	HandlerFuncs []FileHandleFunc
}

func NewFileHandler() *FileHandler {
	return &FileHandler{
		HandlerFuncs: make([]FileHandleFunc, 0),
	}
}

func (h *FileHandler) AddHandleFunc(handleFunc func(ownerID domain.ItemID, item domain.Item)) {
	h.HandlerFuncs = append(h.HandlerFuncs, handleFunc)
}

func (h *FileHandler) Process(paths ...string) error {
	for _, path := range paths {
		err := filepath.WalkDir(path,
			func(pathArg string, dirArg fs.DirEntry, errArg error) error {
				if errArg != nil {
					fmt.Printf("preventing walk through %s: %v\n", pathArg, errArg)
					return errArg
				}

				// Process the owner
				ownerID := domain.ItemID(filepath.Dir(pathArg))

				// Process the item
				stat, errArg := os.Stat(pathArg)
				if errArg != nil {
					fmt.Printf("error getting stat for directory %s: %v\n", pathArg, errArg)
					return errArg
				}
				item := domain.NewItem(domain.ItemID(pathArg), domain.ItemName(stat.Name()), filepath.Ext(pathArg), stat.IsDir(), stat.Size())

				for _, handleFunc := range h.HandlerFuncs {
					handleFunc(ownerID, *item)
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
