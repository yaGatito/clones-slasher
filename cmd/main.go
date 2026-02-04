package main

import (
	"cloneslasher/internal/adapters/formatter"
	"cloneslasher/internal/adapters/handler"
	storage "cloneslasher/internal/adapters/memstorage"
	"cloneslasher/internal/app"
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
	cloneSeeker := app.NewCloneSeeker(itemStorage, fileHandler)

	path1 := filepath.Join("data", "test_1")
	path2 := filepath.Join("data", "test_2")
	path3 := filepath.Join("data", "test_3")

	err := cloneSeeker.Process(path1, path2, path3)
	if err != nil {
		panic(err)
	}

	// TODO: Content formation!!! Then it needs to be able to count size!!! by size the equality
	namesakes := cloneSeeker.GetItemNamesakes()
	file, err := os.OpenFile(filepath.Join("data", "namesakes.json"), os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	namesakesDTO := formatter.MapItemNamesakesToDTO(namesakes)
	err = json.NewEncoder(file).Encode(namesakesDTO)
	if err != nil {
		panic(err)
	}

	clones := cloneSeeker.GetItemClones()
	file, err = os.OpenFile(filepath.Join("data", "clones.json"), os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	clonesDTO := formatter.MapItemClonesToDTO(clones)
	err = json.NewEncoder(file).Encode(clonesDTO)
	if err != nil {
		panic(err)
	}

}
