package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("main.gohtml"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "main.gohtml", `test test test`)
	if err != nil {
		log.Fatalln(err)
	}
}
