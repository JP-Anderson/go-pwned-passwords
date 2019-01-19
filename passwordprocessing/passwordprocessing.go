package passwordprocessing

import (
	"crypto/sha1"
	"encoding/hex"
	"strings"
)

func StringToSha1Hex(input string) string {
	eightBitInts := PasswordStringToSHABytes(input)
	hex := strings.ToUpper(hex.EncodeToString(eightBitInts))
	return hex
}

func PasswordStringToSHABytes(password string) []uint8 {
	var bytes = []byte(password)
	hasher := sha1.New()
	hasher.Write(bytes)
	sha := hasher.Sum(nil)
	return sha
}