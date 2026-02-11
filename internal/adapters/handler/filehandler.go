package handler

import (
	"cloneslasher/internal/domain"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

type FileHandleFunc func(ownerPath string, item domain.Item)

type FileHandler struct {
	handlerFuncs []FileHandleFunc
}

func NewFileHandler() *FileHandler {
	return &FileHandler{
		handlerFuncs: make([]FileHandleFunc, 0),
	}
}

func (h *FileHandler) AddHandleFunc(handleFunc func(ownerPath string, item domain.Item)) {
	h.handlerFuncs = append(h.handlerFuncs, handleFunc)
}

func (h *FileHandler) Process(paths []string) error {
	for _, path := range paths {
		err := filepath.WalkDir(path,
			func(pathArg string, dirArg fs.DirEntry, errArg error) error {
				if errArg != nil {
					fmt.Printf("preventing walk through %s: %v\n", pathArg, errArg)
					return errArg
				}

				// Process the owner
				ownerID := filepath.Dir(pathArg)

				// Process the item
				stat, errArg := os.Stat(pathArg)
				if errArg != nil {
					fmt.Printf("error getting stat for directory %s: %v\n", pathArg, errArg)
					return errArg
				}
				item := domain.NewItem(pathArg, stat.Name(), filepath.Ext(pathArg), stat.IsDir(), stat.Size())

				for _, handleFunc := range h.handlerFuncs {
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
