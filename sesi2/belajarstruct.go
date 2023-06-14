package main

import (
	"fmt"
	"strconv"
)

func main() {
	greet("ronald", "cikupa")

	namalengkap := returnValue("ronald")
	fmt.Println("Single Return : " + namalengkap)

	awal, akhir := returnValueMultiple("ronald")

	fmt.Println("Multiple REturn : " + awal)
	fmt.Println("Multiple REturn : " + akhir)

	fmt.Println("Recursive function " + strconv.Itoa(fibonacci(8)))

}

func fibonacci(n int) int {
	if n == 0 || n == 1 {
		return n
	} else {
		return fibonacci(n-1) + fibonacci(n-2)
	}
}

func greet(name, address string) {
	fmt.Println("Function : " + name)
	fmt.Println("Function : " + address)
}

func returnValue(name string) string {
	namalengkap := "hallo " + name
	return namalengkap
}

func returnValueMultiple(name string) (awal string, akhir string) {
	awal = "hallo " + name
	akhir = "sampai jumpa " + name
	return awal, akhir // atau cukup return saja karena sudah menggunakan predefine return value
}

type User struct {
	name     string
	email    string
	password string
}
