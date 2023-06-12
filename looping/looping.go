package main

import (
	"fmt"
	"strconv"
)

func main() {
	iniLooping()
}

func iniLooping() {
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

func iniLoopingWitParameter(value int) string {
	if value%3 == 0 && value%5 == 0 {
		return "Fizzbuzz"
	} else if value%5 == 0 {
		return "buzz"
	} else if value%3 == 0 {
		return "fizz"
	} else {
		return strconv.Itoa(value)
	}

}
