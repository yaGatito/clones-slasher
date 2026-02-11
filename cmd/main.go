package main

import (
	"cloneslasher/internal/adapters/handler"
	storage "cloneslasher/internal/adapters/memstorage"
	"cloneslasher/internal/adapters/terminal"
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

	cmd := terminal.ParseArgs(os.Args)
	err := cloneSeeker.ProcessCommand(cmd)
	if err != nil {
		panic(err)
	}

	reportPath := cmd.ReportPath
	if reportPath == "" {
		reportPath = filepath.Dir(cmd.Paths[0])
	}

	if cmd.CloneCheck {
		cloneSeeker.ReportClones(reportPath)
	}

	if cmd.NamesakesCheck {
		cloneSeeker.ReportNamesakes(reportPath)
	}
}
