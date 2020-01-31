package main

import (
	"fmt"
	"io/ioutil"
)

type entry struct {
	Name string
	Done bool
}

type ToDo struct {
	User string
	List []entry
}

func readFile(file string) {
        fileContents, err := ioutil.ReadFile(file)
        if err != nil {
            panic(err)
        }
        fmt.Print(string(fileContents))
}

func writeFile() {
	contents := readFile("first-post.txt")
	bytesToWrite := []byte(contents)

	err := ioutil.WriteFile("new-file.txt", bytesToWrite, 0644)
	if err != nil {
		panic(err)
	}
}

func main() {
	readFile("first-post.txt")
	writeFile()
}