package main

import "fmt"

func main() {
	// fmt.Println("Я main и я начался")
	// defer func() {
	// 	fmt.Println("Я main и я закончился")
	// }()
	// hello()

	// number := 10

	// Здесь мы получаем адрес нашей перменной
	// pointer := &number
	// fmt.Println(pointer)

	// pointer := &number
	// foo(pointer)

	// number := 5
	// pointer := &number
	// foo(pointer)
	// fmt.Println(number)
	// fmt.Println("Разъяемененный поинтер", *pointer)

	name := "Dimash"
	ptr := &name
	fmt.Println("Before change", name)
	changeName(ptr)
	fmt.Println("After changed", name)
}

func changeName(name *string) {
	*name = "Almas"
}

// а здесь вы указываем
func foo(n *int) {
	fmt.Println(n)
	fmt.Println(*n)

	*n = 10

}

func boo(n int) {
	n = 10
}

func hello() {
	fmt.Println("Hello world1")
	fmt.Println("Hello world2")
	fmt.Println("Hello world3")

}
