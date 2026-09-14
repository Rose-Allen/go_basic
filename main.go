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

	res := multiply(5, 4)
	fmt.Println(res)

	fmt.Print(isEven(10))
	fmt.Print(isEven(7))

	fmt.Println(max(10, 20))
	fmt.Println(max(50, 7))

	fmt.Println(canLogin(25, false)) // true
	fmt.Println(canLogin(16, false)) // false
	fmt.Println(canLogin(25, true))  // false

	name, age := getPerson()
	fmt.Println(name)
	fmt.Println(age)

	sum, multiply := calculate(5, 3)

	fmt.Println(sum)      // 8
	fmt.Println(multiply) // 15

	_, age2 := getPerson()
	fmt.Println(age2)

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

func multiply(a int, b int) int {
	res := a * b
	return res
}

func isEven(num int) bool {
	if num%2 == 0 {
		return true
	} else {
		return false
	}
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func canLogin(age int, isBlocked bool) bool {
	if age >= 18 && !isBlocked {
		return true
	} else {
		return false
	}
}

func getPerson() (string, int) {
	return "Dimash", 25
}

func calculate(a int, b int) (int, int) {
	return a + b, a * b
}
