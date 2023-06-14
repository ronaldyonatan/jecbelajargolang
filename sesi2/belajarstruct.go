package main

import (
	"fmt"
	"strconv"
)

type User struct {
	Name     string
	Email    string
	Password string
}

var users []User

func main() {
	greet("ronald", "cikupa")

	namalengkap := returnValue("ronald")
	fmt.Println("Single Return : " + namalengkap)

	awal, akhir := returnValueMultiple("ronald")

	fmt.Println("Multiple REturn : " + awal)
	fmt.Println("Multiple REturn : " + akhir)

	fmt.Println("Recursive function " + strconv.Itoa(fibonacci(8)))

	//trial struct dan slice
	user1 := User{Name: "Tuan", Email: "tuan@tuan.com", Password: "password1"}
	user2 := User{Name: "Ronald", Email: "ronald@ronald.com", Password: "password2"}
	user3 := User{Name: "Yonatan", Email: "yonatan@yonatan.com", Password: "password3"}

	Register(user1)
	Register(user2)
	Register(user3)

	userList := Get()
	for _, user := range userList {
		fmt.Printf("Name: %s, Email: %s, Password: %s\n", user.Name, user.Email, user.Password)
	}
}

func Register(user User) {
	users = append(users, user)
}

func Get() []User {
	return users
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
