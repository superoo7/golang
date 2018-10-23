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
	name := map[string]string {
		"HCl": "Hydrocloric Acid",
		"NaCl": "Sodium Chloride",
		"H2O": "Water"	}

	err := tpl.Execute(os.Stdout, name)

	if err != nil {
		log.Fatalln(err)
	}
}