package main

import (
	// "fmt"
	"html/template"
	"io/ioutil"
	"os"
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

func writeTemplateToFile(filename string, data interface{}) {
	t := template.Must(template.New("template.tmpl").ParseFiles(filename))

	f, err := os.Create("first-post.html")
	if err != nil {
		panic(err)
	}

	err = t.Execute(f, data)
	if err != nil {
		panic(err)
	}
}

func main() {
	newEntry := entry{Name: "New ToDo", Done: false, Content: "Finish this project"}
	entryList := []entry{newEntry}

	newToDo := ToDo{User: "Max", List: entryList}
	renderTemplate("template.tmpl", newToDo)
	writeTemplateToFile("template.tmpl", newToDo)
}
