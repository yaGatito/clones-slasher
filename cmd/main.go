package main

import (
	"cloneslasher/internal/adapters/formatter"
	storage "cloneslasher/internal/adapters/memstorage"
	"cloneslasher/internal/app"
	"os"
	"path/filepath"
)

func main() {
	run()
	// arr := []domain.Item{
	// 	domain.Item{
	// 		Path:      "asdasd",
	// 		Name:      "asdasasd",
	// 		Size:      213,
	// 		Extension: ".weqq",
	// 		IsFolder:  false,
	// 	},
	// 	domain.Item{
	// 		Path:      "asdasd",
	// 		Name:      "123ew12ew",
	// 		Size:      1231231,
	// 		Extension: ".asd",
	// 		IsFolder:  false,
	// 	},
	// 	domain.Item{
	// 		Path:      "asdasd",
	// 		Name:      "asdasasd",
	// 		Size:      213,
	// 		IsFolder:  true,
	// 		Content: []domain.Item{
	// 			domain.Item{
	// 				Path: "asdasd",
	// 				Name: "sdad",
	// 			},
	// 		},
	// 	},
	// }
	// fmt.Println(arr[1:][0])
}

func run() {
	fileStorage := storage.NewItemStorage()
	folderStorage := storage.NewItemStorage()
	cloneSeeker := app.NewCloneSeeker(fileStorage, folderStorage)

	path1 := filepath.Join("data", "test_2")
	path2 := filepath.Join("data", "test_2")
	path3 := filepath.Join("data", "test_3")

	err := cloneSeeker.Process(path1, path2, path3)
	if err != nil {
		panic(err)
	}

	namesakesRawData := cloneSeeker.GetFoldersNamesakes()
	foldersRes := formatter.MapCollectionToDTO(namesakesRawData)
	bytes, err := formatter.MapToJson(foldersRes)
	jsonPath := filepath.Join("data", "folder_namesakes.json")
	err = os.WriteFile(jsonPath, bytes.Bytes(), 0666)
	if err != nil {
		panic(err)
	}

	// clones := cloneSeeker.GetFolderClones()
}
