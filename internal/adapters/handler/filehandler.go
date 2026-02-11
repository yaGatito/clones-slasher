package handler

import (
	"cloneslasher/internal/domain"
	"errors"
	"fmt"
	"io/fs"
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

func (h *FileHandler) Process(roots []string) error {
	for _, root := range roots {
		err := filepath.WalkDir(root,
			func(pathArg string, dirEntryArg fs.DirEntry, errArg error) error {
				if errArg != nil {
					if errors.Is(errArg, fs.ErrPermission) {
						return fs.SkipDir
					} else {
						fmt.Printf("warning: preventing walk through %s: %v\n", pathArg, errArg)
						return nil
					}
				}

				// Process the owner
				ownerPath := filepath.Dir(pathArg)

				stat, err := dirEntryArg.Info()
				if err != nil {
					return nil
				}

				// Process the item
				item := domain.NewItem(pathArg, stat.Name(), filepath.Ext(pathArg), stat.IsDir(), stat.Size())
				for _, handleFunc := range h.handlerFuncs {
					handleFunc(ownerPath, *item)
				}

				return nil
			})

		if err != nil {
			return err
		}

	}
	return nil
}
