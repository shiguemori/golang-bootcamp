package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	newFile, err := os.Create("a.txt")

	if err != nil {
		log.Fatal(err)
	}

	err = os.Truncate("a.txt", 0)

	if err != nil {
		log.Fatal(err)
	}

	newFile.Close()

	file, err := os.Open("a.txt")

	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	file, err = os.OpenFile("a.txt", os.O_APPEND, 0644)

	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	var fileInfo os.FileInfo
	fileInfo, err = os.Stat("a.txt")

	p := fmt.Println
	p("File Name:", fileInfo.Name())
	p("Size in bytes:", fileInfo.Size())
	p("Last modified:", fileInfo.ModTime())
	p("Is Directory? ", fileInfo.IsDir())
	p("Pemissions:", fileInfo.Mode())

	fileInfo, err = os.Stat("b.txt")

	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("The file does not exist")
		}
	}

	oldPath := "a.txt"
	newPath := "aaa.txt"
	err = os.Rename(oldPath, newPath)

	if err != nil {
		log.Fatal(err)
	}

	err = os.Remove("aa.txt")

	if err != nil {
		log.Fatal(err)
	}

}
