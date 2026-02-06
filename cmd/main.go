package main

import (
	"cloneslasher/internal/adapters/formatter"
	"cloneslasher/internal/adapters/handler"
	storage "cloneslasher/internal/adapters/memstorage"
	"cloneslasher/internal/app"
	"cloneslasher/pkg/slicex"
	"encoding/json"
	"os"
	"path/filepath"
)

func main() {
	run()
}

func run() {
	itemStorage := storage.NewItemStorage()
	fileHandler := handler.NewFileHandler()
	fileHandler.AddHandleFunc(itemStorage.AddItem)
	cloneSeeker := app.NewCloneSeeker(itemStorage, fileHandler)

	path1 := filepath.Join("data", "test_1")
	path2 := filepath.Join("data", "test_2")
	path3 := filepath.Join("data", "test_3")
	fileHandler.Process(path1, path2, path3)

	file, err := os.OpenFile(filepath.Join("data", "namesakes.json"), os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	namesakesDTO := slicex.Map(cloneSeeker.GetItemNamesakes(), formatter.MapItemNamesakesToDTO)
	err = json.NewEncoder(file).Encode(namesakesDTO)
	if err != nil {
		panic(err)
	}

	file, err = os.OpenFile(filepath.Join("data", "clones.json"), os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	clonesDTO := slicex.Map(cloneSeeker.GetItemClones(), formatter.MapItemClonesToDTO)
	err = json.NewEncoder(file).Encode(clonesDTO)
	if err != nil {
		panic(err)
	}
}
