package init

import (
	"fmt"
	"github.com/nizigama/one/helpers"
	"log"
	"os"
)

type initPaths struct {
	rootPath         string
	structureFolders []string
}

type initFile struct {
	filePath string
	fileData []byte
}

func init() {

	structureConfig := getStructureFolders()

	initFiles := getInitialFiles()

	for _, folderName := range structureConfig.structureFolders {
		err := helpers.CreateFolderIfNotExists(fmt.Sprintf("%s/%s", structureConfig.rootPath, folderName))

		if err != nil {
			log.Fatalln(err)
		}
	}

	for _, file := range initFiles {
		err := helpers.CreateFileIfNotExists(fmt.Sprintf("%s/%s", structureConfig.rootPath, file.filePath), file.fileData)

		if err != nil {
			log.Fatalln(err)
		}
	}
}

func getStructureFolders() initPaths {

	rootPath, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	return initPaths{
		rootPath: rootPath,
		structureFolders: []string{
			"app/console",
			"app/http/controllers",
			"app/http/middlewares",
			"app/http/requests",
			"app/models",
			"app/services",
			"config",
			"database/migrations",
			"public",
			"resources/css",
			"resources/js",
			"resources/views",
			"routes",
			"storage/app",
			"storage/logs",
			"tests/feature",
			"tests/unit",
		},
	}
}

func getInitialFiles() []initFile {

	return []initFile{
		{
			filePath: ".env",
			fileData: []byte(envData),
		},
		{
			filePath: ".env.example",
			fileData: []byte(envData),
		},
		{
			filePath: "storage/logs/one.log",
		},
		{
			filePath: "routes/web.go",
			fileData: []byte(webRouteData),
		},
		{
			filePath: "app/http/main.go",
			fileData: []byte(mainHttpData),
		},
	}
}
