package main

import "fmt"

func main() {
	height := 3

	// Use switch on the height variable.
	switch {
	case height <= 4:
		fmt.Println("Short")
		fallthrough
	case height <= 5:
		fmt.Println("Normal")
		fallthrough
	case height > 5:
		fmt.Println("Tall")
	}
}
