package main

import "fmt"

func main() {
	for counter := 1; counter <= 15; counter++ {
		if counter%3 == 0 && counter%5 == 0 {
			fmt.Println("Fizzbuzz")
		} else if counter%5 == 0 {
			fmt.Println("buzz")
		} else if counter%3 == 0 {
			fmt.Println("fizz")
		} else {
			fmt.Println(counter)
		}

	}
}
