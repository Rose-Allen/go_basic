package main

import "fmt"

var number int = 5

func main() {
	// fmt.Print("Hello World")
	// fmt.Print("Hello World")
	// fmt.Print("Hello World")

	// fmt.Println("Start func main")
	// square(5)
	// fmt.Println("End of main")
	// fmt.Println("")
	// hello()

	// officianWork("Максим")

	summa := sum(2, 2)
	fmt.Println(summa)
}

func square(x int) {
	fmt.Println("Принимаем переменную  - ", x)
	fmt.Println("Квадрат - ", x*x)
}

func hello() {
	fmt.Println("Hello world")
}

func officianWork(name string) {
	fmt.Println("Накрываю на стол")
	fmt.Println("Привет", name)
	fmt.Println("Я принял заказ!")
	fmt.Println("Я принес блюдо")
	fmt.Println("")

}

func sum(a int, b int) int {
	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
	c := a + b
	return c
}
