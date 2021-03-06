package main

import (
	//	"io/ioutil"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	//"strconv"
)

var (
	newFile *os.File
	err     error
	scan    string
)

var (
	fileInfo os.FileInfo
)

func getInfo(file string) {
	fileInfo, err = os.Stat(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("File Name:", fileInfo.Name())
	fmt.Println("Size in bytes:", fileInfo.Size())
	fmt.Println("Permissions:", fileInfo.Mode())
	fmt.Println("Last Modified:", fileInfo.ModTime())
	fmt.Println("Is Directory: ", fileInfo.IsDir())
	fmt.Printf("System Interface: %T\n", fileInfo.Sys())
	fmt.Printf("System Info: %+v\n\n", fileInfo.Sys())
}

func scanner() string {
	fmt.Printf("Name and path for new file: ")
	reader := bufio.NewReader(os.Stdin)
	//var path string
	scan, _ := reader.ReadString('\n')
	scan = strings.Replace(scan, "\n", "", -1)
	file := scan
	return file
}

func NewFile() {
	//Sfmt.Printf("Name and path for new file")
	file := scanner()
	//path := strings.Replace(scan, "\n", "", -1)
	fmt.Println("The file you are creating is", file)
	newFile, err = os.Create(file)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(newFile)
	newFile.Close()
}

func truncate() {
	err := os.Truncate("test.txt", 100)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	fmt.Println("File Manager")
	fmt.Println("---------------------")
	fmt.Println("Press 1 for new file")
	fmt.Println("Press 2 for file info")
	fmt.Println("Press 3 for file truncation")
	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("key", char)

	switch char {
	case '1':
		NewFile()
		break
	case '2':
		getInfo(file)
	case '3':
		truncate()
		break
	}

}
