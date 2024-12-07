package main

import "fmt"

func main() {
	a := `sadf
			fas 
		df
		`
	for b := range a {
		if a[b] == '\t' {
			fmt.Println("Tab")
		} else if a[b] == '\n' {
			fmt.Println("Newline")
		}
	}
}
