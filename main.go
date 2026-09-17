package main

import "fmt"

type User struct {
	Name        string
	Age         int
	PhoneNumber string
	IsClose     bool
	Rating      float64
}

func NewUser(
	name string,
	age int,
	phoneNumber string,
	isClose bool,
	rating float64,
) User {
	fmt.Println("Валидирую имя: ")
	if name == "" {
		fmt.Println("Не прошло Валидацию имя ")
		return User{}
	}

	fmt.Println("Валидирую возраст: ")
	if age <= 0 || age >= 150 {
		fmt.Println("Не прошло Валидацию возраст ")
		return User{}
	}

	fmt.Println("Валидирую телефон: ")
	if phoneNumber == "" {
		fmt.Println("Не прошло Валидацию телефона ")
		return User{}
	}
	fmt.Println("Валидирую рейтинг: ")
	if rating < 0.0 || rating >= 10.0 {
		fmt.Println("Не прошло Валидацию рейтинга ")
		return User{}
	}

	return User{
		Name:        name,
		Age:         age,
		PhoneNumber: phoneNumber,
		IsClose:     isClose,
		Rating:      rating,
	}

}

func (u *User) ChangeName(newName string) {
	if newName != "" {
		u.Name = newName
	}

}

func (u *User) ChangeAge(newAge int) {
	if newAge > 0 && newAge <= 150 {
		u.Age = newAge
	}
}

func (u *User) ChangePhoneNumber(newPhoneNumber string) {
	if newPhoneNumber != "" {
		u.PhoneNumber = newPhoneNumber
	}
}

func (u *User) CloseAccount() {
	u.IsClose = true
}

func (u *User) OpenAccount() {
	u.IsClose = false
}

func (u *User) RatingUp(rating float64) {
	if u.Rating+rating <= 10.0 {
		u.Rating += rating
	}
}

func (u *User) RatingDown(rating float64) {
	if u.Rating+rating >= 0.0 {
		u.Rating -= rating
	}
}

// func Greeting(u User) {
// 	fmt.Println("Hello my name is ,", u.Name)
// 	fmt.Println("My Rating is ,", u.Rating)

// 	u.Rating = 3.0
// 	fmt.Println("")

// }

// Метод структуры юзер
func (u User) Greeting() {
	fmt.Println("Hello my name is ,", u.Name)
	fmt.Println("My Rating is ,", u.Rating)

	u.Rating = 3.0
	fmt.Println("")

}

// Метод структуры юзер
func (u User) Goodbye() {
	fmt.Println("Bye my name is ,", u.Name)
	fmt.Println("My Rating is ,", u.Rating)
	fmt.Println("")
}

// func Goodbye(u User) {
// 	fmt.Println("Bye my name is ,", u.Name)
// 	fmt.Println("My Rating is ,", u.Rating)
// 	fmt.Println("")
// }

// func (u User) RatingUp(rating float64) {
// 	if u.Rating+rating <= 10.0 {
// 		u.Rating = u.Rating + rating
// 		fmt.Println("I added Rating")
// 	} else {
// 		fmt.Println("I don't approve validation")
// 	}
// }

// Если указатель есть мы можем поменять то что внутри
// func RatingUp(u *User, rating float64) {
// 	if u.Rating+rating <= 10.0 {
// 		u.Rating = u.Rating + rating
// 		fmt.Println("I added Rating")
// 	} else {
// 		fmt.Println("I don't approve validation")
// 	}
// }

// func (u *User) RatingUp(rating float64) {
// 	if u.Rating+rating <= 10.0 {
// 		u.Rating = u.Rating + rating
// 		fmt.Println("I added Rating")
// 	} else {
// 		fmt.Println("I don't approve validation")
// 	}
// }

func main() {
	// user := User{
	// 	Name:        "Dimash",
	// 	Age:         25,
	// 	PhoneNumber: "7078886644",
	// 	isClose:     false,
	// 	Rating:      5.5,
	// }
	// fmt.Println(user)

	// fmt.Println("Before change", user.Name)
	// user.Name = "Vasya"
	// fmt.Println("After change", user.Name)

	// user := User{
	// 	Name:   "Vasya",
	// 	Rating: 6.0,
	// }
	// fmt.Println("User before", user)
	// fmt.Println("")

	// // ptr := &user
	// // Greeting(user)
	// user.Greeting()
	// user.Goodbye()
	// // RatingUp(ptr, 2.0)
	// // user.RatingUp(ptr, 2.0)
	// // Goodbye(user)
	// user.RatingUp(3.0)
	// fmt.Println("User after", user)

	user := NewUser(
		"Dimash",
		25,
		"77082223311",
		false,
		5.5,
	)

	fmt.Println("User: ", user)
}
