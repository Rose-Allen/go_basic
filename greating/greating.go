package greating

import "fmt"

func SayHello() {
	fmt.Println("SayHello from greating")
	i := GiveMeInt()
	fmt.Println(i)
}
