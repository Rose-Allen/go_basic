package main

import (
	"fmt"
	"study/greating"
	"study/user"

	"github.com/k0kubun/pp"
	// "github.com/k0kubun/pp"
)

func main() {
	greating.SayHello()
	greating.BadMessage()

	// u := user.User{'Max', 25}
	u := user.NewUser("Nax", 25)
	fmt.Println(u)
	pp.Println(u)

	// u.Age = 500
	// fmt.Println(u)
}
