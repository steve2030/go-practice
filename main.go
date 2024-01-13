package main

import "fmt"

func main() {
	var firstName string
	var secondName string

	fmt.Print("Enter your FirstName:")
	fmt.Scan(&firstName)
	fmt.Print("Enter your SecondName:")
	fmt.Scan(&secondName)

	fmt.Printf("My FistName is:%v and my Second name is: %v", firstName, secondName)

}

// print("hello steve")
