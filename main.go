package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// fmt.Print("Hello World")
	// fmt.Print("Hello World")
	// fmt.Print("Hello World")'

	// for i := 1; i <= 5; i++ {
	// 	fmt.Print("Hello World")
	// }

	// fmt.Println("До")
	// number:= 999;
	// fmt.Println(number)

	// for i:=1; i<= 5; i++{
	// 	score:= 5;
	// 	score = score +3;
	// 	fmt.Println("Hello world", i)
	// 	fmt.Println("score = ", score)
	// 	fmt.Println("number = ", number)
	// 	number++;
	// }
	// fmt.Println("После number",number);
	//🟥

	// score := 0
	// fmt.Println("Get Ready")
	// fmt.Println("Счет: ", score)
	// fmt.Println("")
	// for i := 1; i <= 5; i++ {

	// 	fmt.Println("==========================================")

	// 	fmt.Println("Вы подлетаете к трубе! ", i)
	// 	fmt.Println("🐥 🟩 🟩 ")
	// 	fmt.Println("")

	// 	fmt.Println("Вы подлетаете через трубу! ", i)
	// 	fmt.Println(" 🟩 🐥 🟩 ")
	// 	fmt.Println("")

	// 	fmt.Println("Вы пролетели через трубу! ", i)
	// 	fmt.Println(" 🟩 🟩 🐥 ")
	// 	fmt.Println("")

	// 	score++

	// 	fmt.Println("Счет: ", score)

	// }

	// fmt.Println("Начинаю уровень генерации трубы!")

	// for i := 1; i <= 10; i++ {
	// 	fmt.Println("Труба номер", i)
	// 	fmt.Println("=========================")
	// 	if i%2 == 0 {
	// 		fmt.Println("🟥 🟥")
	// 	} else {
	// 		fmt.Println("🟩 🟩")

	// 	}

	// 	fmt.Println("=========================")
	// 	fmt.Println("")
	// }

	// fmt.Println("Закончил уровень генерации трубы")

	// fmt.Println("Hello world")
	// fmt.Println("Рандомная цифра", rand.Intn(10))

	score := 0
	fmt.Println("Get Ready")
	fmt.Println("Счет: ", score)
	fmt.Println("")
	for {

		fmt.Println("==========================================")

		fmt.Println("Вы подлетаете к трубе! ")
		fmt.Println("🐥 🟩 🟩 ")
		fmt.Println("")

		fmt.Println("Вы подлетаете через трубу! ")
		fmt.Println(" 🟩 🐥 🟩 ")
		fmt.Println("")

		if rand.Intn(4) == 1 {
			fmt.Println("Я врезался в трубу")

			break
		}

		fmt.Println("Вы пролетели через трубу! ")
		fmt.Println(" 🟩 🟩 🐥 ")
		fmt.Println("")

		score++
		fmt.Println("Счет: ", score)
		fmt.Println("==========================================")
		fmt.Println("")
		time.Sleep(500 * time.Millisecond)
	}

	fmt.Println("GAME OVER", score)

}
