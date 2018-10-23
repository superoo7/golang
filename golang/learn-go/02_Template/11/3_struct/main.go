package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type chemical struct {
	Base string
	Acid string
}

func init() {
	tpl = template.Must(template.ParseFiles("main.gohtml"))
}

func main() {
	salt := chemical{
		Base: "NaOH",
		Acid: "HCl",
	}

	err := tpl.Execute(os.Stdout, salt)

	if err != nil {
		log.Fatalln(err)
	}
}
