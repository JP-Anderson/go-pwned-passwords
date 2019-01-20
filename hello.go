package main

import (
	"fmt"
	"./passwordprocessing"
	"./pwnapi"
)

func main() {
	var password = "Password1"
	passwordShaHex := passwordprocessing.StringToSha1Hex(password)
	fmt.Printf("Sha1(%v) = %v \n", password, passwordShaHex)

	firstFive := passwordShaHex[0:5]
	theRest := passwordShaHex[5:]

	fmt.Printf("%v will be sent to api and %v is compared to the returned options\n", firstFive, theRest)
	
	pwned := pwnapi.PasswordHashIsPwned(passwordShaHex)
	if pwned {
		fmt.Println("Pwned.")
	} else {	
		fmt.Println("Not pwned. Well done!")
	}
}