package main

import "fmt"

func main() {
	fmt.Println("Я main и я начался")
	defer func() {
		fmt.Println("Я main и я закончился")
	}()
	hello()

}

func hello() {
	fmt.Println("Hello world1")
	fmt.Println("Hello world2")
	fmt.Println("Hello world3")

}
