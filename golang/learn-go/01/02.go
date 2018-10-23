package main

import "fmt"

func main() {

	test := 3
	change(&test)

	fmt.Println(test)
}

func change(a *int) {
	*a = *a + *a
}
