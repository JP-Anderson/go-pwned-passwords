package main

import (
	"fmt"
	"./passwordprocessing"
)

func main() {

	var password = "Password1"
	fmt.Printf("Secure password: %v \n", password)
		
	shaBytes := passwordprocessing.PasswordStringToSHABytes(password)

	fmt.Printf("Sha: %v (of type %T) \n", shaBytes, shaBytes)	
}