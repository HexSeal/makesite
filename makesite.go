package main

import (
	// "fmt"
	"html/template"
	"io/ioutil"
	"os"
	"strings"
	"flag"
	//"github.com/kr/fs"
)

type entry struct {
	Name    string
	Content string
	Done    bool
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

func FileExtensionConverter(name string) string {
	extension := ".html"
	s := strings.Split(name, ".")[0] + extension
	return s
}

func writeTemplateToFile(templateName string, fileName string) {
	c := entry{Content: readFile(fileName)}
	t := template.Must(template.New("template.tmpl").ParseFiles(templateName))

	name := FileExtensionConverter(fileName)
	f, err := os.Create(name)
	if err != nil {
		panic(err)
	}

	err = t.Execute(f, c)
	if err != nil {
		panic(err)
	}
}

func main() {
	// Mock Data
	// newEntry := entry{Name: "New ToDo", Done: false, Content: "Finish this project"}
	// entryList := []entry{newEntry}
	// newToDo := ToDo{User: "Max", List: entryList}

	var filename string
	flag.StringVar(&filename, "filename", "new-post.txt", "The name of the text file we want to save in HTML")

	arg := os.Args[1]
// 	if err := fs.Parse(os.Args[1:]); err != nil {
// 		os.Exit(100)
// }
	writeTemplateToFile("template.tmpl", arg)
}
