package main

import (
	"flag"
	"fmt"

	"task-7/download"
)


func main() {

	fileURL := flag.String("url", "", "url")
	fileName := flag.String("file", "", "file name")

	flag.Parse()

	err := download.DownloadFile(*fileURL, *fileName)

	if err != nil {
		fmt.Println(err)
	}
	
}


