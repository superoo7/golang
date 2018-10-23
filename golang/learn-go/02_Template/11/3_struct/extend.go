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
	tpl = template.Must(template.ParseFiles("extend.gohtml"))
}

func main() {
	salt := chemical{
		Base: "NaOH",
		Acid: "HCl",
	}


	kcl := chemical{
		Base: "KOH",
		Acid: "HCl",
	}

	nacl := chemical{
		Base: "NaOH",
		Acid: "HCl",
	}

	chemicals := []chemical{salt, kcl, nacl}
	err := tpl.Execute(os.Stdout, chemicals)

	if err != nil {
		log.Fatalln(err)
	}
}
