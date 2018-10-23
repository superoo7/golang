package main

import (
	"text/template"
	"os"
	"log"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("main.gohtml"))
}

func main() {
	name := []string{"Person1", "Person2", "Person3", "Person4"}

	err := tpl.Execute(os.Stdout, name)
	if err != nil {
		log.Fatalln(err)
	}
}