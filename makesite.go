package main

import (
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"strings"
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

func writeTranslate(filename string, lang string) {
	FileText := readFile(filename)
	model := "nmt"

	contents, error := translateText(lang, FileText, model)
	if error != nil {
		panic(error)
	}
	bytesToWrite := []byte(contents)

	err := ioutil.WriteFile(filename, bytesToWrite, 0644)
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
	// Converts a file extension from .txt to .html, or any other file you want set by extension.
	extension := ".html"
	s := strings.Split(name, ".")[0] + extension
	return s
}

func TextFileCheck(name string) bool {
	// Returns true if the file is a text file, otherwise false
	extension := "txt"
	for i := range name {
		if name[i] == '.' {
			s := strings.Split(name, ".")[1]
			if s == extension {
				return true
			}
		}
	}
	return false
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

func parser() {
	var dir string
	flag.StringVar(&dir, "dir", ".", "The directory with the text files that we're going to convert to HTML. Default is current directory. Use -dir=")

	var lang string
	flag.StringVar(&lang, "t", "en", "The language the text will be translated into. Default is english. Use -t=, followed by Google's language abbreviation.")
	flag.Parse()

	fmt.Println("Input:", dir)

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if TextFileCheck(file.Name()) == true {
			writeTranslate(file.Name(), lang)
			fmt.Println(file.Name())
			writeTemplateToFile("template.tmpl", file.Name())
		}
	}
}

func main() {
	parser()
	//translateText("es", "Hello world!", "nmt")
}
