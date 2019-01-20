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
	fmt.Printf("%v will be sent to api and %v is compared to the returned options", firstFive, theRest)
	
	resp := pwnapi.PasswordHashIsPwned(firstFive)
	fmt.Printf("%v", resp)
}