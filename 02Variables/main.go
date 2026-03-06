package main

import "fmt"

const LoginToken string = "abc1234"

func main() {
	var username string = "Nalinikant"
	fmt.Println(username)
	fmt.Printf("Variable username is of type %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable isLoggedIn is of type %T \n", isLoggedIn)

	var smallVal int = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable smallVal is of type %T \n", smallVal)

	// No var style
	numberOfUsers := 3000
	fmt.Println(numberOfUsers)
	
}



