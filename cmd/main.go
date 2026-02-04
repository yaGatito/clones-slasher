package main

import (
	"cloneslasher/internal/adapters/formatter"
	"cloneslasher/internal/adapters/handler"
	storage "cloneslasher/internal/adapters/memstorage"
	"cloneslasher/internal/app"
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
	namesakesRawData := cloneSeeker.GetItemNamesakes()

	foldersRes := formatter.MapCollectionToDTO(namesakesRawData)
	// fmt.Println("=========================")
	// for _, itemNamesakes := range foldersRes {
	// 	fmt.Printf("name: %s; namesakes: %v\n", itemNamesakes.Name, itemNamesakes.Namesakes)
	// }
	// fmt.Println("=========================")
	bytes, err := formatter.MapToJson(foldersRes)
	jsonPath := filepath.Join("data", "namesakes.json")
	err = os.WriteFile(jsonPath, bytes.Bytes(), 0666)
	if err != nil {
		panic(err)
	}

	// clones := cloneSeeker.GetItemClones()
	// fmt.Println(clones)
}
