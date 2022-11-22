package main

import (
	"log"

	filetool "github.com/takoyaki-3/file-tool"
	csvtag "github.com/takoyaki-3/go-csv-tag/v3"
)

func main() {

	err, fileList := filetool.DirWalk("/",filetool.DirWalkOption{})
	if err != nil {
		log.Fatalln(err)
	}

	err = csvtag.DumpToFile(fileList,"files.csv")
	if err != nil {
		log.Fatalln(err)
	}
}
