package main

import "fmt"

func main() {
	name := "superoo7"

	template := `
	<!DOCTYPE html>
	<html>
	<head>
	<title>Hello</title>
	</head>
	<body>
	<h1>` + name + `</h1>
	</body>
	</html>
	`

	fmt.Println(template)
}

// go run main1.go > index.html
