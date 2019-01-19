package main

import (
	"fmt"
	"crypto/sha1"
)

func main() {

	var password = "Password1"
	fmt.Printf("Secure password: %v \n", password)
	
	var bytes = []byte(password)
	fmt.Printf("Password bytes: %v \n", bytes)
	
	hasher := sha1.New()
	hasher.Write(bytes)
	sha := hasher.Sum(nil)

	fmt.Printf("Sha: %v (of type %T) \n", sha, sha)
	
}