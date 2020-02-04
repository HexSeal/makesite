package main

import (
	// "fmt"
	"html/template"
	"io/ioutil"
	"os"
	"strings"
)

type entry struct {
	Name    string
	Content string
	Done    bool
}

type ToDo struct {
	User string
	List []entry
}

func readFile(file string) string {
	fileContents, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	return string(fileContents)
}

func writeFile() {
	contents := readFile("first-post.txt")
	bytesToWrite := []byte(contents)

	err := ioutil.WriteFile("new-file.txt", bytesToWrite, 0644)
	if err != nil {
		panic(err)
	}
}

func renderTemplate(name string, data interface{}) {
	t := template.Must(template.New("template.tmpl").ParseFiles(name))

	err := t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}

func FileExtensionConverter(name string) {
	return strings.Split(name, ".")[0] + ".html"
}

func writeTemplateToFile(filename string, data interface{}) {
	t := template.Must(template.New("template.tmpl").ParseFiles(filename))

	name := FileExtensionConverter(filename)
	f, err := os.Create(name)
	if err != nil {
		panic(err)
	}

	err = t.Execute(f, data)
	if err != nil {
		panic(err)
	}
}

func main() {
	// Mock Data
	// newEntry := entry{Name: "New ToDo", Done: false, Content: "Finish this project"}
	// entryList := []entry{newEntry}
	// newToDo := ToDo{User: "Max", List: entryList}

	arg := os.Args[1]
	renderTemplate("template.tmpl", readFile(arg))
	writeTemplateToFile("template.tmpl", readFile(arg))
}
