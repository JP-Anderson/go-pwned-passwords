package passwordprocessing

import (
	"crypto/sha1"
)

func PasswordStringToSHABytes(password string) []uint8 {
	var bytes = []byte(password)
	hasher := sha1.New()
	hasher.Write(bytes)
	sha := hasher.Sum(nil)
	return sha
}