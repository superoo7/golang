package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}


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
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("index.gohtml"))
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	if len(s) >= 3 {
		s = s[:3]
	}
	return s
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
		Model:        "Accord",
		Doors:        4,
	}

	saga := car{
		Manufacturer: "Proton",
		Model:        "Saga",
		Doors:        4,
	}

	chem := []chemical{nacl, kcl}
	cars := []car{accord, saga}

	data := item{chem, cars}

	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", data)

	if err != nil {
		log.Fatalln(err)
	}
}
