package main

import (
	"text/template"
	"os"
	"log"
)

var tpl *template.Template

type car struct {
	Manufacturer string
	Model        string
	Doors        int
}

type chemical struct {
	Base string
	Acid string
}

type item struct {
	Salt      []chemical
	Transport []car
}

func init() {
	tpl = template.Must(template.ParseFiles("main.gohtml"))
}

func main() {
	nacl := chemical{
		Base: "NaOH",
		Acid: "HCl",
	}

	kcl := chemical{
		Base: "KOH",
		Acid: "HCl",
	}

	accord := car{
		Manufacturer: "Honda",
		Model: "Accord",
		Doors: 4,
	}

	saga := car{
		Manufacturer: "Proton",
		Model: "Saga",
		Doors: 4,
	}

	chem := []chemical{nacl, kcl}
	cars := []car{accord, saga}

	data := item{chem, cars,}

	err := tpl.Execute(os.Stdout, data)

	if err != nil {
		log.Fatalln(err)
	}

}
