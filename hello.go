package main

import "fmt"

func main() {

	var password = "Password1"
	fmt.Printf("Secure password: %v \n", password)
	
	var bytes = []byte(password)
	fmt.Printf("Password bytes: %v \n", bytes)
}