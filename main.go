package main

import "fmt"

type User struct {
	Name        string
	Age         int
	PhoneNumber string
	isClose     bool
	Rating      float64
}

func main() {
	user := User{
		Name:        "Dimash",
		Age:         25,
		PhoneNumber: "7078886644",
		isClose:     false,
		Rating:      5.5,
	}
	fmt.Print(user)
}
